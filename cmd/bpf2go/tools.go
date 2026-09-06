//go:build !windows

package main

func splitCFlagsFromArgs(in []string) (args, cflags []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

func splitArguments(in string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }
