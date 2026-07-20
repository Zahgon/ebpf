//go:build !windows

package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/cilium/ebpf/btf"
	"github.com/cilium/ebpf/cmd/bpf2go/gen"
)

const helpText = `Usage: %[1]s [options] <ident> <source file> [-- <C flags>]

ident is used as the stem of all generated Go types and functions, and
must be a valid Go identifier.

source is a single C file that is compiled using the specified compiler
(usually some version of clang).

You can pass options to the compiler by appending them after a '--' argument
or by supplying -cflags. Flags passed as arguments take precedence
over flags passed via -cflags. Additionally, the program expands quotation
marks in -cflags. This means that -cflags 'foo "bar baz"' is passed to the
compiler as two arguments "foo" and "bar baz".

The program expects GOPACKAGE to be set in the environment, and should be invoked
via go generate. The generated files are written to the current directory.

Some options take defaults from the environment. Variable name is mentioned
next to the respective option.

Options:

`

func run(stdout io.Writer, args []string) (err error) { _ = "STUB: not implemented"; return nil }

type bpf2go struct {
	stdout  io.Writer
	verbose bool

	sourceFile string

	outputDir string

	outputStem string

	outputSuffix string

	pkg string

	identStem string

	targetArches map[gen.Target]gen.GoArches

	cc string

	strip            string
	disableStripping bool

	cFlags          []string
	skipGlobalTypes bool

	cTypes cTypes

	tags buildTags

	makeBase string
}

func (b2g *bpf2go) Debugln(a ...any) { _ = "STUB: not implemented"; return }

func newB2G(stdout io.Writer, args []string) (*bpf2go, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type cTypes []string

var _ flag.Value = (*cTypes)(nil)

func (ct *cTypes) String() string { _ = "STUB: not implemented"; return "" }

const validCTypeChars = `[a-z0-9_]`

var reValidCType = regexp.MustCompile(`(?i)^` + validCTypeChars + `+$`)

func (ct *cTypes) Set(value string) error { _ = "STUB: not implemented"; return nil }

func getEnv(key, defaultVal string) string { _ = "STUB: not implemented"; return "" }

func getBool(key string, defaultVal bool) bool { _ = "STUB: not implemented"; return false }

func (b2g *bpf2go) convertAll() (err error) { _ = "STUB: not implemented"; return nil }

func (b2g *bpf2go) convert(tgt gen.Target, goarches gen.GoArches) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (b2g *bpf2go) removeOldOutputFiles(outputStem string, tgt gen.Target) error {
	_ = "STUB: not implemented"
	return nil
}

func printTargets(w io.Writer) { _ = "STUB: not implemented"; return }

func collectCTypes(types *btf.Spec, names []string) ([]btf.Type, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const gopackageEnv = "GOPACKAGE"

func main() {
	if err := run(os.Stdout, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
