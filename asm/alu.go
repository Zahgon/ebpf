package asm

//go:generate go tool stringer -output alu_string.go -type=Source,Endianness,ALUOp

type Source uint16

const sourceMask OpCode = 0x0008

const (
	InvalidSource Source = 0xffff

	ImmSource Source = 0x0000

	RegSource Source = 0x0008
)

type Endianness uint8

const endianMask = sourceMask

const (
	InvalidEndian Endianness = 0xff

	LE Endianness = 0x00

	BE Endianness = 0x08
)

type ALUOp uint16

const aluMask OpCode = 0x3ff0

const (
	InvalidALUOp ALUOp = 0xffff

	Add ALUOp = 0x0000

	Sub ALUOp = 0x0010

	Mul ALUOp = 0x0020

	Div ALUOp = 0x0030

	SDiv ALUOp = Div + 0x0100

	Or ALUOp = 0x0040

	And ALUOp = 0x0050

	LSh ALUOp = 0x0060

	RSh ALUOp = 0x0070

	Neg ALUOp = 0x0080

	Mod ALUOp = 0x0090

	SMod ALUOp = Mod + 0x0100

	Xor ALUOp = 0x00a0

	Mov ALUOp = 0x00b0

	MovSX8 ALUOp = Mov + 0x0100

	MovSX16 ALUOp = Mov + 0x0200

	MovSX32 ALUOp = Mov + 0x0300

	ArSh ALUOp = 0x00c0

	Swap ALUOp = 0x00d0
)

func HostTo(endian Endianness, dst Register, size Size) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}

func BSwap(dst Register, size Size) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}

func (op ALUOp) Op(source Source) OpCode { _ = "STUB: not implemented"; return *new(OpCode) }

func (op ALUOp) Reg(dst, src Register) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}

func (op ALUOp) Imm(dst Register, value int32) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}

func (op ALUOp) Op32(source Source) OpCode { _ = "STUB: not implemented"; return *new(OpCode) }

func (op ALUOp) Reg32(dst, src Register) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}

func (op ALUOp) Imm32(dst Register, value int32) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}
