//go:build windows

package dwrite

import (
	"math"
	"syscall"
	"unsafe"

	"github.com/bupjae/direct"
)

var IIDFactory = &syscall.GUID{0xb859ee5a, 0xd838, 0x4b5b, [8]byte{0xa2, 0xe8, 0x1a, 0xdc, 0x7d, 0x93, 0xdb, 0x48}}

type IFactory struct{ direct.IUnknown }

//sys _CreateFactory(factoryType FactoryType, riid *syscall.GUID, factory **IFactory) (hr uintptr) = dwrite.DWriteCreateFactory

func CreateFactory(factoryType FactoryType) (*IFactory, error) {
	var r *IFactory
	hr := _CreateFactory(factoryType, IIDFactory, &r)
	if hr&0x80000000 != 0 {
		return nil, direct.HResult(hr)
	}
	return r, nil
}

func (inf *IFactory) CreateTextFormat(fontFamilyName string, fontCollection *IFontCollection, fontWeight FontWeight, fontStyle FontStyle, fontStretch FontStretch, fontSize float32, localName string) (*ITextFormat, error) {
	var r *ITextFormat
	fn, err := syscall.UTF16PtrFromString(fontFamilyName)
	if err != nil {
		return nil, err
	}
	ln, err := syscall.UTF16PtrFromString(localName)
	if err != nil {
		return nil, err
	}
	hr, _, _ := syscall.SyscallN(inf.VTable[15], uintptr(unsafe.Pointer(inf)),
		uintptr(unsafe.Pointer(fn)),
		uintptr(unsafe.Pointer(fontCollection)),
		uintptr(fontWeight),
		uintptr(fontStyle),
		uintptr(fontStretch),
		uintptr(math.Float32bits(fontSize)),
		uintptr(unsafe.Pointer(ln)),
		uintptr(unsafe.Pointer(&r)),
	)
	if hr&0x80000000 != 0 {
		return nil, direct.HResult(hr)
	}
	return r, nil
}
