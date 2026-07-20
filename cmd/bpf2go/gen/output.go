//go:build !windows

package gen

import (
	_ "embed"
	"go/build/constraint"
	"io"
	"text/template"

	"github.com/cilium/ebpf/btf"
)

//go:embed output.tpl
var commonRaw string

var commonTemplate = template.Must(template.New("common").Parse(commonRaw))

type templateName string

func (n templateName) maybeExport(str string) string { _ = "STUB: not implemented"; return "" }

func (n templateName) Bytes() string { _ = "STUB: not implemented"; return "" }

func (n templateName) Specs() string { _ = "STUB: not implemented"; return "" }

func (n templateName) ProgramSpecs() string { _ = "STUB: not implemented"; return "" }

func (n templateName) MapSpecs() string { _ = "STUB: not implemented"; return "" }

func (n templateName) VariableSpecs() string { _ = "STUB: not implemented"; return "" }

func (n templateName) Load() string { _ = "STUB: not implemented"; return "" }

func (n templateName) LoadObjects() string { _ = "STUB: not implemented"; return "" }

func (n templateName) Objects() string { _ = "STUB: not implemented"; return "" }

func (n templateName) Maps() string { _ = "STUB: not implemented"; return "" }

func (n templateName) Variables() string { _ = "STUB: not implemented"; return "" }

func (n templateName) Programs() string { _ = "STUB: not implemented"; return "" }

func (n templateName) CloseHelper() string { _ = "STUB: not implemented"; return "" }

type GenerateArgs struct {
	Package string

	Stem string

	Constraints constraint.Expr

	Maps []string

	Variables []string

	Programs []string

	Types []btf.Type

	ObjectFile string

	Output io.Writer

	Identifier func(string) string
}

func Generate(args GenerateArgs) error { _ = "STUB: not implemented"; return nil }

func sortTypes(typeNames map[btf.Type]string) ([]btf.Type, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toUpperFirst(str string) string { _ = "STUB: not implemented"; return "" }
