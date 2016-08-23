package vt

/*
#include <linux/vt.h>
*/
import "C"
/*
 * These constants are also useful for user-level apps (e.g., VC
 * resizing).
 */
const(
	MIN_NR_CONSOLES int = C.MIN_NR_CONSOLES       /* must be at least 1 */
	MAX_NR_CONSOLES int = C.MAX_NR_CONSOLES	/* serial lines start at 64 */
	MAX_NR_USER_CONSOLES int = C.MAX_NR_USER_CONSOLES	/* must be root to allocate above this */
)

/* Note: the ioctl VT_GETSTATE does not work for
   consoles 16 and higher (since it returns a short) */

/* 0x56 is 'V', to avoid collision with termios and kd */

const VT_OPENQRY	= C.VT_OPENQRY	/* find available vt */

//char mode;		 vt mode
//char waitv;		 if set, hang on writes if not active
//short relsig;		 signal to raise on release req
//short acqsig;		 signal to raise on acquisition
//short frsig;		 unused (set to 0)
type Vt_mode C.struct_vt_mode

const(
	VT_GETMODE uintptr = C.VT_GETMODE	/* get mode of active vt */
	VT_SETMODE uintptr = C.VT_SETMODE	/* set mode of active vt */
	VT_AUTO	 int = C.VT_AUTO	/* auto vt switching */
	VT_PROCESS int = C.VT_PROCESS	/* process controls switching */
	VT_ACKACQ int = C.VT_ACKACQ	/* acknowledge switch */
)

//unsigned short v_active;	 active vt
//unsigned short v_signal;	 signal to send
//unsigned short v_state;		 vt bitmask
type Vt_stat C.struct_vt_stat

const(
	VT_GETSTATE uintptr = C.VT_GETSTATE	/* get global vt state info */
	VT_SENDSIG uintptr = C.VT_SENDSIG	/* signal to send to bitmask of vts */

	VT_RELDISP uintptr = C.VT_RELDISP	/* release display */

	VT_ACTIVATE uintptr = C.VT_ACTIVATE	/* make vt active */
	VT_WAITACTIVE uintptr = C.VT_WAITACTIVE	/* wait for vt active */
	VT_DISALLOCATE	uintptr = C.VT_DISALLOCATE  /* free memory associated to vt */
)

//unsigned short v_rows;		/* number of rows */
//unsigned short v_cols;		/* number of columns */
//unsigned short v_scrollsize;	/* number of lines of scrollback */
type Vt_size C.struct_vt_sizes

const VT_RESIZE	uintptr = C.VT_RESIZE	/* set kernel's idea of screensize */

//unsigned short v_rows;	/* number of rows */
//unsigned short v_cols;	/* number of columns */
//unsigned short v_vlin;	/* number of pixel rows on screen */
//unsigned short v_clin;	/* number of pixel rows per character */
//unsigned short v_vcol;	/* number of pixel columns on screen */
//unsigned short v_ccol;	/* number of pixel columns per character */
type Vt_consize C.struct_vt_consize

const(
	VT_RESIZEX uintptr = C.VT_RESIZEX   /* set kernel's idea of screensize + more */
	VT_LOCKSWITCH  uintptr = C.VT_LOCKSWITCH  /* disallow vt switching */
	VT_UNLOCKSWITCH uintptr = C.VT_UNLOCKSWITCH  /* allow vt switching */
	VT_GETHIFONTMASK uintptr = C.VT_GETHIFONTMASK  /* return hi font mask */
)

//unsigned int event;
//#define VT_EVENT_SWITCH		0x0001	/* Console switch */
//#define VT_EVENT_BLANK		0x0002	/* Screen blank */
//#define VT_EVENT_UNBLANK	0x0004	/* Screen unblank */
//#define VT_EVENT_RESIZE		0x0008	/* Resize display */
//#define VT_MAX_EVENT		0x000F
//unsigned int oldev;		/* Old console */
//unsigned int newev;		/* New console (if changing) */
//unsigned int pad[4];		/* Padding for expansion */
type Vt_event C.struct_vt_event
const(
	VT_EVENT_SWITCH uintptr = C.VT_EVENT_SWITCH	/* Console switch */
	VT_EVENT_BLANK	uintptr = C.VT_EVENT_BLANK	/* Screen blank */
	VT_EVENT_UNBLANK uintptr = C.VT_EVENT_UNBLANK	/* Screen unblank */
	VT_EVENT_RESIZE uintptr = C.VT_EVENT_RESIZE	/* Resize display */
	VT_MAX_EVENT	uintptr = C.VT_MAX_EVENT
)

const VT_WAITEVENT uintptr = C.VT_WAITEVENT	/* Wait for an event */

type Vt_setactivate C.struct_vt_setactivate


const VT_SETACTIVATE uintptr = C.VT_SETACTIVATE	/* Activate and set the mode of a console */
