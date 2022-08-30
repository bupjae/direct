//go:build windows

package main

import (
	"math"
	"strconv"
	"syscall"
	"time"
	"unsafe"

	"github.com/bupjae/direct"
	"github.com/bupjae/direct/d2d"
	"github.com/bupjae/direct/dwrite"
)

var (
	hWnd HWND

	d2dFactory             *d2d.IFactory
	renderTarget           *d2d.IHwndRenderTarget
	gradientStopCollection *d2d.IGradientStopCollection
	brushBlack             *d2d.IBrush
	brushGradient          *d2d.IBrush

	dwriteFactory *dwrite.IFactory
	textFormat    *dwrite.ITextFormat

	ticker *time.Ticker
	done   chan struct{}
)

func createIndependentResource() {
	if d2dFactory == nil {
		f, err := d2d.CreateFactory(d2d.D2D1_FACTORY_TYPE_SINGLE_THREADED, &d2d.FactoryOptions{DebugLevel: d2d.D2D1_DEBUG_LEVEL_NONE})
		if err != nil {
			panic(err)
		}
		d2dFactory = f
	}
	if dwriteFactory == nil {
		f, err := dwrite.CreateFactory(dwrite.DWRITE_FACTORY_TYPE_SHARED)
		if err != nil {
			panic(err)
		}
		dwriteFactory = f
	}
	if textFormat == nil {
		t, err := dwriteFactory.CreateTextFormat("Arial", nil, dwrite.DWRITE_FONT_WEIGHT_REGULAR, dwrite.DWRITE_FONT_STYLE_NORMAL, dwrite.DWRITE_FONT_STRETCH_NORMAL, 10, "en-US")
		if err != nil {
			panic(err)
		}
		if err = t.SetTextAlignment(dwrite.DWRITE_TEXT_ALIGNMENT_CENTER); err != nil {
			panic(err)
		}
		if err = t.SetParagraphAlignment(dwrite.DWRITE_PARAGRAPH_ALIGNMENT_CENTER); err != nil {
			panic(err)
		}
		textFormat = t
	}
	if ticker == nil {
		ticker = time.NewTicker(time.Second / 10)
	}
	if done == nil {
		done = make(chan struct{})
	}
	go func(ticker <-chan time.Time, done <-chan struct{}) {
		for {
			select {
			case <-ticker:
				InvalidateRect(hWnd, nil, false)
			case <-done:
				return
			}
		}
	}(ticker.C, done)
}

func createDependentResource() {
	var rc RECT
	if err := GetClientRect(hWnd, &rc); err != nil {
		panic(err)
	}
	pixelSize := d2d.SizeU{uint32(rc.Right), uint32(rc.Bottom)}

	if renderTarget == nil {
		t, err := d2dFactory.CreateHwndRenderTarget(
			&d2d.RenderTargetProperties{},
			&d2d.HwndRenderTargetProperties{
				Hwnd:      uintptr(hWnd),
				PixelSize: pixelSize,
			},
		)
		if err != nil {
			panic(err)
		}
		renderTarget = t
	} else {
		if renderTarget.GetPixelSize() != pixelSize {
			if err := renderTarget.Resize(&pixelSize); err != nil {
				panic(err)
			}
		}
	}

	dipSize := renderTarget.GetSize()
	d := dipSize.Height
	if d > dipSize.Width {
		d = dipSize.Width
	}
	d *= 0.005
	renderTarget.SetTransform(&d2d.Matrix3x2F{{d, 0}, {0, d}, {dipSize.Width / 2, dipSize.Height / 2}})

	if gradientStopCollection == nil {
		g, err := renderTarget.CreateGradientStopCollection(
			[]d2d.GradientStop{
				{0, d2d.ColorF{1, 0, 0, 1}},
				{0.5, d2d.ColorF{0, 1, 0, 1}},
				{1, d2d.ColorF{0, 0, 1, 1}},
			},
			d2d.D2D1_GAMMA_1_0, d2d.D2D1_EXTEND_MODE_CLAMP)
		if err != nil {
			panic(err)
		}
		gradientStopCollection = g
	}

	if brushBlack == nil {
		b, err := renderTarget.CreateSolidColorBrush(
			&d2d.ColorF{0, 0, 0, 1},
			&d2d.BrushProperties{
				Opacity:   1,
				Transform: d2d.Matrix3x2F{{1, 0}, {0, 1}, {0, 0}}})
		if err != nil {
			panic(err)
		}
		brushBlack = &b.IBrush
	}

	if brushGradient == nil {
		b, err := renderTarget.CreateLinearGradientBrush(
			&d2d.LinearGradientBrushProperties{
				StartPoint: d2d.Point2F{-50, -50},
				EndPoint:   d2d.Point2F{50, 50},
			},
			&d2d.BrushProperties{
				Opacity:   1,
				Transform: d2d.Matrix3x2F{{1, 0}, {0, 1}, {0, 0}},
			},
			gradientStopCollection)
		if err != nil {
			panic(err)
		}
		brushGradient = &b.IBrush
	}
}

func releaseDependentResource() {
	if brushGradient != nil {
		brushGradient.Release()
		brushGradient = nil
	}
	if brushBlack != nil {
		brushBlack.Release()
		brushBlack = nil
	}
	if gradientStopCollection != nil {
		gradientStopCollection.Release()
		gradientStopCollection = nil
	}
	if renderTarget != nil {
		renderTarget.Release()
		renderTarget = nil
	}
}

func releaseIndependentResource() {
	releaseDependentResource()

	if done != nil {
		close(done)
		done = nil
	}
	if ticker != nil {
		ticker.Stop()
		ticker = nil
	}
	if textFormat != nil {
		textFormat.Release()
		textFormat = nil
	}
	if dwriteFactory != nil {
		dwriteFactory.Release()
		dwriteFactory = nil
	}
	if d2dFactory != nil {
		d2dFactory.Release()
		d2dFactory = nil
	}
}

func onPaint() {
	renderTarget.BeginDraw()
	renderTarget.Clear(&d2d.ColorF{0x87 / 255.0, 0xce / 255.0, 0xeb / 255.0, 1}) // SkyBlue = 0x87CEEB,
	renderTarget.FillEllipse(&d2d.Ellipse{d2d.Point2F{0, 0}, 80, 80}, brushGradient)
	t := time.Now()
	hh, mm, ss := float64(t.Hour()), float64(t.Minute()), float64(t.Second())
	s, c := math.Sincos((hh + mm/60 + ss/3600 - 3) / 6 * math.Pi)
	renderTarget.DrawLine(d2d.Point2F{0, 0}, d2d.Point2F{float32(c * 45), float32(s * 45)}, brushBlack, 3, nil)
	s, c = math.Sincos((mm + ss/60 - 15) / 30 * math.Pi)
	renderTarget.DrawLine(d2d.Point2F{0, 0}, d2d.Point2F{float32(c * 60), float32(s * 60)}, brushBlack, 2, nil)
	s, c = math.Sincos((ss - 15) / 30 * math.Pi)
	renderTarget.DrawLine(d2d.Point2F{0, 0}, d2d.Point2F{float32(c * 75), float32(s * 75)}, brushBlack, 1, nil)
	for i := 1; i <= 12; i++ {
		s, c = math.Sincos((float64(i)/6 - 0.5) * math.Pi)
		renderTarget.DrawText(strconv.Itoa(i), textFormat, &d2d.RectF{float32(c*90 - 10), float32(s*90 - 10), float32(c*90 + 10), float32(s*90 + 10)}, brushBlack, d2d.D2D1_DRAW_TEXT_OPTIONS_CLIP, direct.DWRITE_MEASURING_MODE_NATURAL)
	}
	if _, _, err := renderTarget.EndDraw(); err != nil {
		if err == d2d.D2DERR_RECREATE_TARGET {
			if brushGradient != nil {
				brushGradient.Release()
				brushGradient = nil
			}
			if brushBlack != nil {
				brushBlack.Release()
				brushBlack = nil
			}
			if gradientStopCollection != nil {
				gradientStopCollection.Release()
				gradientStopCollection = nil
			}
			if renderTarget != nil {
				renderTarget.Release()
				renderTarget = nil
			}
		} else {
			panic(err)
		}
	} else {
		ValidateRect(hWnd, nil)
	}
}

func wndProc(hWnd HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	switch msg {
	case WM_CREATE:
		createIndependentResource()
		return 0
	case WM_PAINT:
		createDependentResource()
		onPaint()
		return 0
	case WM_SIZE:
		InvalidateRect(hWnd, nil, false)
	case WM_DESTROY:
		releaseIndependentResource()
		PostQuitMessage(0)
		return 0
	}
	return DefWindowProc(hWnd, msg, wParam, lParam)
}

func main() {
	titlePtr, _ := syscall.UTF16PtrFromString("Clock")
	var instance HMODULE
	err := GetModuleHandle(0, nil, &instance)
	if err != nil {
		panic(err)
	}
	icon, err := LoadIcon(0, IDI_APPLICATION)
	if err != nil {
		panic(err)
	}
	cursor, err := LoadCursor(0, IDC_ARROW)
	if err != nil {
		panic(err)
	}
	if _, err = RegisterClass(&WNDCLASSEX{
		Size:      UINT(unsafe.Sizeof(WNDCLASSEX{})),
		WndProc:   syscall.NewCallback(wndProc),
		Instance:  instance,
		Icon:      icon,
		Cursor:    cursor,
		ClassName: titlePtr,
	}); err != nil {
		panic(err)
	}
	hWnd, err = CreateWindow(
		0,
		"Clock",
		"Clock",
		WS_OVERLAPPEDWINDOW,
		CW_USEDEFAULT,
		CW_USEDEFAULT,
		CW_USEDEFAULT,
		CW_USEDEFAULT,
		0,
		0,
		instance,
		nil,
	)
	if err != nil {
		panic(err)
	}
	ShowWindow(hWnd, SW_SHOWDEFAULT)

	var msg MSG
	for {
		r, err := GetMessage(&msg, 0, 0, 0)
		if err != nil {
			panic(err)
		}
		if !r {
			break
		}
		TranslateMessage(&msg)
		DispatchMessage(&msg)
	}
}
