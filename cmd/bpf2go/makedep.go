//go:build !windows

package main

import (
	"io"
)

func adjustDependencies(w io.Writer, baseDir string, deps []dependency) error {
	_ = "STUB: not implemented"
	return nil
}

type dependency struct {
	file          string
	prerequisites []string
}

func parseDependencies(baseDir string, in io.Reader) ([]dependency, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
