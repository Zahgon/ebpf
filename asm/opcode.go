package asm

//go:generate go tool stringer -output opcode_string.go -type=Class

type Class uint8

const classMask OpCode = 0x07

const (
	LdClass Class = 0x00

	LdXClass Class = 0x01

	StClass Class = 0x02

	StXClass Class = 0x03

	ALUClass Class = 0x04

	JumpClass Class = 0x05

	Jump32Class Class = 0x06

	ALU64Class Class = 0x07
)

func (cls Class) IsLoad() bool { _ = "STUB: not implemented"; return false }

func (cls Class) IsStore() bool { _ = "STUB: not implemented"; return false }

func (cls Class) isLoadOrStore() bool { _ = "STUB: not implemented"; return false }

func (cls Class) IsALU() bool { _ = "STUB: not implemented"; return false }

func (cls Class) IsJump() bool { _ = "STUB: not implemented"; return false }

func (cls Class) isJumpOrALU() bool { _ = "STUB: not implemented"; return false }

type OpCode uint32

const InvalidOpCode OpCode = 0xffff

func (op OpCode) bpfOpCode() (byte, error) { _ = "STUB: not implemented"; return 0, nil }

func (op OpCode) rawInstructions() int { _ = "STUB: not implemented"; return 0 }

func (op OpCode) IsDWordLoad() bool { _ = "STUB: not implemented"; return false }

func (op OpCode) Class() Class { _ = "STUB: not implemented"; return *new(Class) }

func (op OpCode) Mode() Mode { _ = "STUB: not implemented"; return *new(Mode) }

func (op OpCode) Size() Size { _ = "STUB: not implemented"; return *new(Size) }

func (op OpCode) AtomicOp() AtomicOp { _ = "STUB: not implemented"; return *new(AtomicOp) }

func (op OpCode) Source() Source { _ = "STUB: not implemented"; return *new(Source) }

func (op OpCode) ALUOp() ALUOp { _ = "STUB: not implemented"; return *new(ALUOp) }

func (op OpCode) Endianness() Endianness { _ = "STUB: not implemented"; return *new(Endianness) }

func (op OpCode) JumpOp() JumpOp { _ = "STUB: not implemented"; return *new(JumpOp) }

func (op OpCode) SetMode(mode Mode) OpCode { _ = "STUB: not implemented"; return *new(OpCode) }

func (op OpCode) SetSize(size Size) OpCode { _ = "STUB: not implemented"; return *new(OpCode) }

func (op OpCode) SetAtomicOp(atomic AtomicOp) OpCode {
	_ = "STUB: not implemented"
	return *new(OpCode)
}

func (op OpCode) SetSource(source Source) OpCode { _ = "STUB: not implemented"; return *new(OpCode) }

func (op OpCode) SetALUOp(alu ALUOp) OpCode { _ = "STUB: not implemented"; return *new(OpCode) }

func (op OpCode) SetJumpOp(jump JumpOp) OpCode { _ = "STUB: not implemented"; return *new(OpCode) }

func (op OpCode) String() string { _ = "STUB: not implemented"; return "" }

func valid(value, mask OpCode) bool { _ = "STUB: not implemented"; return false }
