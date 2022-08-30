//go:build windows

package d2d

import (
	"math"
	"syscall"
	"unicode/utf16"
	"unsafe"

	"github.com/bupjae/direct"
	"github.com/bupjae/direct/dwrite"
)

type IRenderTarget struct{ IResource }

func (inf *IRenderTarget) CreateSolidColorBrush(color *ColorF, prop *BrushProperties) (*ISolidColorBrush, error) {
	var r *ISolidColorBrush
	hr, _, _ := syscall.SyscallN(inf.VTable[8], uintptr(unsafe.Pointer(inf)), uintptr(unsafe.Pointer(color)), uintptr(unsafe.Pointer(prop)), uintptr(unsafe.Pointer(&r)))
	if hr&0x80000000 != 0 {
		return nil, direct.HResult(hr)
	}
	return r, nil
}

func (inf *IRenderTarget) CreateGradientStopCollection(stop []GradientStop, gamma Gamma, mode ExtendMode) (*IGradientStopCollection, error) {
	var r *IGradientStopCollection
	if len(stop) == 0 {
		return nil, syscall.EINVAL
	}
	hr, _, _ := syscall.SyscallN(inf.VTable[9], uintptr(unsafe.Pointer(inf)), uintptr(unsafe.Pointer(&stop[0])), uintptr(len(stop)), uintptr(gamma), uintptr(mode), uintptr(unsafe.Pointer(&r)))
	if hr&0x80000000 != 0 {
		return nil, direct.HResult(hr)
	}
	return r, nil
}

func (inf *IRenderTarget) CreateLinearGradientBrush(lprop *LinearGradientBrushProperties, bprop *BrushProperties, stop *IGradientStopCollection) (*ILinearGradientBrush, error) {
	var r *ILinearGradientBrush
	hr, _, _ := syscall.SyscallN(inf.VTable[10], uintptr(unsafe.Pointer(inf)), uintptr(unsafe.Pointer(lprop)), uintptr(unsafe.Pointer(bprop)), uintptr(unsafe.Pointer(stop)), uintptr(unsafe.Pointer(&r)))
	if hr&0x80000000 != 0 {
		return nil, direct.HResult(hr)
	}
	return r, nil
}

func (inf *IResource) DrawLine(point0 Point2F, point1 Point2F, brush *IBrush, strokeWidth float32, strokeStyle *IStrokeStyle) {
	_, _, _ = syscall.SyscallN(inf.VTable[15], uintptr(unsafe.Pointer(inf)), point0.encode(), point1.encode(), uintptr(unsafe.Pointer(brush)), uintptr(math.Float32bits(strokeWidth)), uintptr(unsafe.Pointer(strokeStyle)))
}

func (inf *IResource) FillEllipse(ellipse *Ellipse, brush *IBrush) {
	_, _, _ = syscall.SyscallN(inf.VTable[21], uintptr(unsafe.Pointer(inf)), uintptr(unsafe.Pointer(ellipse)), uintptr(unsafe.Pointer(brush)))
}

func (inf *IResource) DrawText(text string, textFormat *dwrite.ITextFormat, layoutRect *RectF, defaultFillBrush *IBrush, options DrawTextOptions, measuringMode direct.MeasuringMode) {
	var p *uint16
	t := utf16.Encode([]rune(text))
	if len(t) > 0 {
		p = &t[0]
	}
	_, _, _ = syscall.SyscallN(inf.VTable[27], uintptr(unsafe.Pointer(inf)),
		uintptr(unsafe.Pointer(p)),
		uintptr(len(t)),
		uintptr(unsafe.Pointer(textFormat)),
		uintptr(unsafe.Pointer(layoutRect)),
		uintptr(unsafe.Pointer(defaultFillBrush)),
		uintptr(options),
		uintptr(measuringMode),
	)
}

func (inf *IResource) SetTransform(matrix *Matrix3x2F) {
	_, _, _ = syscall.SyscallN(inf.VTable[30], uintptr(unsafe.Pointer(inf)), uintptr(unsafe.Pointer(matrix)))
}

func (inf *IResource) Clear(clearColor *ColorF) {
	_, _, _ = syscall.SyscallN(inf.VTable[47], uintptr(unsafe.Pointer(inf)), uintptr(unsafe.Pointer(clearColor)))
}

func (inf *IResource) BeginDraw() {
	_, _, _ = syscall.SyscallN(inf.VTable[48], uintptr(unsafe.Pointer(inf)))
}

func (inf *IResource) EndDraw() (uint64, uint64, error) {
	var t1, t2 uint64
	hr, _, _ := syscall.SyscallN(inf.VTable[49], uintptr(unsafe.Pointer(inf)), uintptr(unsafe.Pointer(&t1)), uintptr(unsafe.Pointer(&t2)))
	if hr&0x80000000 != 0 {
		return t1, t2, direct.HResult(hr)
	}
	return t1, t2, nil
}

func (inf *IResource) GetSize() SizeF {
	var r SizeF
	_, _, _ = syscall.SyscallN(inf.VTable[53], uintptr(unsafe.Pointer(inf)), uintptr(unsafe.Pointer(&r)))
	return r
}

func (inf *IResource) GetPixelSize() SizeU {
	var r SizeU
	_, _, _ = syscall.SyscallN(inf.VTable[54], uintptr(unsafe.Pointer(inf)), uintptr(unsafe.Pointer(&r)))
	return r
}
