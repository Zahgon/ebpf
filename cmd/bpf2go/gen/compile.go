//go:build !windows

package gen

type CompileArgs struct {
	CC string

	Strip string

	Flags []string

	Workdir string

	Source string

	Dest string

	Target           Target
	DisableStripping bool
}

func insertDefaultFlags(flags []string) []string { _ = "STUB: not implemented"; return nil }

func Compile(args CompileArgs) error { _ = "STUB: not implemented"; return nil }
