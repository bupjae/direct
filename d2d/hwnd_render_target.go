//go:build windows

package d2d

import (
	"syscall"
	"unsafe"

	"github.com/bupjae/direct"
)

type IHwndRenderTarget struct{ IRenderTarget }

func (inf *IHwndRenderTarget) Resize(pixelSize *SizeU) error {
	hr, _, _ := syscall.SyscallN(inf.VTable[58], uintptr(unsafe.Pointer(inf)), uintptr(unsafe.Pointer(pixelSize)))
	if hr&0x80000000 != 0 {
		return direct.HResult(hr)
	}
	return nil
}
