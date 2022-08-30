//go:build windows

package direct

type MeasuringMode uint32

const (
	DWRITE_MEASURING_MODE_NATURAL     MeasuringMode = 0
	DWRITE_MEASURING_MODE_GDI_CLASSIC MeasuringMode = 1
	DWRITE_MEASURING_MODE_GDI_NATURAL MeasuringMode = 2
)
