package asm

//go:generate go tool stringer -output func_string.go -type=BuiltinFunc

type BuiltinFunc uint32

func BuiltinFuncForPlatform(plat string, value uint32) (BuiltinFunc, error) {
	_ = "STUB: not implemented"
	return *new(BuiltinFunc), nil
}

func (fn BuiltinFunc) Call() Instruction { _ = "STUB: not implemented"; return *new(Instruction) }
