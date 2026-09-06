//go:build !linux

package rlimit

func RemoveMemlock() error { _ = "STUB: not implemented"; return nil }
