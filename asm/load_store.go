package asm

//go:generate go tool stringer -output load_store_string.go -type=Mode,Size

type Mode uint8

const modeMask OpCode = 0xe0

const (
	InvalidMode Mode = 0xff

	ImmMode Mode = 0x00

	AbsMode Mode = 0x20

	IndMode Mode = 0x40

	MemMode Mode = 0x60

	MemSXMode Mode = 0x80

	AtomicMode Mode = 0xc0
)

const atomicMask OpCode = 0x0001_ff00

type AtomicOp uint32

const (
	InvalidAtomic AtomicOp = 0xffff_ffff

	AddAtomic AtomicOp = AtomicOp(Add) << 8

	FetchAdd AtomicOp = AddAtomic | fetch

	AndAtomic AtomicOp = AtomicOp(And) << 8

	FetchAnd AtomicOp = AndAtomic | fetch

	OrAtomic AtomicOp = AtomicOp(Or) << 8

	FetchOr AtomicOp = OrAtomic | fetch

	XorAtomic AtomicOp = AtomicOp(Xor) << 8

	FetchXor AtomicOp = XorAtomic | fetch

	Xchg AtomicOp = 0x0000_e000 | fetch

	CmpXchg AtomicOp = 0x0000_f000 | fetch

	fetch AtomicOp = 0x0000_0100

	loadAcquire AtomicOp = 0x0001_0000

	storeRelease AtomicOp = 0x0001_1000
)

func (op AtomicOp) String() string { _ = "STUB: not implemented"; return "" }

func (op AtomicOp) OpCode(size Size) OpCode { _ = "STUB: not implemented"; return *new(OpCode) }

func (op AtomicOp) Mem(dst, src Register, size Size, offset int16) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}

func LoadAcquire(dst, src Register, size Size, offset int16) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}

func StoreRelease(dst, src Register, size Size, offset int16) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}

type Size uint8

const sizeMask OpCode = 0x18

const (
	InvalidSize Size = 0xff

	DWord Size = 0x18

	Word Size = 0x00

	Half Size = 0x08

	Byte Size = 0x10
)

func (s Size) Sizeof() int { _ = "STUB: not implemented"; return 0 }

func LoadMemOp(size Size) OpCode { _ = "STUB: not implemented"; return *new(OpCode) }

func LoadMemSXOp(size Size) OpCode { _ = "STUB: not implemented"; return *new(OpCode) }

func LoadMem(dst, src Register, offset int16, size Size) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}

func LoadMemSX(dst, src Register, offset int16, size Size) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}

func LoadImmOp(size Size) OpCode { _ = "STUB: not implemented"; return *new(OpCode) }

func LoadImm(dst Register, value int64, size Size) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}

func LoadMapPtr(dst Register, fd int) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}

func LoadMapValue(dst Register, fd int, offset uint32) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}

func LoadIndOp(size Size) OpCode { _ = "STUB: not implemented"; return *new(OpCode) }

func LoadInd(dst, src Register, offset int32, size Size) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}

func LoadAbsOp(size Size) OpCode { _ = "STUB: not implemented"; return *new(OpCode) }

func LoadAbs(offset int32, size Size) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}

func StoreMemOp(size Size) OpCode { _ = "STUB: not implemented"; return *new(OpCode) }

func StoreMem(dst Register, offset int16, src Register, size Size) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}

func StoreImmOp(size Size) OpCode { _ = "STUB: not implemented"; return *new(OpCode) }

func StoreImm(dst Register, offset int16, value int64, size Size) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}

func StoreXAddOp(size Size) OpCode { _ = "STUB: not implemented"; return *new(OpCode) }

func StoreXAdd(dst, src Register, size Size) Instruction {
	_ = "STUB: not implemented"
	return *new(Instruction)
}
