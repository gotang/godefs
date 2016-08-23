package kd

/*
#include <linux/kd.h>
#include <linux/types.h>
*/
import "C"


/* 0x4B is 'K', to avoid collision with termios and vt */
const(
	GIO_FONT	uintptr = C.GIO_FONT	/* gets font in expanded form */
	PIO_FONT	uintptr = C.PIO_FONT	/* use font in expanded form */

	GIO_FONTX	uintptr = C.GIO_FONTX	/* get font using struct consolefontdesc */
	PIO_FONTX	uintptr = C.PIO_FONTX	/* set font using struct consolefontdesc */
)

//unsigned short charcount;	/* characters in font (256 or 512) */
//unsigned short charheight;	/* scan lines per character (1-32) */
//char *chardata;		/* font data in expanded form */
type Consolefontdesc C.struct_consolefontdesc

const(
	PIO_FONTRESET uintptr = C.PIO_FONTRESET	/* reset to default font */

	GIO_CMAP uintptr = C.GIO_CMAP	/* gets colour palette on VGA+ */
	PIO_CMAP uintptr = C.PIO_CMAP	/* sets colour palette on VGA+ */

	KIOCSOUND uintptr = C.KIOCSOUND	/* start sound generation (0 for off) */
	KDMKTONE uintptr = C.KDMKTONE	/* generate tone */

	KDGETLED uintptr = C.KDGETLED	/* return current led state */
	KDSETLED uintptr = C.KDSETLED	/* set led state [lights, not flags] */
	LED_SCR	int = C.LED_SCR	/* scroll lock led */
	LED_NUM	int = C.LED_NUM	/* num lock led */
	LED_CAP	int = C.LED_CAP	/* caps lock led */

	KDGKBTYPE uintptr = C.KDGKBTYPE	/* get keyboard type */
	KB_84	int = C.KB_84
	KB_101	int = C.KB_101 	/* this is what we always answer */
	KB_OTHER int = C.KB_OTHER

	KDADDIO	uintptr = C.KDADDIO	/* add i/o port as valid */
	KDDELIO	uintptr = C.KDDELIO	/* del i/o port as valid */
	KDENABIO	uintptr = C.KDENABIO	/* enable i/o to video board */
	KDDISABIO	uintptr = C.KDDISABIO	/* disable i/o to video board */

	KDSETMODE	uintptr = C.KDSETMODE	/* set text/graphics mode */
	KD_TEXT	int = C.KD_TEXT
	KD_GRAPHICS int = C.KD_GRAPHICS
	KD_TEXT0 int = C.KD_TEXT0	/* obsolete */
	KD_TEXT1 int = C.KD_TEXT1	/* obsolete */
	KDGETMODE uintptr = C.KDGETMODE	/* get current mode */

	KDMAPDISP uintptr = C.KDMAPDISP	/* map display into address space */
	KDUNMAPDISP	uintptr = C.KDUNMAPDISP	/* unmap display from address space */

	E_TABSZ	int = C.E_TABSZ
	GIO_SCRNMAP	uintptr = C.GIO_SCRNMAP	/* get screen mapping from kernel */
	PIO_SCRNMAP	uintptr = C.PIO_SCRNMAP	/* put screen mapping table in kernel */
	GIO_UNISCRNMAP  uintptr = C.GIO_UNISCRNMAP	/* get full Unicode screen mapping */
	PIO_UNISCRNMAP  uintptr = C.PIO_UNISCRNMAP /* set full Unicode screen mapping */

	GIO_UNIMAP uintptr = C.GIO_UNIMAP	/* get unicode-to-font mapping from kernel */
)

type Unipair C.struct_unipair

type Unimapdesc C.struct_unimapdesc

const PIO_UNIMAP uintptr = C.PIO_UNIMAP	/* put unicode-to-font mapping in kernel */
const  PIO_UNIMAPCLR uintptr = C.PIO_UNIMAPCLR	/* clear table, possibly advise hash algorithm */

type Unimapinit C.struct_unimapinit

const(
	UNI_DIRECT_BASE uintptr = C.UNI_DIRECT_BASE	/* start of Direct Font Region */
	UNI_DIRECT_MASK uintptr = C.UNI_DIRECT_MASK	/* Direct Font Region bitmask */

	K_RAW	int = C.K_RAW
	K_XLATE	int = C.K_XLATE
	K_MEDIUMRAW int = C.K_MEDIUMRAW
	K_UNICODE int = C.K_UNICODE
	K_OFF	int = C.K_OFF
	KDGKBMODE	uintptr = C.KDGKBMODE	/* gets current keyboard mode */
	KDSKBMODE	uintptr = C.KDSKBMODE	/* sets current keyboard mode */

	K_METABIT int = C.K_METABIT
	K_ESCPREFIX int = C.K_ESCPREFIX
	KDGKBMETA uintptr = C.KDGKBMETA	/* gets meta key handling mode */
	KDSKBMETA uintptr = C.KDSKBMETA	/* sets meta key handling mode */

	K_SCROLLLOCK	int = C.K_SCROLLLOCK
	K_NUMLOCK	int = C.K_NUMLOCK
	K_CAPSLOCK	int = C.K_CAPSLOCK
	KDGKBLED	uintptr = C.KDGKBLED	/* get led flags (not lights) */
	KDSKBLED	uintptr = C.KDSKBLED	/* set led flags (not lights) */
)


type Kbentry C.struct_kbentry

const(
	K_NORMTAB int =	C.K_NORMTAB
	K_SHIFTTAB int = C.K_SHIFTTAB
	K_ALTTAB int = C.K_ALTTAB
	K_ALTSHIFTTAB int = C.K_ALTSHIFTTAB
	KDGKBENT uintptr = C.KDGKBENT	/* gets one entry in translation table */
	KDSKBENT uintptr = C.KDSKBENT	/* sets one entry in translation table */
)

type Kbsentry C.struct_kbsentry

const KDGKBSENT	uintptr = C.KDGKBSENT	/* gets one function key string entry */
const KDSKBSENT	uintptr = C.KDSKBSENT	/* sets one function key string entry */

type Kbdiacr C.struct_kbdiacr

//unsigned int kb_cnt;    /* number of entries in following array */
//struct kbdiacr kbdiacr[256];    /* MAX_DIACR from keyboard.h */
type Kbdiacrs C.struct_kbdiacrs

const KDGKBDIACR     uintptr = C.KDGKBDIACR  /* read kernel accent table */
const KDSKBDIACR      uintptr = C.KDSKBDIACR  /* write kernel accent table */

type Kbdiacruc C.struct_kbdiacruc

//unsigned int kb_cnt;    /* number of entries in following array */
//struct kbdiacruc kbdiacruc[256];    /* MAX_DIACR from keyboard.h */
type Kbdiacrsuc C.struct_kbdiacrsuc

const KDGKBDIACRUC    uintptr = C.KDGKBDIACRUC  /* read kernel accent table - UCS */
const KDSKBDIACRUC    uintptr = C.KDSKBDIACRUC  /* write kernel accent table - UCS */

type Kbkeycode C.struct_kbkeycode

const KDGETKEYCODE	uintptr = C.KDGETKEYCODE	/* read kernel keycode table entry */
const KDSETKEYCODE	uintptr = C.KDSETKEYCODE	/* write kernel keycode table entry */

const KDSIGACCEPT	uintptr = C.KDSIGACCEPT	/* accept kbd generated signals */


//int delay;	/* in msec; <= 0: don't change */
//int period;	/* in msec; <= 0: don't change */
///* earlier this field was misnamed "rate" */
type Kbd_repeat C.struct_kbd_repeat


const KDKBDREP        uintptr = C.KDKBDREP  /* set keyboard delay/repeat rate;
				 * actually used values are returned */

const KDFONTOP	uintptr = C.KDFONTOP	/* font operations */

//
//unsigned int op;	/* operation code KD_FONT_OP_* */
//unsigned int flags;	/* KD_FONT_FLAG_* */
//unsigned int width, height;	/* font size */
//unsigned int charcount;
//unsigned char *data;	/* font data with height fixed to 32 */

type Console_font_op C.struct_console_font_op

//unsigned int width, height;	/* font size */
//unsigned int charcount;
//unsigned char *data;	/* font data with height fixed to 32 */
type Console_font C.struct_console_font

const(
	KD_FONT_OP_SET	int = C.KD_FONT_OP_SET	/* Set font */
	KD_FONT_OP_GET	int = C.KD_FONT_OP_GET	/* Get font */
	KD_FONT_OP_SET_DEFAULT	int = C.KD_FONT_OP_SET_DEFAULT	/* Set font to default, data points to name / NULL */
	KD_FONT_OP_COPY	int = C.KD_FONT_OP_COPY	/* Copy from another console */

	KD_FONT_FLAG_DONT_RECALC int = C.KD_FONT_FLAG_DONT_RECALC	/* Don't recalculate hw charcell size [compat] */
)

/* note: 0x4B00-0x4B4E all have had a value at some time;
   don't reuse for the time being */
/* note: 0x4B60-0x4B6D, 0x4B70-0x4B72 used above */

