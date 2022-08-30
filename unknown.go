//go:build windows

package direct

import (
	"syscall"
	"unsafe"
)

var IIDUnknown = &syscall.GUID{0x00000000, 0x0000, 0x0000, [8]byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

type IUnknown struct{ *VTable }

func (inf *IUnknown) QueryInterface(riid *syscall.GUID) (unsafe.Pointer, error) {
	var r unsafe.Pointer
	hr, _, _ := syscall.SyscallN(inf.VTable[0], uintptr(unsafe.Pointer(inf)), uintptr(unsafe.Pointer(riid)), uintptr(unsafe.Pointer(&r)))
	if hr&0x80000000 != 0 {
		return nil, HResult(hr)
	}
	return r, nil
}

func (inf *IUnknown) AddRef() uint32 {
	r, _, _ := syscall.SyscallN(inf.VTable[1], uintptr(unsafe.Pointer(inf)))
	return uint32(r)
}

func (inf *IUnknown) Release() uint32 {
	r, _, _ := syscall.SyscallN(inf.VTable[2], uintptr(unsafe.Pointer(inf)))
	return uint32(r)
}
