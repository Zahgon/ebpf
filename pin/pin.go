package pin

import "io"

type Pin struct {
	Path   string
	Object io.Closer
}

func (p *Pin) close() { _ = "STUB: not implemented"; return }

func (p *Pin) Take() { _ = "STUB: not implemented"; return }
