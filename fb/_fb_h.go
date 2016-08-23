package fb

import (
 "github.com/gotang/godefs"
 "unsafe"
)
/*
#include <linux/fb.h>
*/
import "C"

/* Definitions of frame buffers						*/

var(
	FBIO_CURSOR int
	FBIOGET_VBLANK int
	FBIO_WAITFORVSYNC int
)
func init() {
  FBIO_CURSOR      =      godefs.IOWR('F', 0x08, int(unsafe.Sizeof(Fb_cursor{})))
  FBIOGET_VBLANK	 =	godefs.IOR('F', 0x12, int(unsafe.Sizeof(Fb_vblank{})))
  FBIO_WAITFORVSYNC	= 	godefs.IOW('F', 0x20, int(unsafe.Sizeof(uint32(0))))
}
const FB_MAX	int = C.FB_MAX	/* sufficient for now */

/* ioctls
   0x46 is 'F'								*/
const FBIOGET_VSCREENINFO	uintptr = C.FBIOGET_VSCREENINFO
const FBIOPUT_VSCREENINFO	uintptr = C.FBIOPUT_VSCREENINFO
const FBIOGET_FSCREENINFO	uintptr = C.FBIOGET_FSCREENINFO
const FBIOGETCMAP		uintptr = C.FBIOGETCMAP
const FBIOPUTCMAP		uintptr = C.FBIOPUTCMAP
const FBIOPAN_DISPLAY		uintptr = C.FBIOPAN_DISPLAY
//const FBIO_CURSOR      =      _IOWR('F', 0x08, struct fb_cursor)
/* 0x4607-0x460B are defined below */
/* #define FBIOGET_MONITORSPEC	0x460C */
/* #define FBIOPUT_MONITORSPEC	0x460D */
/* #define FBIOSWITCH_MONIBIT	0x460E */
const FBIOGET_CON2FBMAP		uintptr = C.FBIOGET_CON2FBMAP
const FBIOPUT_CON2FBMAP	        uintptr = C.FBIOPUT_CON2FBMAP
const FBIOBLANK		uintptr = C.FBIOBLANK		/* arg: 0 or vesa level + 1 */

const FBIO_ALLOC             uintptr = C.FBIO_ALLOC
const FBIO_FREE              uintptr = C.FBIO_FREE
const FBIOGET_GLYPH           uintptr = C.FBIOGET_GLYPH
const FBIOGET_HWCINFO         uintptr = C.FBIOGET_HWCINFO
const FBIOPUT_MODEINFO        uintptr = C.FBIOPUT_MODEINFO
const FBIOGET_DISPINFO        uintptr = C.FBIOGET_DISPINFO


const (
	 FB_TYPE_PACKED_PIXELS	=	iota	/* Packed Pixels	*/
	 FB_TYPE_PLANES				/* Non interleaved planes */
	 FB_TYPE_INTERLEAVED_PLANES		/* Interleaved planes	*/
	 FB_TYPE_TEXT				/* Text/attributes	*/
	 FB_TYPE_VGA_PLANES			/* EGA/VGA planes	*/
	 FB_TYPE_FOURCC				/* Type identified by a V4L2 FOURCC */
)

const (
 FB_AUX_TEXT_MDA		= 0	/* Monochrome text */
 FB_AUX_TEXT_CGA		=1	/* CGA/EGA/VGA Color text */
 FB_AUX_TEXT_S3_MMIO		=2/* S3 MMIO fasttext */
 FB_AUX_TEXT_MGA_STEP16		=3/* MGA Millenium I: text, attr, 14 reserved bytes */
 FB_AUX_TEXT_MGA_STEP8		=4/* other MGAs:      text, attr,  6 reserved bytes */
 FB_AUX_TEXT_SVGA_GROUP		=8/* 8-15: SVGA tileblit compatible modes */
 FB_AUX_TEXT_SVGA_MASK		=7/* lower three bits says step */
 FB_AUX_TEXT_SVGA_STEP2		=8/* SVGA text mode:  text, attr */
 FB_AUX_TEXT_SVGA_STEP4		=9/* SVGA text mode:  text, attr,  2 reserved bytes */
 FB_AUX_TEXT_SVGA_STEP8		=10/* SVGA text mode:  text, attr,  6 reserved bytes */
 FB_AUX_TEXT_SVGA_STEP16	=11	/* SVGA text mode:  text, attr, 14 reserved bytes */
 FB_AUX_TEXT_SVGA_LAST		=15/* reserved up to 15 */
)

const(
 FB_AUX_VGA_PLANES_VGA4		= iota	/* 16 color planes (EGA/VGA) */
 FB_AUX_VGA_PLANES_CFB4			/* CFB4 in planes (VGA) */
 FB_AUX_VGA_PLANES_CFB8			/* CFB8 in planes (VGA) */
)

const (
 FB_VISUAL_MONO01		= iota	/* Monochr. 1=Black 0=White */
 FB_VISUAL_MONO10			/* Monochr. 1=White 0=Black */
 FB_VISUAL_TRUECOLOR			/* True color	*/
 FB_VISUAL_PSEUDOCOLOR			/* Pseudo color (like atari) */
 FB_VISUAL_DIRECTCOLOR			/* Direct color */
 FB_VISUAL_STATIC_PSEUDOCOLOR		/* Pseudo color readonly */
 FB_VISUAL_FOURCC			/* Visual identified by a V4L2 FOURCC */
)

const(
 FB_ACCEL_NONE		= 0	/* no hardware accelerator	*/
 FB_ACCEL_ATARIBLITT	=1	/* Atari Blitter		*/
 FB_ACCEL_AMIGABLITT	=2	/* Amiga Blitter                */
 FB_ACCEL_S3_TRIO64	=3	/* Cybervision64 (S3 Trio64)    */
 FB_ACCEL_NCR_77C32BLT	=4	/* RetinaZ3 (NCR 77C32BLT)      */
 FB_ACCEL_S3_VIRGE	=5	/* Cybervision64/3D (S3 ViRGE)	*/
 FB_ACCEL_ATI_MACH64GX	=6	/* ATI Mach 64GX family		*/
 FB_ACCEL_DEC_TGA	=7	/* DEC 21030 TGA		*/
 FB_ACCEL_ATI_MACH64CT	=8	/* ATI Mach 64CT family		*/
 FB_ACCEL_ATI_MACH64VT	=9	/* ATI Mach 64CT family VT class */
 FB_ACCEL_ATI_MACH64GT	=10	/* ATI Mach 64CT family GT class */
 FB_ACCEL_SUN_CREATOR	=11	/* Sun Creator/Creator3D	*/
 FB_ACCEL_SUN_CGSIX	=12	/* Sun cg6			*/
 FB_ACCEL_SUN_LEO	=13	/* Sun leo/zx			*/
 FB_ACCEL_IMS_TWINTURBO	=14	/* IMS Twin Turbo		*/
 FB_ACCEL_3DLABS_PERMEDIA2 =15	/* 3Dlabs Permedia 2		*/
 FB_ACCEL_MATROX_MGA2064W =16	/* Matrox MGA2064W (Millenium)	*/
 FB_ACCEL_MATROX_MGA1064SG =17	/* Matrox MGA1064SG (Mystique)	*/
 FB_ACCEL_MATROX_MGA2164W =18	/* Matrox MGA2164W (Millenium II) */
 FB_ACCEL_MATROX_MGA2164W_AGP =19	/* Matrox MGA2164W (Millenium II) */
 FB_ACCEL_MATROX_MGAG100	=20	/* Matrox G100 (Productiva G100) */
 FB_ACCEL_MATROX_MGAG200	=21	/* Matrox G200 (Myst, Mill, ...) */
 FB_ACCEL_SUN_CG14	=22	/* Sun cgfourteen		 */
 FB_ACCEL_SUN_BWTWO	=23	/* Sun bwtwo			*/
 FB_ACCEL_SUN_CGTHREE	=24	/* Sun cgthree			*/
 FB_ACCEL_SUN_TCX	=25	/* Sun tcx			*/
 FB_ACCEL_MATROX_MGAG400 =26		/* Matrox G400			*/
 FB_ACCEL_NV3		=27	/* nVidia RIVA 128              */
 FB_ACCEL_NV4		= 28	/* nVidia RIVA TNT		*/
 FB_ACCEL_NV5		= 29	/* nVidia RIVA TNT2		*/
 FB_ACCEL_CT_6555x	= 30	/* C&T 6555x			*/
 FB_ACCEL_3DFX_BANSHEE	= 31	/* 3Dfx Banshee			*/
 FB_ACCEL_ATI_RAGE128	= 32	/* ATI Rage128 family		*/
 FB_ACCEL_IGS_CYBER2000	= 33	/* CyberPro 2000		*/
 FB_ACCEL_IGS_CYBER2010	= 34	/* CyberPro 2010		*/
 FB_ACCEL_IGS_CYBER5000	= 35	/* CyberPro 5000		*/
 FB_ACCEL_SIS_GLAMOUR    = 36	/* SiS 300/630/540              */
 FB_ACCEL_3DLABS_PERMEDIA3 = 37	/* 3Dlabs Permedia 3		*/
 FB_ACCEL_ATI_RADEON	= 38	/* ATI Radeon family		*/
 FB_ACCEL_I810      = 39           /* Intel 810/815                */
 FB_ACCEL_SIS_GLAMOUR_2  = 40 	/* SiS 315, 650, 740		*/
 FB_ACCEL_SIS_XABRE      = 41	/* SiS 330 ("Xabre")		*/
 FB_ACCEL_I830              =42   /* Intel 830M/845G/85x/865G     */
 FB_ACCEL_NV_10              =43  /* nVidia Arch 10               */
 FB_ACCEL_NV_20      =44          /* nVidia Arch 20               */
 FB_ACCEL_NV_30         =45       /* nVidia Arch 30               */
 FB_ACCEL_NV_40           =46     /* nVidia Arch 40               */
 FB_ACCEL_XGI_VOLARI_V	=47	/* XGI Volari V3XT, V5, V8      */
 FB_ACCEL_XGI_VOLARI_Z	=48	/* XGI Volari Z7                */
 FB_ACCEL_OMAP1610	=49	/* TI OMAP16xx                  */
 FB_ACCEL_TRIDENT_TGUI	=50	/* Trident TGUI			*/
 FB_ACCEL_TRIDENT_3DIMAGE =51	/* Trident 3DImage		*/
 FB_ACCEL_TRIDENT_BLADE3D =52	/* Trident Blade3D		*/
 FB_ACCEL_TRIDENT_BLADEXP =53	/* Trident BladeXP		*/
 FB_ACCEL_CIRRUS_ALPINE   =53	/* Cirrus Logic 543x/544x/5480	*/
 FB_ACCEL_NEOMAGIC_NM2070 =90	/* NeoMagic NM2070              */
 FB_ACCEL_NEOMAGIC_NM2090 =91	/* NeoMagic NM2090              */
 FB_ACCEL_NEOMAGIC_NM2093 =92	/* NeoMagic NM2093              */
 FB_ACCEL_NEOMAGIC_NM2097 =93	/* NeoMagic NM2097              */
 FB_ACCEL_NEOMAGIC_NM2160 =94	/* NeoMagic NM2160              */
 FB_ACCEL_NEOMAGIC_NM2200 =95	/* NeoMagic NM2200              */
 FB_ACCEL_NEOMAGIC_NM2230 =96	/* NeoMagic NM2230              */
 FB_ACCEL_NEOMAGIC_NM2360 =97	/* NeoMagic NM2360              */
 FB_ACCEL_NEOMAGIC_NM2380 =98	/* NeoMagic NM2380              */
 FB_ACCEL_PXA3XX	=99	 	/* PXA3xx			*/
)

const(
 FB_ACCEL_SAVAGE4        =0x80	/* S3 Savage4                   */
 FB_ACCEL_SAVAGE3D       =0x81	/* S3 Savage3D                  */
 FB_ACCEL_SAVAGE3D_MV    =0x82	/* S3 Savage3D-MV               */
 FB_ACCEL_SAVAGE2000     =0x83	/* S3 Savage2000                */
 FB_ACCEL_SAVAGE_MX_MV   =0x84	/* S3 Savage/MX-MV              */
 FB_ACCEL_SAVAGE_MX      =0x85	/* S3 Savage/MX                 */
 FB_ACCEL_SAVAGE_IX_MV   =0x86	/* S3 Savage/IX-MV              */
 FB_ACCEL_SAVAGE_IX      =0x87	/* S3 Savage/IX                 */
 FB_ACCEL_PROSAVAGE_PM   =0x88	/* S3 ProSavage PM133           */
 FB_ACCEL_PROSAVAGE_KM   =0x89	/* S3 ProSavage KM133           */
 FB_ACCEL_S3TWISTER_P    =0x8a	/* S3 Twister                   */
 FB_ACCEL_S3TWISTER_K    =0x8b	/* S3 TwisterK                  */
 FB_ACCEL_SUPERSAVAGE    =0x8c    /* S3 Supersavage               */
 FB_ACCEL_PROSAVAGE_DDR  =0x8d	/* S3 ProSavage DDR             */
 FB_ACCEL_PROSAVAGE_DDRK =0x8e	/* S3 ProSavage DDR-K           */
 FB_ACCEL_PUV3_UNIGFX =	0xa0	/* PKUnity-v3 Unigfx		*/
)

const  FB_CAP_FOURCC	=	1	/* Device supports FOURCC-based formats */

//char id[16];			/* identification string eg "TT Builtin" */
//unsigned long smem_start;	/* Start of frame buffer mem */
///* (physical address) */
//__u32 smem_len;			/* Length of frame buffer mem */
//__u32 type;			/* see FB_TYPE_*		*/
//__u32 type_aux;			/* Interleave for interleaved Planes */
//__u32 visual;			/* see FB_VISUAL_*		*/
//__u16 xpanstep;			/* zero if no hardware panning  */
//__u16 ypanstep;			/* zero if no hardware panning  */
//__u16 ywrapstep;		/* zero if no hardware ywrap    */
//__u32 line_length;		/* length of a line in bytes    */
//unsigned long mmio_start;	/* Start of Memory Mapped I/O   */
///* (physical address) */
//__u32 mmio_len;			/* Length of Memory Mapped I/O  */
//__u32 accel;			/* Indicate to driver which	*/
///*  specific chip/card we have	*/
//__u16 capabilities;		/* see FB_CAP_*			*/
//__u16 reserved[2];		/* Reserved for future compatibility */
type Fb_fix_screeninfo C.struct_fb_fix_screeninfo


/* Interpretation of offset for color fields: All offsets are from the right,
 * inside a "pixel" value, which is exactly 'bits_per_pixel' wide (means: you
 * can use the offset as right argument to <<). A pixel afterwards is a bit
 * stream and is written to video memory as that unmodified.
 *
 * For pseudocolor: offset and length should be the same for all color
 * components. Offset specifies the position of the least significant bit
 * of the pallette index in a pixel value. Length indicates the number
 * of available palette entries (i.e. # of entries = 1 << length).
 */

//__u32 offset;			/* beginning of bitfield	*/
//__u32 length;			/* length of bitfield		*/
//__u32 msb_right;		/* != 0 : Most significant bit is */
///* right */
type Fb_bitfield C.struct_fb_bitfield

const (
 FB_NONSTD_HAM	=	1	/* Hold-And-Modify (HAM)        */
 FB_NONSTD_REV_PIX_IN_B	=2	/* order of pixels in each byte is reversed */
)

const(
 FB_ACTIVATE_NOW	=	0	/* set values immediately (or vbl)*/
 FB_ACTIVATE_NXTOPEN	=1	/* activate on next open	*/
 FB_ACTIVATE_TEST	=2	/* don't set, round up impossible */
 FB_ACTIVATE_MASK       =15
/* values			*/
 FB_ACTIVATE_VBL	 =      16	/* activate values on next vbl  */
 FB_CHANGE_CMAP_VBL     =32	/* change colormap on vbl	*/
 FB_ACTIVATE_ALL	=       64	/* change all VCs on this fb	*/
 FB_ACTIVATE_FORCE     =128	/* force apply even when no change*/
 FB_ACTIVATE_INV_MODE  =256       /* invalidate videomode */

 FB_ACCELF_TEXT	=	1	/* (OBSOLETE) see fb_info.flags and vc_mode */

)

const(
    FB_SYNC_HOR_HIGH_ACT	=1	/* horizontal sync high active	*/
    FB_SYNC_VERT_HIGH_ACT	=2	/* vertical sync high active	*/
    FB_SYNC_EXT		=4	/* external sync		*/
    FB_SYNC_COMP_HIGH_ACT	=8	/* composite sync high active   */
    FB_SYNC_BROADCAST	=16	/* broadcast video timings      */
   /* vtotal = 144d/288n/576i => PAL  */
   /* vtotal = 121d/242n/484i => NTSC */
    FB_SYNC_ON_GREEN	=32	/* sync on green */

    FB_VMODE_NONINTERLACED  =0	/* non interlaced */
    FB_VMODE_INTERLACED	=1	/* interlaced	*/
    FB_VMODE_DOUBLE	=	2	/* double scan */
    FB_VMODE_ODD_FLD_FIRST	=4	/* interlaced: top line first */
    FB_VMODE_MASK		=255

    FB_VMODE_YWRAP		=256	/* ywrap instead of panning     */
    FB_VMODE_SMOOTH_XPAN	=512	/* smooth xpan possible (internally used) */
    FB_VMODE_CONUPDATE	=512	/* don't update x/yoffset	*/

   /*
    * Display rotation support
    */
    FB_ROTATE_UR   =   0
    FB_ROTATE_CW    =  1
    FB_ROTATE_UD    =  2
    FB_ROTATE_CCW    = 3

    //PICOS2KHZ(a) (1000000000UL/(a))
    //KHZ2PICOS(a) (1000000000UL/(a))

)

//__u32 xres;			/* visible resolution		*/
//__u32 yres;
//__u32 xres_virtual;		/* virtual resolution		*/
//__u32 yres_virtual;
//__u32 xoffset;			/* offset from virtual to visible */
//__u32 yoffset;			/* resolution			*/
//
//__u32 bits_per_pixel;		/* guess what			*/
//__u32 grayscale;		/* 0 = color, 1 = grayscale,	*/
///* >1 = FOURCC			*/
//struct fb_bitfield red;		/* bitfield in fb mem if true color, */
//struct fb_bitfield green;	/* else only length is significant */
//struct fb_bitfield blue;
//struct fb_bitfield transp;	/* transparency			*/
//
//__u32 nonstd;			/* != 0 Non standard pixel format */
//
//__u32 activate;			/* see FB_ACTIVATE_*		*/
//
//__u32 height;			/* height of picture in mm    */
//__u32 width;			/* width of picture in mm     */
//
//__u32 accel_flags;		/* (OBSOLETE) see fb_info.flags */
//
///* Timing: All values in pixclocks, except pixclock (of course) */
//__u32 pixclock;			/* pixel clock in ps (pico seconds) */
//__u32 left_margin;		/* time from sync to picture	*/
//__u32 right_margin;		/* time from picture to sync	*/
//__u32 upper_margin;		/* time from sync to picture	*/
//__u32 lower_margin;
//__u32 hsync_len;		/* length of horizontal sync	*/
//__u32 vsync_len;		/* length of vertical sync	*/
//__u32 sync;			/* see FB_SYNC_*		*/
//__u32 vmode;			/* see FB_VMODE_*		*/
//__u32 rotate;			/* angle we rotate counter clockwise */
//__u32 colorspace;		/* colorspace for FOURCC-based modes */
//__u32 reserved[4];		/* Reserved for future compatibility */
type Fb_var_screeninfo C.struct_fb_var_screeninfo


//__u32 start;			/* First entry	*/
//__u32 len;			/* Number of entries */
//__u16 *red;			/* Red values	*/
//__u16 *green;
//__u16 *blue;
//__u16 *transp;			/* transparency, can be NULL */
type Fb_cmap C.struct_fb_cmap

type Fb_con2fbmap C.struct_fb_con2fbmap

/* VESA Blanking Levels */
const(
 VESA_NO_BLANKING       = iota
 VESA_VSYNC_SUSPEND
 VESA_HSYNC_SUSPEND
 VESA_POWERDOWN
)

const(
 /* screen: unblanked, hsync: on,  vsync: on */
 FB_BLANK_UNBLANK       = VESA_NO_BLANKING

 /* screen: blanked,   hsync: on,  vsync: on */
 FB_BLANK_NORMAL        = VESA_NO_BLANKING + 1

/* screen: blanked,   hsync: on,  vsync: off */
FB_BLANK_VSYNC_SUSPEND = VESA_VSYNC_SUSPEND + 1

/* screen: blanked,   hsync: off, vsync: on */
FB_BLANK_HSYNC_SUSPEND = VESA_HSYNC_SUSPEND + 1

/* screen: blanked,   hsync: off, vsync: off */
FB_BLANK_POWERDOWN     = VESA_POWERDOWN + 1
)

const(
 FB_VBLANK_VBLANKING	=0x001	/* currently in a vertical blank */
 FB_VBLANK_HBLANKING	=0x002	/* currently in a horizontal blank */
 FB_VBLANK_HAVE_VBLANK	=0x004	/* vertical blanks can be detected */
 FB_VBLANK_HAVE_HBLANK	=0x008	/* horizontal blanks can be detected */
 FB_VBLANK_HAVE_COUNT	=0x010	/* global retrace counter is available */
 FB_VBLANK_HAVE_VCOUNT	=0x020	/* the vcount field is valid */
 FB_VBLANK_HAVE_HCOUNT	=0x040	/* the hcount field is valid */
 FB_VBLANK_VSYNCING	=0x080	/* currently in a vsync */
 FB_VBLANK_HAVE_VSYNC	=0x100	/* verical syncs can be detected */
)

//__u32 flags;			/* FB_VBLANK flags */
//__u32 count;			/* counter of retraces since boot */
//__u32 vcount;			/* current scanline position */
//__u32 hcount;			/* current scandot position */
//__u32 reserved[4];		/* reserved for future compatibility */
type Fb_vblank C.struct_fb_vblank

const(
/* Internal HW accel */
 ROP_COPY =0
 ROP_XOR  =1
)

type Fb_copyarea C.struct_fb_copyarea

type Fb_fillrect C.struct_fb_fillrect

//__u32 dx;		/* Where to place image */
//__u32 dy;
//__u32 width;		/* Size of image */
//__u32 height;
//__u32 fg_color;		/* Only used when a mono bitmap */
//__u32 bg_color;
//__u8  depth;		/* Depth of the image */
//const char *data;	/* Pointer to image data */
//struct fb_cmap cmap;	/* color map info */
type Fb_image C.struct_fb_image

const(
/*
* hardware cursor control
*/

 FB_CUR_SETIMAGE =0x01
 FB_CUR_SETPOS   =0x02
 FB_CUR_SETHOT   =0x04
 FB_CUR_SETCMAP  =0x08
 FB_CUR_SETSHAPE =0x10
 FB_CUR_SETSIZE	=0x20
 FB_CUR_SETALL  = 0xFF
)

type Fbcurpos C.struct_fbcurpos

//__u16 set;		/* what to set */
//__u16 enable;		/* cursor on/off */
//__u16 rop;		/* bitop operation */
//const char *mask;	/* cursor mask bits */
//struct fbcurpos hot;	/* cursor hot spot */
//struct fb_image	image;	/* Cursor image */
type Fb_cursor C.struct_fb_cursor

const(
//#ifdef CONFIG_FB_BACKLIGHT
/* Settings for the generic backlight code */
 FB_BACKLIGHT_LEVELS	= 128
 FB_BACKLIGHT_MAX	= 0xFF
)
