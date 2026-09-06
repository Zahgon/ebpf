package asm

type Register uint8

const R0 Register = 0

const (
	R1 Register = R0 + 1 + iota
	R2
	R3
	R4
	R5
)

const (
	R6 Register = R5 + 1 + iota
	R7
	R8
	R9
)

const (
	R10 Register = R9 + 1
	RFP          = R10
)

const (
	PseudoMapFD     = R1
	PseudoMapValue  = R2
	PseudoBtfId     = R3
	PseudoCall      = R1
	PseudoFunc      = R4
	PseudoKfuncCall = R2
	PseudoMayGoto   = R0
)

func (r Register) String() string { _ = "STUB: not implemented"; return "" }
