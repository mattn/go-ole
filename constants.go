package ole

import ()

const (
	CLSCTX_INPROC_SERVER   = 1
	CLSCTX_INPROC_HANDLER  = 2
	CLSCTX_LOCAL_SERVER    = 4
	CLSCTX_INPROC_SERVER16 = 8
	CLSCTX_REMOTE_SERVER   = 16
	CLSCTX_ALL             = CLSCTX_INPROC_SERVER | CLSCTX_INPROC_HANDLER | CLSCTX_LOCAL_SERVER
	CLSCTX_INPROC          = CLSCTX_INPROC_SERVER | CLSCTX_INPROC_HANDLER
	CLSCTX_SERVER          = CLSCTX_INPROC_SERVER | CLSCTX_LOCAL_SERVER | CLSCTX_REMOTE_SERVER
)

const (
	COINIT_APARTMENTTHREADED = 0x2
	COINIT_MULTITHREADED     = 0x0
	COINIT_DISABLE_OLE1DDE   = 0x4
	COINIT_SPEED_OVER_MEMORY = 0x8
)

const (
	DISPATCH_METHOD         = 1
	DISPATCH_PROPERTYGET    = 2
	DISPATCH_PROPERTYPUT    = 4
	DISPATCH_PROPERTYPUTREF = 8
)

const (
	S_OK           = 0x00000000
	E_UNEXPECTED   = 0x8000FFFF
	E_NOTIMPL      = 0x80004001
	E_OUTOFMEMORY  = 0x8007000E
	E_INVALIDARG   = 0x80070057
	E_NOINTERFACE  = 0x80004002
	E_POINTER      = 0x80004003
	E_HANDLE       = 0x80070006
	E_ABORT        = 0x80004004
	E_FAIL         = 0x80004005
	E_ACCESSDENIED = 0x80070005
	E_PENDING      = 0x8000000A
)

const (
	CC_FASTCALL = iota
	CC_CDECL
	CC_MSCPASCAL
	CC_PASCAL = CC_MSCPASCAL
	CC_MACPASCAL
	CC_STDCALL
	CC_FPFASTCALL
	CC_SYSCALL
	CC_MPWCDECL
	CC_MPWPASCAL
	CC_MAX = CC_MPWPASCAL
)

const (
	VT_EMPTY           = 0x0
	VT_NULL            = 0x1
	VT_I2              = 0x2
	VT_I4              = 0x3
	VT_R4              = 0x4
	VT_R8              = 0x5
	VT_CY              = 0x6
	VT_DATE            = 0x7
	VT_BSTR            = 0x8
	VT_DISPATCH        = 0x9
	VT_ERROR           = 0xa
	VT_BOOL            = 0xb
	VT_VARIANT         = 0xc
	VT_UNKNOWN         = 0xd
	VT_DECIMAL         = 0xe
	VT_I1              = 0x10
	VT_UI1             = 0x11
	VT_UI2             = 0x12
	VT_UI4             = 0x13
	VT_I8              = 0x14
	VT_UI8             = 0x15
	VT_INT             = 0x16
	VT_UINT            = 0x17
	VT_VOID            = 0x18
	VT_HRESULT         = 0x19
	VT_PTR             = 0x1a
	VT_SAFEARRAY       = 0x1b
	VT_CARRAY          = 0x1c
	VT_USERDEFINED     = 0x1d
	VT_LPSTR           = 0x1e
	VT_LPWSTR          = 0x1f
	VT_RECORD          = 0x24
	VT_INT_PTR         = 0x25
	VT_UINT_PTR        = 0x26
	VT_FILETIME        = 0x40
	VT_BLOB            = 0x41
	VT_STREAM          = 0x42
	VT_STORAGE         = 0x43
	VT_STREAMED_OBJECT = 0x44
	VT_STORED_OBJECT   = 0x45
	VT_BLOB_OBJECT     = 0x46
	VT_CF              = 0x47
	VT_CLSID           = 0x48
	VT_BSTR_BLOB       = 0xfff
	VT_VECTOR          = 0x1000
	VT_ARRAY           = 0x2000
	VT_BYREF           = 0x4000
	VT_RESERVED        = 0x8000
	VT_ILLEGAL         = 0xffff
	VT_ILLEGALMASKED   = 0xfff
	VT_TYPEMASK        = 0xfff
)

const (
	DISPID_UNKNOWN     = -1
	DISPID_VALUE       = 0
	DISPID_PROPERTYPUT = -3
	DISPID_NEWENUM     = -4
	DISPID_EVALUATE    = -5
	DISPID_CONSTRUCTOR = -6
	DISPID_DESTRUCTOR  = -7
	DISPID_COLLECT     = -8
)

const (
	TKIND_ENUM      = 1
	TKIND_RECORD    = 2
	TKIND_MODULE    = 3
	TKIND_INTERFACE = 4
	TKIND_DISPATCH  = 5
	TKIND_COCLASS   = 6
	TKIND_ALIAS     = 7
	TKIND_UNION     = 8
	TKIND_MAX       = 9
)

// Safe Array Feature Flags

const (
	FADF_AUTO        = 0x0001
	FADF_STATIC      = 0x0002
	FADF_EMBEDDED    = 0x0004
	FADF_FIXEDSIZE   = 0x0010
	FADF_RECORD      = 0x0020
	FADF_HAVEIID     = 0x0040
	FADF_HAVEVARTYPE = 0x0080
	FADF_BSTR        = 0x0100
	FADF_UNKNOWN     = 0x0200
	FADF_DISPATCH    = 0x0400
	FADF_VARIANT     = 0x0800
	FADF_RESERVED    = 0xF008
)
