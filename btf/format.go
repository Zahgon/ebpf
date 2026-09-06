package btf

import (
	"errors"
	"strings"
)

var errNestedTooDeep = errors.New("nested too deep")

type GoFormatter struct {
	w strings.Builder

	Names map[Type]string

	Identifier func(string) string

	EnumIdentifier func(name, element string) string
}

func (gf *GoFormatter) TypeDeclaration(name string, typ Type) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (gf *GoFormatter) identifier(s string) string { _ = "STUB: not implemented"; return "" }

func (gf *GoFormatter) enumIdentifier(name, element string) string {
	_ = "STUB: not implemented"
	return ""
}

func (gf *GoFormatter) writeTypeDecl(name string, typ Type) error {
	_ = "STUB: not implemented"
	return nil
}

func (gf *GoFormatter) writeType(typ Type, depth int) error { _ = "STUB: not implemented"; return nil }

func (gf *GoFormatter) writeTypeLit(typ Type, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func (gf *GoFormatter) writeIntLit(i *Int) error { _ = "STUB: not implemented"; return nil }

func (gf *GoFormatter) writeStructLit(size uint32, members []Member, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func (gf *GoFormatter) writeStructField(m Member, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func (gf *GoFormatter) writeDatasecLit(ds *Datasec, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func (gf *GoFormatter) writePadding(bytes uint32) { _ = "STUB: not implemented"; return }

func skipQualifiers(typ Type) Type { _ = "STUB: not implemented"; return *new(Type) }
