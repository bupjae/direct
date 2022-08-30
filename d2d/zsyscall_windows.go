// Code generated by 'go generate'; DO NOT EDIT.

package d2d

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modd2d1 = windows.NewLazySystemDLL("d2d1.dll")

	procD2D1CreateFactory = modd2d1.NewProc("D2D1CreateFactory")
)

func _CreateFactory(factoryType FactoryType, riid *syscall.GUID, factoryOptions *FactoryOptions, factory **IFactory) (hr uintptr) {
	r0, _, _ := syscall.Syscall6(procD2D1CreateFactory.Addr(), 4, uintptr(factoryType), uintptr(unsafe.Pointer(riid)), uintptr(unsafe.Pointer(factoryOptions)), uintptr(unsafe.Pointer(factory)), 0, 0)
	hr = uintptr(r0)
	return
}
