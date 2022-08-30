//go:build windows

package d2d

import (
	"syscall"
	"unsafe"

	"github.com/bupjae/direct"
)

var IIDFactory = &syscall.GUID{0x06152247, 0x6f50, 0x465a, [8]byte{0x92, 0x45, 0x11, 0x8b, 0xfd, 0x3b, 0x60, 0x07}}

type IFactory struct{ direct.IUnknown }

//sys _CreateFactory(factoryType FactoryType, riid *syscall.GUID, factoryOptions *FactoryOptions, factory **IFactory) (hr uintptr) = d2d1.D2D1CreateFactory

func CreateFactory(factoryType FactoryType, options *FactoryOptions) (*IFactory, error) {
	var r *IFactory
	if options == nil {
		options = &FactoryOptions{}
	}
	hr := _CreateFactory(factoryType, IIDFactory, options, &r)
	if hr&0x80000000 != 0 {
		return nil, direct.HResult(hr)
	}
	return r, nil
}

func (inf *IFactory) CreateHwndRenderTarget(renderTargetProperties *RenderTargetProperties, hwndRenderTargetProperties *HwndRenderTargetProperties) (*IHwndRenderTarget, error) {
	var r *IHwndRenderTarget
	hr, _, _ := syscall.SyscallN(inf.VTable[14], uintptr(unsafe.Pointer(inf)), uintptr(unsafe.Pointer(renderTargetProperties)), uintptr(unsafe.Pointer(hwndRenderTargetProperties)), uintptr(unsafe.Pointer(&r)))
	if hr&0x80000000 != 0 {
		return nil, direct.HResult(hr)
	}
	return r, nil
}
