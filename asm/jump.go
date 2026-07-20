package asm

//go:generate go tool stringer -output jump_string.go -type=JumpOp

type JumpOp uint8

const jumpMask OpCode = 0xf0

const (
	InvalidJumpOp JumpOp = 0xff

	Ja JumpOp = 0x00

	JEq JumpOp = 0x10

	JGT JumpOp = 0x20

	JGE JumpOp = 0x30

	JSet JumpOp = 0x40

	JNE JumpOp = 0x50

	JSGT JumpOp = 0x60

	JSGE JumpOp = 0x70

	Call JumpOp = 0x80

	Exit JumpOp = 0x90

	JLT JumpOp = 0xa0

	JLE JumpOp = 0xb0

	JSLT JumpOp = 0xc0

	JSLE JumpOp = 0xd0

	JCOND JumpOp = 0xe0
)

func Return() Instruction { _ = "STUB: not implemented"; return *new(Instruction) }

func (op JumpOp) Op(source Source) OpCode { _ = "STUB: not implemented"; return *new(OpCode) }

func (op JumpOp) Imm(dst Register, value int32, label string) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}

func (op JumpOp) Imm32(dst Register, value int32, label string) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}

func (op JumpOp) Reg(dst, src Register, label string) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}

func (op JumpOp) Reg32(dst, src Register, label string) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}

func (op JumpOp) opCode(class Class, source Source) OpCode {
	_ = "STUB: not implemented"
	return *new(OpCode)
}

func LongJump(label string) Instruction { _ = "STUB: not implemented"; return *new(Instruction) }

func (op JumpOp) Label(label string) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}
