// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs _mxcfb_h.go

package fb

import (
	"github.com/gotang/godefs"
	"unsafe"
)

const (
	MXC_DISP_SPEC_DEV	= 0
	MXC_DISP_DDC_DEV	= 1
)

const (
	MXCFB_REFRESH_OFF	= iota
	MXCFB_REFRESH_AUTO
	MXCFB_REFRESH_PARTIAL
)

const (
	FB_SYNC_OE_LOW_ACT	uintptr	= 0x80000000
	FB_SYNC_CLK_LAT_FALL	uintptr	= 0x40000000
	FB_SYNC_DATA_INVERT	uintptr	= 0x20000000
	FB_SYNC_CLK_IDLE_EN	uintptr	= 0x10000000
	FB_SYNC_SHARP_MODE	uintptr	= 0x8000000
	FB_SYNC_SWAP_RGB	uintptr	= 0x4000000
	FB_ACCEL_TRIPLE_FLAG		= 0x0
	FB_ACCEL_DOUBLE_FLAG		= 0x1
)

type Mxcfb_gbl_alpha struct {
	Enable	int32
	Alpha	int32
}

type Mxcfb_loc_alpha struct {
	Enable		int32
	In_pixel	int32
	Phy_addr0	uint64
	Phy_addr1	uint64
}

type Mxcfb_color_key struct {
	Enable	int32
	Key	uint32
}

type Mxcfb_pos struct {
	X	uint16
	Y	uint16
}

type Mxcfb_gamma struct {
	Enable	int32
	Constk	[16]int32
	Slopek	[16]int32
}

type Mxcfb_rect struct {
	Top	uint32
	Left	uint32
	Width	uint32
	Height	uint32
}

const (
	GRAYSCALE_8BIT		= 0x1
	GRAYSCALE_8BIT_INVERTED	= 0x2
	GRAYSCALE_4BIT		= 0x3
	GRAYSCALE_4BIT_INVERTED	= 0x4

	AUTO_UPDATE_MODE_REGION_MODE	= 0
	AUTO_UPDATE_MODE_AUTOMATIC_MODE	= 1

	UPDATE_SCHEME_SNAPSHOT		= 0
	UPDATE_SCHEME_QUEUE		= 1
	UPDATE_SCHEME_QUEUE_AND_MERGE	= 2

	UPDATE_MODE_PARTIAL	= 0x0
	UPDATE_MODE_FULL	= 0x1

	WAVEFORM_MODE_AUTO	= 257

	TEMP_USE_AMBIENT	= 0x1000

	EPDC_FLAG_ENABLE_INVERSION	= 0x01
	EPDC_FLAG_FORCE_MONOCHROME	= 0x02
	EPDC_FLAG_USE_CMAP		= 0x04
	EPDC_FLAG_USE_ALT_BUFFER	= 0x100
	EPDC_FLAG_TEST_COLLISION	= 0x200
	EPDC_FLAG_GROUP_UPDATE		= 0x400
	EPDC_FLAG_USE_DITHERING_Y1	= 0x2000
	EPDC_FLAG_USE_DITHERING_Y4	= 0x4000

	FB_POWERDOWN_DISABLE	= -1
)

type Mxcfb_alt_buffer_data struct {
	Phys_addr		uint32
	Width			uint32
	Height			uint32
	Alt_update_region	Mxcfb_rect
}

type Mxcfb_update_data struct {
	Update_region	Mxcfb_rect
	Waveform_mode	uint32
	Update_mode	uint32
	Update_marker	uint32
	Temp		int32
	Flags		uint32
	Alt_buffer_data	Mxcfb_alt_buffer_data
}

type Mxcfb_update_marker_data struct {
	Update_marker	uint32
	Collision_test	uint32
}

type Mxcfb_waveform_modes struct {
	Init	int32
	Du	int32
	Gc4	int32
	Gc8	int32
	Gc16	int32
	Gc32	int32
}

type Mxcfb_csc_matrix struct {
	Param [5][3]int32
}

var (
	MXCFB_WAIT_FOR_VSYNC	int
	MXCFB_SET_GBL_ALPHA	int
	MXCFB_SET_CLR_KEY	int
	MXCFB_SET_OVERLAY_POS	int
	MXCFB_GET_FB_IPU_CHAN	int
	MXCFB_SET_LOC_ALPHA	int
	MXCFB_SET_LOC_ALP_BUF	int
	MXCFB_SET_GAMMA		int
	MXCFB_GET_FB_IPU_DI	int
	MXCFB_GET_DIFMT		int
	MXCFB_GET_FB_BLANK	int
	MXCFB_SET_DIFMT		int
	MXCFB_CSC_UPDATE	int

	MXCFB_SET_WAVEFORM_MODES	int
	MXCFB_SET_TEMPERATURE		int
	MXCFB_SET_AUTO_UPDATE_MODE	int
	MXCFB_SEND_UPDATE		int
	MXCFB_SET_PWRDOWN_DELAY		int
	MXCFB_GET_PWRDOWN_DELAY		int
	MXCFB_SET_UPDATE_SCHEME		int
	MXCFB_GET_WORK_BUFFER		int
	MXCFB_DISABLE_EPDC_ACCESS	int
	MXCFB_ENABLE_EPDC_ACCESS	int
)

func init() {
	MXCFB_WAIT_FOR_VSYNC = godefs.IOW('F', 0x20, int(unsafe.Sizeof(uint32(0))))
	MXCFB_SET_GBL_ALPHA = godefs.IOW('F', 0x21, int(unsafe.Sizeof(Mxcfb_gbl_alpha{})))
	MXCFB_SET_CLR_KEY = godefs.IOW('F', 0x22, int(unsafe.Sizeof(Mxcfb_color_key{})))
	MXCFB_SET_OVERLAY_POS = godefs.IOWR('F', 0x24, int(unsafe.Sizeof(Mxcfb_pos{})))
	MXCFB_GET_FB_IPU_CHAN = godefs.IOR('F', 0x25, int(unsafe.Sizeof(uint32(0))))
	MXCFB_SET_LOC_ALPHA = godefs.IOWR('F', 0x26, int(unsafe.Sizeof(Mxcfb_loc_alpha{})))
	MXCFB_SET_LOC_ALP_BUF = godefs.IOW('F', 0x27, int(unsafe.Sizeof(uint64(0))))
	MXCFB_SET_GAMMA = godefs.IOW('F', 0x28, int(unsafe.Sizeof(Mxcfb_gamma{})))
	MXCFB_GET_FB_IPU_DI = godefs.IOR('F', 0x29, int(unsafe.Sizeof(uint32(0))))
	MXCFB_GET_DIFMT = godefs.IOR('F', 0x2A, int(unsafe.Sizeof(uint32(0))))
	MXCFB_GET_FB_BLANK = godefs.IOR('F', 0x2B, int(unsafe.Sizeof(uint32(0))))
	MXCFB_SET_DIFMT = godefs.IOW('F', 0x2C, int(unsafe.Sizeof(uint32(0))))
	MXCFB_CSC_UPDATE = godefs.IOW('F', 0x2D, int(unsafe.Sizeof(Mxcfb_csc_matrix{})))

	MXCFB_SET_WAVEFORM_MODES = godefs.IOW('F', 0x2B, int(unsafe.Sizeof(Mxcfb_waveform_modes{})))
	MXCFB_SET_TEMPERATURE = godefs.IOW('F', 0x2C, int(unsafe.Sizeof(uint32(0))))
	MXCFB_SET_AUTO_UPDATE_MODE = godefs.IOW('F', 0x2D, int(unsafe.Sizeof(uint32(0))))
	MXCFB_SEND_UPDATE = godefs.IOW('F', 0x2E, int(unsafe.Sizeof(Mxcfb_update_marker_data{})))
	MXCFB_SET_PWRDOWN_DELAY = godefs.IOW('F', 0x30, int(unsafe.Sizeof(uint32(0))))
	MXCFB_GET_PWRDOWN_DELAY = godefs.IOR('F', 0x31, int(unsafe.Sizeof(uint32(0))))
	MXCFB_SET_UPDATE_SCHEME = godefs.IOW('F', 0x32, int(unsafe.Sizeof(uint32(0))))
	MXCFB_GET_WORK_BUFFER = godefs.IOWR('F', 0x34, int(unsafe.Sizeof(uint64(0))))
	MXCFB_DISABLE_EPDC_ACCESS = godefs.IO('F', 0x35)
	MXCFB_ENABLE_EPDC_ACCESS = godefs.IO('F', 0x36)
}
