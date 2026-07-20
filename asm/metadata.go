package asm

type Metadata struct {
	head *metaElement
}

type metaElement struct {
	next       *metaElement
	key, value any
}

func (m *Metadata) find(key any) *metaElement { _ = "STUB: not implemented"; return nil }

func (m *Metadata) remove(r *metaElement) { _ = "STUB: not implemented"; return }

func (m *Metadata) Set(key, value any) { _ = "STUB: not implemented"; return }

func (m *Metadata) Get(key any) any { _ = "STUB: not implemented"; return *new(any) }
