//go:build windows

package dwrite

import (
	"syscall"
	"unsafe"

	"github.com/bupjae/direct"
)

type ITextFormat struct{ direct.IUnknown }

func (inf *ITextFormat) SetTextAlignment(textAlignment TextAlignment) error {
	hr, _, _ := syscall.SyscallN(inf.VTable[3], uintptr(unsafe.Pointer(inf)), uintptr(textAlignment))
	if hr&0x80000000 != 0 {
		return direct.HResult(hr)
	}
	return nil
}

func (inf *ITextFormat) SetParagraphAlignment(paragraphAlignment ParagraphAlignment) error {
	hr, _, _ := syscall.SyscallN(inf.VTable[4], uintptr(unsafe.Pointer(inf)), uintptr(paragraphAlignment))
	if hr&0x80000000 != 0 {
		return direct.HResult(hr)
	}
	return nil
}
