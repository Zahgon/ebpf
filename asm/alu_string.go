package asm

func _() { _ = "STUB: not implemented"; return }

const (
	_Source_name_0 = "ImmSource"
	_Source_name_1 = "RegSource"
	_Source_name_2 = "InvalidSource"
)

func (i Source) String() string { _ = "STUB: not implemented"; return "" }

func _() { _ = "STUB: not implemented"; return }

const (
	_Endianness_name_0 = "LE"
	_Endianness_name_1 = "BE"
	_Endianness_name_2 = "InvalidEndian"
)

func (i Endianness) String() string { _ = "STUB: not implemented"; return "" }

func _() { _ = "STUB: not implemented"; return }

const _ALUOp_name = "AddSubMulDivOrAndLShRShNegModXorMovArShSwapSDivSModMovSX8MovSX16MovSX32InvalidALUOp"

var _ALUOp_map = map[ALUOp]string{
	0:     _ALUOp_name[0:3],
	16:    _ALUOp_name[3:6],
	32:    _ALUOp_name[6:9],
	48:    _ALUOp_name[9:12],
	64:    _ALUOp_name[12:14],
	80:    _ALUOp_name[14:17],
	96:    _ALUOp_name[17:20],
	112:   _ALUOp_name[20:23],
	128:   _ALUOp_name[23:26],
	144:   _ALUOp_name[26:29],
	160:   _ALUOp_name[29:32],
	176:   _ALUOp_name[32:35],
	192:   _ALUOp_name[35:39],
	208:   _ALUOp_name[39:43],
	304:   _ALUOp_name[43:47],
	400:   _ALUOp_name[47:51],
	432:   _ALUOp_name[51:57],
	688:   _ALUOp_name[57:64],
	944:   _ALUOp_name[64:71],
	65535: _ALUOp_name[71:83],
}

func (i ALUOp) String() string { _ = "STUB: not implemented"; return "" }
