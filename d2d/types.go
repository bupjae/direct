//go:build windows

package d2d

import "math"

type (
	ColorF struct{ R, G, B, A float32 }

	Matrix3x2F = [3][2]float32
	Matrix4x3F = [4][3]float32
	Matrix4x4F = [4][4]float32
	Matrix5x4F = [5][4]float32

	Point2F struct{ X, Y float32 }
	Point2U struct{ X, Y uint32 }
	Point2L struct{ X, Y int32 }

	RectF struct{ Left, Top, Right, Bottom float32 }
	RectU struct{ Left, Top, Right, Bottom uint32 }
	RectL struct{ Left, Top, Right, Bottom int32 }

	SizeF struct{ Width, Height float32 }
	SizeU struct{ Width, Height uint32 }

	Vector2F struct{ X, Y float32 }
	Vector3F struct{ X, Y, Z float32 }
	Vector4F struct{ X, Y, Z, W float32 }
)

type BrushProperties struct {
	Opacity   float32
	Transform Matrix3x2F
}

type Ellipse struct {
	Point2F
	RadiusX, RadiusY float32
}

type FactoryOptions struct {
	DebugLevel DebugLevel
}

type GradientStop struct {
	Position float32
	ColorF
}

type HwndRenderTargetProperties struct {
	Hwnd           uintptr
	PixelSize      SizeU
	PresentOptions PresentOptions
}

type LinearGradientBrushProperties struct {
	StartPoint, EndPoint Point2F
}

type PixelFormat struct {
	Format    DxgiFormat
	AlphaMode AlphaMode
}

type RenderTargetProperties struct {
	Type        RenderTargetType
	PixelFormat PixelFormat
	DpiX, DpiY  float32
	Usage       RenderTargetUsage
	MinLevel    FeatureLevel
}

func (p Point2F) encode() uintptr {
	return (uintptr(math.Float32bits(p.Y)) << 32) | uintptr(math.Float32bits(p.X))
}
