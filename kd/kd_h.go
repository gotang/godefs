// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs _kd_h.go

package kd

const (
	GIO_FONT	uintptr	= 0x4b60
	PIO_FONT	uintptr	= 0x4b61

	GIO_FONTX	uintptr	= 0x4b6b
	PIO_FONTX	uintptr	= 0x4b6c
)

type Consolefontdesc struct {
	Charcount	uint16
	Charheight	uint16
	Chardata	*uint8
}

const (
	PIO_FONTRESET	uintptr	= 0x4b6d

	GIO_CMAP	uintptr	= 0x4b70
	PIO_CMAP	uintptr	= 0x4b71

	KIOCSOUND	uintptr	= 0x4b2f
	KDMKTONE	uintptr	= 0x4b30

	KDGETLED	uintptr	= 0x4b31
	KDSETLED	uintptr	= 0x4b32
	LED_SCR		int	= 0x1
	LED_NUM		int	= 0x2
	LED_CAP		int	= 0x4

	KDGKBTYPE	uintptr	= 0x4b33
	KB_84		int	= 0x1
	KB_101		int	= 0x2
	KB_OTHER	int	= 0x3

	KDADDIO		uintptr	= 0x4b34
	KDDELIO		uintptr	= 0x4b35
	KDENABIO	uintptr	= 0x4b36
	KDDISABIO	uintptr	= 0x4b37

	KDSETMODE	uintptr	= 0x4b3a
	KD_TEXT		int	= 0x0
	KD_GRAPHICS	int	= 0x1
	KD_TEXT0	int	= 0x2
	KD_TEXT1	int	= 0x3
	KDGETMODE	uintptr	= 0x4b3b

	KDMAPDISP	uintptr	= 0x4b3c
	KDUNMAPDISP	uintptr	= 0x4b3d

	E_TABSZ		int	= 0x100
	GIO_SCRNMAP	uintptr	= 0x4b40
	PIO_SCRNMAP	uintptr	= 0x4b41
	GIO_UNISCRNMAP	uintptr	= 0x4b69
	PIO_UNISCRNMAP	uintptr	= 0x4b6a

	GIO_UNIMAP	uintptr	= 0x4b66
)

type Unipair struct {
	Unicode	uint16
	Fontpos	uint16
}

type Unimapdesc struct {
	Ct		uint16
	Pad_cgo_0	[2]byte
	Entries		*Unipair
}

const PIO_UNIMAP uintptr = 0x4b67
const PIO_UNIMAPCLR uintptr = 0x4b68

type Unimapinit struct {
	Hashsize	uint16
	Hashstep	uint16
	Hashlevel	uint16
}

const (
	UNI_DIRECT_BASE	uintptr	= 0xf000
	UNI_DIRECT_MASK	uintptr	= 0x1ff

	K_RAW		int	= 0x0
	K_XLATE		int	= 0x1
	K_MEDIUMRAW	int	= 0x2
	K_UNICODE	int	= 0x3
	K_OFF		int	= 0x4
	KDGKBMODE	uintptr	= 0x4b44
	KDSKBMODE	uintptr	= 0x4b45

	K_METABIT	int	= 0x3
	K_ESCPREFIX	int	= 0x4
	KDGKBMETA	uintptr	= 0x4b62
	KDSKBMETA	uintptr	= 0x4b63

	K_SCROLLLOCK	int	= 0x1
	K_NUMLOCK	int	= 0x2
	K_CAPSLOCK	int	= 0x4
	KDGKBLED	uintptr	= 0x4b64
	KDSKBLED	uintptr	= 0x4b65
)

type Kbentry struct {
	Table	uint8
	Index	uint8
	Value	uint16
}

const (
	K_NORMTAB	int	= 0x0
	K_SHIFTTAB	int	= 0x1
	K_ALTTAB	int	= 0x2
	K_ALTSHIFTTAB	int	= 0x3
	KDGKBENT	uintptr	= 0x4b46
	KDSKBENT	uintptr	= 0x4b47
)

type Kbsentry struct {
	Func	uint8
	String	[512]uint8
}

const KDGKBSENT uintptr = 0x4b48
const KDSKBSENT uintptr = 0x4b49

type Kbdiacr struct {
	Diacr	uint8
	Base	uint8
	Result	uint8
}

type Kbdiacrs struct {
	Cnt	uint32
	Kbdiacr	[256]Kbdiacr
}

const KDGKBDIACR uintptr = 0x4b4a
const KDSKBDIACR uintptr = 0x4b4b

type Kbdiacruc struct {
	Diacr	uint32
	Base	uint32
	Result	uint32
}

type Kbdiacrsuc struct {
	Cnt		uint32
	Kbdiacruc	[256]Kbdiacruc
}

const KDGKBDIACRUC uintptr = 0x4bfa
const KDSKBDIACRUC uintptr = 0x4bfb

type Kbkeycode struct {
	Scancode	uint32
	Keycode		uint32
}

const KDGETKEYCODE uintptr = 0x4b4c
const KDSETKEYCODE uintptr = 0x4b4d

const KDSIGACCEPT uintptr = 0x4b4e

type Kbd_repeat struct {
	Delay	int32
	Period	int32
}

const KDKBDREP uintptr = 0x4b52

const KDFONTOP uintptr = 0x4b72

type Console_font_op struct {
	Op		uint32
	Flags		uint32
	Width		uint32
	Height		uint32
	Charcount	uint32
	Data		*uint8
}

type Console_font struct {
	Width		uint32
	Height		uint32
	Charcount	uint32
	Data		*uint8
}

const (
	KD_FONT_OP_SET		int	= 0x0
	KD_FONT_OP_GET		int	= 0x1
	KD_FONT_OP_SET_DEFAULT	int	= 0x2
	KD_FONT_OP_COPY		int	= 0x3

	KD_FONT_FLAG_DONT_RECALC	int	= 0x1
)
