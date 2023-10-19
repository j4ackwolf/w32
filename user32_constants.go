package w32

// https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-enumdisplaymonitors
const (
	WINSTA_ALL_ACCESS        = 0x37F
	WINSTA_ACCESSCLIPBOARD   = 0x0004
	WINSTA_ACCESSGLOBALATOMS = 0x0020
	WINSTA_CREATEDESKTOP     = 0x0008
	WINSTA_ENUMDESKTOPS      = 0x0001
	WINSTA_ENUMERATE         = 0x0100
	WINSTA_EXITWINDOWS       = 0x0040
	WINSTA_READATTRIBUTES    = 0x0002
	WINSTA_READSCREEN        = 0x0200
	WINSTA_WRITEATTRIBUTES   = 0x0010
)

const (
	DESKTOP_CREATEMENU      = 0x0004
	DESKTOP_CREATEWINDOW    = 0x0002
	DESKTOP_ENUMERATE       = 0x0040
	DESKTOP_HOOKCONTROL     = 0x0008
	DESKTOP_JOURNALPLAYBACK = 0x0020
	DESKTOP_JOURNALRECORD   = 0x0010
	DESKTOP_READOBJECTS     = 0x0001
	DESKTOP_SWITCHDESKTOP   = 0x0100
	DESKTOP_WRITEOBJECTS    = 0x0080
)

const (
	CURSOR_SHOWING    = 0x00000001
	CURSOR_SUPPRESSED = 0x00000002
)
