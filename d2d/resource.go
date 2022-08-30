//go:build windows

package d2d

import (
	"syscall"
	"unsafe"

	"github.com/bupjae/direct"
)

type IResource struct{ direct.IUnknown }

func (inf *IResource) GetFactory() *IFactory {
	r, _, _ := syscall.SyscallN(inf.VTable[3], uintptr(unsafe.Pointer(inf)))
	//goland:noinspection GoVetUnsafePointer
	return (*IFactory)(unsafe.Pointer(r))
}
