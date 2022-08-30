//go:build windows

package main

import "syscall"

type (
	ATOM      = uint16
	DWORD     = uint32
	HANDLE    = syscall.Handle
	HBRUSH    = HANDLE
	HCURSOR   = HANDLE
	HICON     = HANDLE
	HINSTANCE = HANDLE
	HMENU     = HANDLE
	HMODULE   = HANDLE
	HWND      = HANDLE
	LONG      = int32
	LPARAM    = uintptr
	LRESULT   = uintptr
	UINT      = uint32
	WPARAM    = uintptr
)
type POINT struct {
	X, Y LONG
}

type WNDCLASSEX struct {
	Size       UINT
	Style      UINT
	WndProc    uintptr
	ClsExtra   int32
	WndExtra   int32
	Instance   HINSTANCE
	Icon       HICON
	Cursor     HCURSOR
	Background HBRUSH
	MenuName   *uint16
	ClassName  *uint16
	IconSm     HICON
}

type MSG struct {
	Wnd     HWND
	Message UINT
	WParam  WPARAM
	LParam  LPARAM
	Time    DWORD
	Pt      POINT
	Private DWORD
}

type RECT struct {
	Left, Top, Right, Bottom LONG
}

const (
	IDI_APPLICATION uintptr = 32512

	IDC_ARROW uintptr = 32512

	WS_CAPTION          = 0x00C00000
	WS_MAXIMIZEBOX      = 0x00010000
	WS_MINIMIZEBOX      = 0x00020000
	WS_OVERLAPPED       = 0x00000000
	WS_OVERLAPPEDWINDOW = WS_OVERLAPPED | WS_CAPTION | WS_SYSMENU | WS_THICKFRAME | WS_MINIMIZEBOX | WS_MAXIMIZEBOX
	WS_SYSMENU          = 0x00080000
	WS_THICKFRAME       = 0x00040000

	CW_USEDEFAULT = ^int32(0x7FFFFFFF)

	SW_SHOWDEFAULT = 10

	WM_CREATE  = 0x0001
	WM_DESTROY = 0x0002
	WM_SIZE    = 0x0005
	WM_PAINT   = 0x000F
)

//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output zsyscall_windows.go win32.go

//sys GetModuleHandle(flags DWORD, moduleName *uint16, module *HMODULE) (err error) = kernel32.GetModuleHandleExW
//sys CreateWindow(exStyle DWORD, className string, windowName string, style DWORD, x int32, y int32, width int32, height int32, wndParent HWND, menu HMENU, instance HINSTANCE, param unsafe.Pointer) (wnd HWND, err error) = user32.CreateWindowExW
//sys DefWindowProc(wnd HWND, msg UINT, wParam WPARAM, lParam LPARAM) (r LRESULT) = user32.DefWindowProcW
//sys DispatchMessage(msg *MSG) (r LRESULT) = user32.DispatchMessageW
//sys GetMessage(msg *MSG, wnd HWND, msgFilterMin UINT, msgFilterMax UINT) (b bool, err error) [r0==^uintptr(0)] = user32.GetMessageW
//sys LoadCursor(instance HINSTANCE, cursorName uintptr) (cursor HCURSOR, err error) = user32.LoadCursorW
//sys PostQuitMessage(exitCode int32) = user32.PostQuitMessage
//sys RegisterClass(class *WNDCLASSEX) (a ATOM, err error) = user32.RegisterClassExW
//sys ShowWindow(wnd HWND, cmdShow int32) (b bool) = user32.ShowWindow
//sys TranslateMessage(msg *MSG) (b bool) = user32.TranslateMessage
//sys InvalidateRect(wnd HWND, rect *RECT, erase bool) (b bool) = user32.InvalidateRect
//sys LoadIcon(instance HINSTANCE, iconName uintptr) (icon HICON, err error) = user32.LoadIconW
//sys ValidateRect(wnd HWND, rect *RECT) (b bool) = user32.ValidateRect
//sys GetClientRect(wnd HWND, rect *RECT) (err error) = user32.GetClientRect
