package melee

import "math/bits"

type ControllerBits uint32

// Controller Data Bit Flags: xxxx xxxx UDLR UDLR xxxS YXBA xLRZ UDRL
const (
	START_BITS   ControllerBits = 0x1000
	B_BITS                      = 0x0200
	A_BITS                      = 0x0100
	Y_BITS                      = 0x0800
	X_BITS                      = 0x0400
	Z_BITS                      = 0x0010
	L_BITS                      = 0x0040
	R_BITS                      = 0x0020
	D_UP_BITS                   = 0x0008
	D_DOWN_BITS                 = 0x0004
	D_LEFT_BITS                 = 0x0001
	D_RIGHT_BITS                = 0x0002
)

func (c ControllerBits) Has(buttons ControllerBits) bool {
	return c&buttons > 0
}

func (c ControllerBits) Set(buttons ControllerBits) {
	c |= buttons
}

func (c ControllerBits) Clear(buttons ControllerBits) {
	c &= ^(buttons)
}

func (c ControllerBits) Count() int {
	return bits.OnesCount32(uint32(c))
}

//type Controller struct {
//START, Y, X, B, A, L, R, Z bool
//}
