// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs _vt_h.go

package vt

const (
	MIN_NR_CONSOLES		int	= 0x1
	MAX_NR_CONSOLES		int	= 0x3f
	MAX_NR_USER_CONSOLES	int	= 0x3f
)

const VT_OPENQRY = 0x5600

type Vt_mode struct {
	Mode	uint8
	Waitv	uint8
	Relsig	int16
	Acqsig	int16
	Frsig	int16
}

const (
	VT_GETMODE	uintptr	= 0x5601
	VT_SETMODE	uintptr	= 0x5602
	VT_AUTO		int	= 0x0
	VT_PROCESS	int	= 0x1
	VT_ACKACQ	int	= 0x2
)

type Vt_stat struct {
	Active	uint16
	Signal	uint16
	State	uint16
}

const (
	VT_GETSTATE	uintptr	= 0x5603
	VT_SENDSIG	uintptr	= 0x5604

	VT_RELDISP	uintptr	= 0x5605

	VT_ACTIVATE	uintptr	= 0x5606
	VT_WAITACTIVE	uintptr	= 0x5607
	VT_DISALLOCATE	uintptr	= 0x5608
)

type Vt_size struct {
	Rows		uint16
	Cols		uint16
	Scrollsize	uint16
}

const VT_RESIZE uintptr = 0x5609

type Vt_consize struct {
	Rows	uint16
	Cols	uint16
	Vlin	uint16
	Clin	uint16
	Vcol	uint16
	Ccol	uint16
}

const (
	VT_RESIZEX		uintptr	= 0x560a
	VT_LOCKSWITCH		uintptr	= 0x560b
	VT_UNLOCKSWITCH		uintptr	= 0x560c
	VT_GETHIFONTMASK	uintptr	= 0x560d
)

type Vt_event struct {
	Event	uint32
	Oldev	uint32
	Newev	uint32
	Pad	[4]uint32
}

const (
	VT_EVENT_SWITCH		uintptr	= 0x1
	VT_EVENT_BLANK		uintptr	= 0x2
	VT_EVENT_UNBLANK	uintptr	= 0x4
	VT_EVENT_RESIZE		uintptr	= 0x8
	VT_MAX_EVENT		uintptr	= 0xf
)

const VT_WAITEVENT uintptr = 0x560e

type Vt_setactivate struct {
	Console	uint32
	Mode	Vt_mode
}

const VT_SETACTIVATE uintptr = 0x560f
