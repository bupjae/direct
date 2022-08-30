//go:build windows

package d2d

import (
	"github.com/bupjae/direct"
)

type DxgiFormat uint32

const (
	DXGI_FORMAT_UNKNOWN                    DxgiFormat = 0
	DXGI_FORMAT_R32G32B32A32_TYPELESS      DxgiFormat = 1
	DXGI_FORMAT_R32G32B32A32_FLOAT         DxgiFormat = 2
	DXGI_FORMAT_R32G32B32A32_UINT          DxgiFormat = 3
	DXGI_FORMAT_R32G32B32A32_SINT          DxgiFormat = 4
	DXGI_FORMAT_R32G32B32_TYPELESS         DxgiFormat = 5
	DXGI_FORMAT_R32G32B32_FLOAT            DxgiFormat = 6
	DXGI_FORMAT_R32G32B32_UINT             DxgiFormat = 7
	DXGI_FORMAT_R32G32B32_SINT             DxgiFormat = 8
	DXGI_FORMAT_R16G16B16A16_TYPELESS      DxgiFormat = 9
	DXGI_FORMAT_R16G16B16A16_FLOAT         DxgiFormat = 10
	DXGI_FORMAT_R16G16B16A16_UNORM         DxgiFormat = 11
	DXGI_FORMAT_R16G16B16A16_UINT          DxgiFormat = 12
	DXGI_FORMAT_R16G16B16A16_SNORM         DxgiFormat = 13
	DXGI_FORMAT_R16G16B16A16_SINT          DxgiFormat = 14
	DXGI_FORMAT_R32G32_TYPELESS            DxgiFormat = 15
	DXGI_FORMAT_R32G32_FLOAT               DxgiFormat = 16
	DXGI_FORMAT_R32G32_UINT                DxgiFormat = 17
	DXGI_FORMAT_R32G32_SINT                DxgiFormat = 18
	DXGI_FORMAT_R32G8X24_TYPELESS          DxgiFormat = 19
	DXGI_FORMAT_D32_FLOAT_S8X24_UINT       DxgiFormat = 20
	DXGI_FORMAT_R32_FLOAT_X8X24_TYPELESS   DxgiFormat = 21
	DXGI_FORMAT_X32_TYPELESS_G8X24_UINT    DxgiFormat = 22
	DXGI_FORMAT_R10G10B10A2_TYPELESS       DxgiFormat = 23
	DXGI_FORMAT_R10G10B10A2_UNORM          DxgiFormat = 24
	DXGI_FORMAT_R10G10B10A2_UINT           DxgiFormat = 25
	DXGI_FORMAT_R11G11B10_FLOAT            DxgiFormat = 26
	DXGI_FORMAT_R8G8B8A8_TYPELESS          DxgiFormat = 27
	DXGI_FORMAT_R8G8B8A8_UNORM             DxgiFormat = 28
	DXGI_FORMAT_R8G8B8A8_UNORM_SRGB        DxgiFormat = 29
	DXGI_FORMAT_R8G8B8A8_UINT              DxgiFormat = 30
	DXGI_FORMAT_R8G8B8A8_SNORM             DxgiFormat = 31
	DXGI_FORMAT_R8G8B8A8_SINT              DxgiFormat = 32
	DXGI_FORMAT_R16G16_TYPELESS            DxgiFormat = 33
	DXGI_FORMAT_R16G16_FLOAT               DxgiFormat = 34
	DXGI_FORMAT_R16G16_UNORM               DxgiFormat = 35
	DXGI_FORMAT_R16G16_UINT                DxgiFormat = 36
	DXGI_FORMAT_R16G16_SNORM               DxgiFormat = 37
	DXGI_FORMAT_R16G16_SINT                DxgiFormat = 38
	DXGI_FORMAT_R32_TYPELESS               DxgiFormat = 39
	DXGI_FORMAT_D32_FLOAT                  DxgiFormat = 40
	DXGI_FORMAT_R32_FLOAT                  DxgiFormat = 41
	DXGI_FORMAT_R32_UINT                   DxgiFormat = 42
	DXGI_FORMAT_R32_SINT                   DxgiFormat = 43
	DXGI_FORMAT_R24G8_TYPELESS             DxgiFormat = 44
	DXGI_FORMAT_D24_UNORM_S8_UINT          DxgiFormat = 45
	DXGI_FORMAT_R24_UNORM_X8_TYPELESS      DxgiFormat = 46
	DXGI_FORMAT_X24_TYPELESS_G8_UINT       DxgiFormat = 47
	DXGI_FORMAT_R8G8_TYPELESS              DxgiFormat = 48
	DXGI_FORMAT_R8G8_UNORM                 DxgiFormat = 49
	DXGI_FORMAT_R8G8_UINT                  DxgiFormat = 50
	DXGI_FORMAT_R8G8_SNORM                 DxgiFormat = 51
	DXGI_FORMAT_R8G8_SINT                  DxgiFormat = 52
	DXGI_FORMAT_R16_TYPELESS               DxgiFormat = 53
	DXGI_FORMAT_R16_FLOAT                  DxgiFormat = 54
	DXGI_FORMAT_D16_UNORM                  DxgiFormat = 55
	DXGI_FORMAT_R16_UNORM                  DxgiFormat = 56
	DXGI_FORMAT_R16_UINT                   DxgiFormat = 57
	DXGI_FORMAT_R16_SNORM                  DxgiFormat = 58
	DXGI_FORMAT_R16_SINT                   DxgiFormat = 59
	DXGI_FORMAT_R8_TYPELESS                DxgiFormat = 60
	DXGI_FORMAT_R8_UNORM                   DxgiFormat = 61
	DXGI_FORMAT_R8_UINT                    DxgiFormat = 62
	DXGI_FORMAT_R8_SNORM                   DxgiFormat = 63
	DXGI_FORMAT_R8_SINT                    DxgiFormat = 64
	DXGI_FORMAT_A8_UNORM                   DxgiFormat = 65
	DXGI_FORMAT_R1_UNORM                   DxgiFormat = 66
	DXGI_FORMAT_R9G9B9E5_SHAREDEXP         DxgiFormat = 67
	DXGI_FORMAT_R8G8_B8G8_UNORM            DxgiFormat = 68
	DXGI_FORMAT_G8R8_G8B8_UNORM            DxgiFormat = 69
	DXGI_FORMAT_BC1_TYPELESS               DxgiFormat = 70
	DXGI_FORMAT_BC1_UNORM                  DxgiFormat = 71
	DXGI_FORMAT_BC1_UNORM_SRGB             DxgiFormat = 72
	DXGI_FORMAT_BC2_TYPELESS               DxgiFormat = 73
	DXGI_FORMAT_BC2_UNORM                  DxgiFormat = 74
	DXGI_FORMAT_BC2_UNORM_SRGB             DxgiFormat = 75
	DXGI_FORMAT_BC3_TYPELESS               DxgiFormat = 76
	DXGI_FORMAT_BC3_UNORM                  DxgiFormat = 77
	DXGI_FORMAT_BC3_UNORM_SRGB             DxgiFormat = 78
	DXGI_FORMAT_BC4_TYPELESS               DxgiFormat = 79
	DXGI_FORMAT_BC4_UNORM                  DxgiFormat = 80
	DXGI_FORMAT_BC4_SNORM                  DxgiFormat = 81
	DXGI_FORMAT_BC5_TYPELESS               DxgiFormat = 82
	DXGI_FORMAT_BC5_UNORM                  DxgiFormat = 83
	DXGI_FORMAT_BC5_SNORM                  DxgiFormat = 84
	DXGI_FORMAT_B5G6R5_UNORM               DxgiFormat = 85
	DXGI_FORMAT_B5G5R5A1_UNORM             DxgiFormat = 86
	DXGI_FORMAT_B8G8R8A8_UNORM             DxgiFormat = 87
	DXGI_FORMAT_B8G8R8X8_UNORM             DxgiFormat = 88
	DXGI_FORMAT_R10G10B10_XR_BIAS_A2_UNORM DxgiFormat = 89
	DXGI_FORMAT_B8G8R8A8_TYPELESS          DxgiFormat = 90
	DXGI_FORMAT_B8G8R8A8_UNORM_SRGB        DxgiFormat = 91
	DXGI_FORMAT_B8G8R8X8_TYPELESS          DxgiFormat = 92
	DXGI_FORMAT_B8G8R8X8_UNORM_SRGB        DxgiFormat = 93
	DXGI_FORMAT_BC6H_TYPELESS              DxgiFormat = 94
	DXGI_FORMAT_BC6H_UF16                  DxgiFormat = 95
	DXGI_FORMAT_BC6H_SF16                  DxgiFormat = 96
	DXGI_FORMAT_BC7_TYPELESS               DxgiFormat = 97
	DXGI_FORMAT_BC7_UNORM                  DxgiFormat = 98
	DXGI_FORMAT_BC7_UNORM_SRGB             DxgiFormat = 99
	DXGI_FORMAT_FORCE_UINT                 DxgiFormat = 0xffffffff
)

type AlphaMode uint32

const (
	D2D1_ALPHA_MODE_UNKNOWN       AlphaMode = 0
	D2D1_ALPHA_MODE_PREMULTIPLIED AlphaMode = 1
	D2D1_ALPHA_MODE_STRAIGHT      AlphaMode = 2
	D2D1_ALPHA_MODE_IGNORE        AlphaMode = 3
)

type DebugLevel uint32

const (
	D2D1_DEBUG_LEVEL_NONE        DebugLevel = 0
	D2D1_DEBUG_LEVEL_ERROR       DebugLevel = 1
	D2D1_DEBUG_LEVEL_WARNING     DebugLevel = 2
	D2D1_DEBUG_LEVEL_INFORMATION DebugLevel = 3
)

type DrawTextOptions uint32

const (
	D2D1_DRAW_TEXT_OPTIONS_NO_SNAP                       DrawTextOptions = 0
	D2D1_DRAW_TEXT_OPTIONS_CLIP                          DrawTextOptions = 1
	D2D1_DRAW_TEXT_OPTIONS_ENABLE_COLOR_FONT             DrawTextOptions = 2
	D2D1_DRAW_TEXT_OPTIONS_DISABLE_COLOR_BITMAP_SNAPPING DrawTextOptions = 3
	D2D1_DRAW_TEXT_OPTIONS_NONE                          DrawTextOptions = 4
)

type ExtendMode uint32

const (
	D2D1_EXTEND_MODE_CLAMP  ExtendMode = 0
	D2D1_EXTEND_MODE_WRAP   ExtendMode = 1
	D2D1_EXTEND_MODE_MIRROR ExtendMode = 2
)

type FactoryType uint32

const (
	D2D1_FACTORY_TYPE_SINGLE_THREADED FactoryType = 0
	D2D1_FACTORY_TYPE_MULTI_THREADED  FactoryType = 1
)

type FeatureLevel uint32

const (
	D2D1_FEATURE_LEVEL_DEFAULT FeatureLevel = 0
	D2D1_FEATURE_LEVEL_9       FeatureLevel = 0x9100 /* D3D10_FEATURE_LEVEL_9_1 */
	D2D1_FEATURE_LEVEL_10      FeatureLevel = 0xa000 /* D3D10_FEATURE_LEVEL_10_0 */
)

type Gamma uint32

const (
	D2D1_GAMMA_2_2 Gamma = 0
	D2D1_GAMMA_1_0 Gamma = 1
)

type PresentOptions uint32

const (
	D2D1_PRESENT_OPTIONS_NONE            PresentOptions = 0x00000000
	D2D1_PRESENT_OPTIONS_RETAIN_CONTENTS PresentOptions = 0x00000001
	D2D1_PRESENT_OPTIONS_IMMEDIATELY     PresentOptions = 0x00000002
)

type RenderTargetType uint32

const (
	D2D1_RENDER_TARGET_TYPE_DEFAULT  RenderTargetType = 0
	D2D1_RENDER_TARGET_TYPE_SOFTWARE RenderTargetType = 1
	D2D1_RENDER_TARGET_TYPE_HARDWARE RenderTargetType = 2
)

type RenderTargetUsage uint32

const (
	D2D1_RENDER_TARGET_USAGE_NONE                  RenderTargetUsage = 0x00000000
	D2D1_RENDER_TARGET_USAGE_FORCE_BITMAP_REMOTING RenderTargetUsage = 0x00000001
	D2D1_RENDER_TARGET_USAGE_GDI_COMPATIBLE        RenderTargetUsage = 0x00000002
)

const (
	D2DERR_RECREATE_TARGET = direct.HResult(0x8899000C)
)
