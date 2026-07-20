package btf

import (
	"errors"
	"fmt"

	"github.com/cilium/ebpf/asm"
	"github.com/cilium/ebpf/internal"
	"github.com/cilium/ebpf/internal/sys"
)

const maxResolveDepth = 32

type TypeID = sys.TypeID

type Type interface {
	fmt.Formatter

	TypeName() string

	copy() Type
}

var (
	_ Type = (*Int)(nil)
	_ Type = (*Struct)(nil)
	_ Type = (*Union)(nil)
	_ Type = (*Enum)(nil)
	_ Type = (*Fwd)(nil)
	_ Type = (*Func)(nil)
	_ Type = (*Typedef)(nil)
	_ Type = (*Var)(nil)
	_ Type = (*Datasec)(nil)
	_ Type = (*Float)(nil)
	_ Type = (*declTag)(nil)
	_ Type = (*TypeTag)(nil)
	_ Type = (*cycle)(nil)
)

type Void struct{}

func (v *Void) Format(fs fmt.State, verb rune) { _ = "STUB: not implemented"; return }
func (v *Void) TypeName() string               { _ = "STUB: not implemented"; return "" }
func (v *Void) size() uint32                   { _ = "STUB: not implemented"; return 0 }
func (v *Void) copy() Type                     { _ = "STUB: not implemented"; return *new(Type) }

type IntEncoding byte

const (
	Unsigned IntEncoding = 0
	Signed   IntEncoding = 1
	Char     IntEncoding = 2
	Bool     IntEncoding = 4
)

func (ie IntEncoding) String() string { _ = "STUB: not implemented"; return "" }

type Int struct {
	Name string

	Size     uint32
	Encoding IntEncoding
}

func (i *Int) Format(fs fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (i *Int) TypeName() string { _ = "STUB: not implemented"; return "" }
func (i *Int) size() uint32     { _ = "STUB: not implemented"; return 0 }
func (i *Int) copy() Type       { _ = "STUB: not implemented"; return *new(Type) }

type Pointer struct {
	Target Type
}

func (p *Pointer) Format(fs fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (p *Pointer) TypeName() string { _ = "STUB: not implemented"; return "" }
func (p *Pointer) size() uint32     { _ = "STUB: not implemented"; return 0 }
func (p *Pointer) copy() Type       { _ = "STUB: not implemented"; return *new(Type) }

type Array struct {
	Index  Type
	Type   Type
	Nelems uint32
}

func (arr *Array) Format(fs fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (arr *Array) TypeName() string { _ = "STUB: not implemented"; return "" }

func (arr *Array) copy() Type { _ = "STUB: not implemented"; return *new(Type) }

type Struct struct {
	Name string

	Size    uint32
	Members []Member
	Tags    []string
}

func (s *Struct) Format(fs fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (s *Struct) TypeName() string { _ = "STUB: not implemented"; return "" }

func (s *Struct) size() uint32 { _ = "STUB: not implemented"; return 0 }

func (s *Struct) copy() Type { _ = "STUB: not implemented"; return *new(Type) }

func (s *Struct) members() []Member { _ = "STUB: not implemented"; return nil }

type Union struct {
	Name string

	Size    uint32
	Members []Member
	Tags    []string
}

func (u *Union) Format(fs fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (u *Union) TypeName() string { _ = "STUB: not implemented"; return "" }

func (u *Union) size() uint32 { _ = "STUB: not implemented"; return 0 }

func (u *Union) copy() Type { _ = "STUB: not implemented"; return *new(Type) }

func (u *Union) members() []Member { _ = "STUB: not implemented"; return nil }

func memberNames(members []Member, max int) []string { _ = "STUB: not implemented"; return nil }

func copyMembers(orig []Member) []Member { _ = "STUB: not implemented"; return nil }

func copyTags(orig []string) []string { _ = "STUB: not implemented"; return nil }

type composite interface {
	Type
	members() []Member
}

var (
	_ composite = (*Struct)(nil)
	_ composite = (*Union)(nil)
)

type Bits uint32

func (b Bits) Bytes() uint32 { _ = "STUB: not implemented"; return 0 }

type Member struct {
	Name         string
	Type         Type
	Offset       Bits
	BitfieldSize Bits
	Tags         []string
}

type Enum struct {
	Name string

	Size uint32

	Signed bool
	Values []EnumValue
}

func (e *Enum) Format(fs fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (e *Enum) TypeName() string { _ = "STUB: not implemented"; return "" }

type EnumValue struct {
	Name  string
	Value uint64
}

func (e *Enum) size() uint32 { _ = "STUB: not implemented"; return 0 }
func (e *Enum) copy() Type   { _ = "STUB: not implemented"; return *new(Type) }

type FwdKind int

const (
	FwdStruct FwdKind = iota
	FwdUnion
)

func (fk FwdKind) String() string { _ = "STUB: not implemented"; return "" }

type Fwd struct {
	Name string
	Kind FwdKind
}

func (f *Fwd) Format(fs fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (f *Fwd) TypeName() string { _ = "STUB: not implemented"; return "" }

func (f *Fwd) copy() Type { _ = "STUB: not implemented"; return *new(Type) }

func (f *Fwd) matches(typ Type) bool { _ = "STUB: not implemented"; return false }

type Typedef struct {
	Name string
	Type Type
	Tags []string
}

func (td *Typedef) Format(fs fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (td *Typedef) TypeName() string { _ = "STUB: not implemented"; return "" }

func (td *Typedef) copy() Type { _ = "STUB: not implemented"; return *new(Type) }

type Volatile struct {
	Type Type
}

func (v *Volatile) Format(fs fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (v *Volatile) TypeName() string { _ = "STUB: not implemented"; return "" }

func (v *Volatile) qualify() Type { _ = "STUB: not implemented"; return *new(Type) }
func (v *Volatile) copy() Type    { _ = "STUB: not implemented"; return *new(Type) }

type Const struct {
	Type Type
}

func (c *Const) Format(fs fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (c *Const) TypeName() string { _ = "STUB: not implemented"; return "" }

func (c *Const) qualify() Type { _ = "STUB: not implemented"; return *new(Type) }
func (c *Const) copy() Type    { _ = "STUB: not implemented"; return *new(Type) }

type Restrict struct {
	Type Type
}

func (r *Restrict) Format(fs fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (r *Restrict) TypeName() string { _ = "STUB: not implemented"; return "" }

func (r *Restrict) qualify() Type { _ = "STUB: not implemented"; return *new(Type) }
func (r *Restrict) copy() Type    { _ = "STUB: not implemented"; return *new(Type) }

type Func struct {
	Name    string
	Type    Type
	Linkage FuncLinkage
	Tags    []string

	ParamTags [][]string
}

type funcInfoMeta struct{}

func FuncMetadata(ins *asm.Instruction) *Func { _ = "STUB: not implemented"; return nil }

func WithFuncMetadata(ins asm.Instruction, fn *Func) asm.Instruction {
	_ = "STUB: not implemented"
	return *new(asm.Instruction)
}

func (f *Func) Format(fs fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (f *Func) TypeName() string { _ = "STUB: not implemented"; return "" }

func (f *Func) copy() Type { _ = "STUB: not implemented"; return *new(Type) }

type FuncProto struct {
	Return Type
	Params []FuncParam
}

func (fp *FuncProto) Format(fs fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (fp *FuncProto) TypeName() string { _ = "STUB: not implemented"; return "" }

func (fp *FuncProto) copy() Type { _ = "STUB: not implemented"; return *new(Type) }

type FuncParam struct {
	Name string
	Type Type
}

type Var struct {
	Name    string
	Type    Type
	Linkage VarLinkage
	Tags    []string
}

func (v *Var) Format(fs fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (v *Var) TypeName() string { _ = "STUB: not implemented"; return "" }

func (v *Var) copy() Type { _ = "STUB: not implemented"; return *new(Type) }

type Datasec struct {
	Name string
	Size uint32
	Vars []VarSecinfo
}

func (ds *Datasec) Format(fs fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (ds *Datasec) TypeName() string { _ = "STUB: not implemented"; return "" }

func (ds *Datasec) size() uint32 { _ = "STUB: not implemented"; return 0 }

func (ds *Datasec) copy() Type { _ = "STUB: not implemented"; return *new(Type) }

type VarSecinfo struct {
	Type   Type
	Offset uint32
	Size   uint32
}

type Float struct {
	Name string

	Size uint32
}

func (f *Float) Format(fs fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (f *Float) TypeName() string { _ = "STUB: not implemented"; return "" }
func (f *Float) size() uint32     { _ = "STUB: not implemented"; return 0 }
func (f *Float) copy() Type       { _ = "STUB: not implemented"; return *new(Type) }

type declTag struct {
	Type  Type
	Value string

	Index int
}

func (dt *declTag) Format(fs fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (dt *declTag) TypeName() string { _ = "STUB: not implemented"; return "" }
func (dt *declTag) copy() Type       { _ = "STUB: not implemented"; return *new(Type) }

type TypeTag struct {
	Type  Type
	Value string
}

func (tt *TypeTag) Format(fs fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (tt *TypeTag) TypeName() string { _ = "STUB: not implemented"; return "" }
func (tt *TypeTag) qualify() Type    { _ = "STUB: not implemented"; return *new(Type) }
func (tt *TypeTag) copy() Type       { _ = "STUB: not implemented"; return *new(Type) }

type cycle struct {
	root Type
}

func (c *cycle) ID() TypeID                     { _ = "STUB: not implemented"; return *new(TypeID) }
func (c *cycle) Format(fs fmt.State, verb rune) { _ = "STUB: not implemented"; return }
func (c *cycle) TypeName() string               { _ = "STUB: not implemented"; return "" }
func (c *cycle) copy() Type                     { _ = "STUB: not implemented"; return *new(Type) }

type sizer interface {
	size() uint32
}

var (
	_ sizer = (*Int)(nil)
	_ sizer = (*Pointer)(nil)
	_ sizer = (*Struct)(nil)
	_ sizer = (*Union)(nil)
	_ sizer = (*Enum)(nil)
	_ sizer = (*Datasec)(nil)
)

type qualifier interface {
	qualify() Type
}

var (
	_ qualifier = (*Const)(nil)
	_ qualifier = (*Restrict)(nil)
	_ qualifier = (*Volatile)(nil)
	_ qualifier = (*TypeTag)(nil)
)

var errUnsizedType = errors.New("type is unsized")

func Sizeof(typ Type) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func alignof(typ Type) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func Copy(typ Type) Type { _ = "STUB: not implemented"; return *new(Type) }

func copyType(typ Type, ids map[Type]TypeID, copies map[Type]Type, copiedIDs map[Type]TypeID) Type {
	_ = "STUB: not implemented"
	return *new(Type)
}

type typeDeque = internal.Deque[*Type]

type essentialName string

func newEssentialName(name string) essentialName {
	_ = "STUB: not implemented"
	return *new(essentialName)
}

func UnderlyingType(typ Type) Type { _ = "STUB: not implemented"; return *new(Type) }

func QualifiedType(typ Type) Type { _ = "STUB: not implemented"; return *new(Type) }

func As[T Type](typ Type) (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

type formatState struct {
	fmt.State
	depth int
}

type formattableType interface {
	fmt.Formatter
	TypeName() string
}

func formatType(f fmt.State, verb rune, t formattableType, extra ...any) {
	_ = "STUB: not implemented"
	return
}
