package fb

import (
	"github.com/gotang/godefs"
	"unsafe"
)
/*
#include <linux/mxcfb.h>
*/
import "C"

const (
	MXC_DISP_SPEC_DEV = 0
	MXC_DISP_DDC_DEV = 1
)

const(
	MXCFB_REFRESH_OFF =iota
	MXCFB_REFRESH_AUTO
	MXCFB_REFRESH_PARTIAL
)

const(
	 FB_SYNC_OE_LOW_ACT	uintptr=C.FB_SYNC_OE_LOW_ACT
	 FB_SYNC_CLK_LAT_FALL	uintptr=C.FB_SYNC_CLK_LAT_FALL
	 FB_SYNC_DATA_INVERT	uintptr=C.FB_SYNC_DATA_INVERT
	 FB_SYNC_CLK_IDLE_EN	uintptr=C.FB_SYNC_CLK_IDLE_EN
	 FB_SYNC_SHARP_MODE	uintptr=C.FB_SYNC_SHARP_MODE
	 FB_SYNC_SWAP_RGB	uintptr=C.FB_SYNC_SWAP_RGB
	 FB_ACCEL_TRIPLE_FLAG	=C.FB_ACCEL_TRIPLE_FLAG
	 FB_ACCEL_DOUBLE_FLAG	=C.FB_ACCEL_DOUBLE_FLAG
)

type Mxcfb_gbl_alpha C.struct_mxcfb_gbl_alpha

type Mxcfb_loc_alpha C.struct_mxcfb_loc_alpha

type Mxcfb_color_key C.struct_mxcfb_color_key

type Mxcfb_pos C.struct_mxcfb_pos

type Mxcfb_gamma C.struct_mxcfb_gamma

type Mxcfb_rect C.struct_mxcfb_rect

const(
 GRAYSCALE_8BIT			=	0x1
 GRAYSCALE_8BIT_INVERTED	=		0x2
 GRAYSCALE_4BIT                  =        0x3
 GRAYSCALE_4BIT_INVERTED        =         0x4

 AUTO_UPDATE_MODE_REGION_MODE	=	0
 AUTO_UPDATE_MODE_AUTOMATIC_MODE	=	1

 UPDATE_SCHEME_SNAPSHOT		=	0
 UPDATE_SCHEME_QUEUE		=	1
 UPDATE_SCHEME_QUEUE_AND_MERGE	=	2

 UPDATE_MODE_PARTIAL		=	0x0
 UPDATE_MODE_FULL		=	0x1

 WAVEFORM_MODE_AUTO		=	257

 TEMP_USE_AMBIENT		=	0x1000

 EPDC_FLAG_ENABLE_INVERSION	=	0x01
 EPDC_FLAG_FORCE_MONOCHROME	=	0x02
 EPDC_FLAG_USE_CMAP		=	0x04
 EPDC_FLAG_USE_ALT_BUFFER	=	0x100
 EPDC_FLAG_TEST_COLLISION	=	0x200
 EPDC_FLAG_GROUP_UPDATE		=	0x400
 EPDC_FLAG_USE_DITHERING_Y1	=	0x2000
 EPDC_FLAG_USE_DITHERING_Y4	=	0x4000

 FB_POWERDOWN_DISABLE		=	-1
)

//__u32 phys_addr;
//__u32 width;	/* width of entire buffer */
//__u32 height;	/* height of entire buffer */
//struct mxcfb_rect alt_update_region;	/* region within buffer to update */
type Mxcfb_alt_buffer_data C.struct_mxcfb_alt_buffer_data

type Mxcfb_update_data C.struct_mxcfb_update_data

type Mxcfb_update_marker_data C.struct_mxcfb_update_marker_data

/*
 * Structure used to define waveform modes for driver
 * Needed for driver to perform auto-waveform selection
 */
type Mxcfb_waveform_modes C.struct_mxcfb_waveform_modes

/*
 * Structure used to define a 5*3 matrix of parameters for
 * setting IPU DP CSC module related to this framebuffer.
 */
type Mxcfb_csc_matrix C.struct_mxcfb_csc_matrix

var(
	MXCFB_WAIT_FOR_VSYNC int
	MXCFB_SET_GBL_ALPHA    int
	MXCFB_SET_CLR_KEY      int
	MXCFB_SET_OVERLAY_POS  int
	MXCFB_GET_FB_IPU_CHAN 	int
	MXCFB_SET_LOC_ALPHA    int
	MXCFB_SET_LOC_ALP_BUF   int
	MXCFB_SET_GAMMA	       int
	MXCFB_GET_FB_IPU_DI 	int
	MXCFB_GET_DIFMT	       int
	MXCFB_GET_FB_BLANK    int
	MXCFB_SET_DIFMT		int
	MXCFB_CSC_UPDATE	int

	/* IOCTLs for E-ink panel updates */
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
	MXCFB_WAIT_FOR_VSYNC =	godefs.IOW('F', 0x20, int(unsafe.Sizeof(uint32(0))))
	MXCFB_SET_GBL_ALPHA     =godefs.IOW('F', 0x21, int(unsafe.Sizeof(Mxcfb_gbl_alpha{})))
	MXCFB_SET_CLR_KEY       =godefs.IOW('F', 0x22, int(unsafe.Sizeof(Mxcfb_color_key{})))
	MXCFB_SET_OVERLAY_POS   =godefs.IOWR('F', 0x24, int(unsafe.Sizeof(Mxcfb_pos{})))
	MXCFB_GET_FB_IPU_CHAN 	=godefs.IOR('F', 0x25, int(unsafe.Sizeof(uint32(0))))
	MXCFB_SET_LOC_ALPHA     =godefs.IOWR('F', 0x26, int(unsafe.Sizeof(Mxcfb_loc_alpha{})))
	MXCFB_SET_LOC_ALP_BUF    =godefs.IOW('F', 0x27, int(unsafe.Sizeof(uint64(0))))
	MXCFB_SET_GAMMA	       =godefs.IOW('F', 0x28, int(unsafe.Sizeof(Mxcfb_gamma{})))
	MXCFB_GET_FB_IPU_DI 	=godefs.IOR('F', 0x29, int(unsafe.Sizeof(uint32(0))))
	MXCFB_GET_DIFMT	       =godefs.IOR('F', 0x2A, int(unsafe.Sizeof(uint32(0))))
	MXCFB_GET_FB_BLANK     =godefs.IOR('F', 0x2B, int(unsafe.Sizeof(uint32(0))))
	MXCFB_SET_DIFMT		=godefs.IOW('F', 0x2C, int(unsafe.Sizeof(uint32(0))))
	MXCFB_CSC_UPDATE	=godefs.IOW('F', 0x2D, int(unsafe.Sizeof(Mxcfb_csc_matrix{})))

	/* IOCTLs for E-ink panel updates */
	MXCFB_SET_WAVEFORM_MODES	=godefs.IOW('F', 0x2B, int(unsafe.Sizeof(Mxcfb_waveform_modes{})))
	MXCFB_SET_TEMPERATURE		=godefs.IOW('F', 0x2C, int(unsafe.Sizeof(uint32(0))))
	MXCFB_SET_AUTO_UPDATE_MODE	=godefs.IOW('F', 0x2D, int(unsafe.Sizeof(uint32(0))))
	MXCFB_SEND_UPDATE		=godefs.IOW('F', 0x2E, int(unsafe.Sizeof(Mxcfb_update_marker_data{})))
	MXCFB_SET_PWRDOWN_DELAY		=godefs.IOW('F', 0x30, int(unsafe.Sizeof(uint32(0))))
	MXCFB_GET_PWRDOWN_DELAY		=godefs.IOR('F', 0x31, int(unsafe.Sizeof(uint32(0))))
	MXCFB_SET_UPDATE_SCHEME		=godefs.IOW('F', 0x32, int(unsafe.Sizeof(uint32(0))))
	MXCFB_GET_WORK_BUFFER		=godefs.IOWR('F', 0x34, int(unsafe.Sizeof(uint64(0))))
	MXCFB_DISABLE_EPDC_ACCESS	=godefs.IO('F', 0x35)
	MXCFB_ENABLE_EPDC_ACCESS	=godefs.IO('F', 0x36)
}
