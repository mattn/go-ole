// +build windows

package ole

import (
	"reflect"
	"testing"
)

func TestIDispatch(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error(r)
		}
	}()

	var err error

	err = CoInitialize(0)
	if err != nil {
		t.Fatal(err)
	}

	defer CoUninitialize()

	var unknown *IUnknown
	var dispatch *IDispatch

	// oleutil.CreateObject()
	unknown, err = CreateInstance(CLSID_COMEchoTestObject, IID_IUnknown)
	if err != nil {
		t.Error(err)
		return
	}
	defer unknown.Release()

	dispatch, err = unknown.QueryInterface(IID_ICOMEchoTestObject)
	if err != nil {
		t.Fatal(err)
		return
	}
	defer dispatch.Release()

	echoValue := func(method string, value interface{}) (interface{}, bool) {
		var dispid []int32
		var err error

		dispid, err = dispatch.GetIDsOfName([]string{method})
		if err != nil {
			t.Fatal(err)
			return nil, false
		}

		result, err := dispatch.Invoke(dispid[0], DISPATCH_METHOD, value)
		if err != nil {
			t.Fatal(err)
			return nil, false
		}

		return result.Value(), true
	}

	methods := map[string]interface{}{
		"EchoInt8":   int8(1),
		"EchoInt16":  int16(1),
		"EchoInt64":  int64(1),
		"EchoUInt8":  uint8(1),
		"EchoUInt16": uint16(1),
		"EchoUInt64": uint64(1)}

	for method, expected := range methods {
		if actual, passed := echoValue(method, expected); passed {
			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("%s() expected %v did not match %v", method, expected, actual)
			}
		}
	}

	if actual, passed := echoValue("EchoInt32", int32(2)); passed {
		value := actual.(int32)
		if value != int32(2) {
			t.Errorf("%s() expected %v did not match %v", "EchoInt32", int32(2), value)
		}
	}

	if actual, passed := echoValue("EchoUInt32", uint32(4)); passed {
		value := actual.(uint32)
		if value != uint32(4) {
			t.Errorf("%s() expected %v did not match %v", "EchoUInt32", uint32(4), value)
		}
	}

	if actual, passed := echoValue("EchoFloat32", float32(2.2)); passed {
		value := actual.(float32)
		if value != uint32(2.2) {
			t.Errorf("%s() expected %v did not match %v", "EchoFloat32", uint32(2.2), value)
		}
	}

	if actual, passed := echoValue("EchoFloat64", float64(2.2)); passed {
		value := actual.(float64)
		if value != float64(2.2) {
			t.Errorf("%s() expected %v did not match %v", "EchoFloat64", float64(2.2), value)
		}
	}

	if actual, passed := echoValue("EchoString", "Test String"); passed {
		value := actual.(string)
		if value != "Test String" {
			t.Errorf("%s() expected %v did not match %v", "EchoString", "Test String", value)
		}
	}
}
