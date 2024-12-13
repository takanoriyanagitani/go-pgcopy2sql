package pgcopy2sql

import (
	"bytes"
	"fmt"
)

type PgcopyBytes struct{ bytes.Buffer }

func (b *PgcopyBytes) Reset() { b.Buffer.Reset() }

func (b *PgcopyBytes) WriteTo(w ValueWriter) error {
	return w.WriteBytes(b.Bytes())
}

func (b *PgcopyBytes) IsNull() bool   { return false }
func (b *PgcopyBytes) AsValue() Value { return b }
func (b *PgcopyBytes) Bytes() []byte  { return b.Buffer.Bytes() }
func (b *PgcopyBytes) SetBytes(v []byte) {
	b.Reset()
	_, _ = b.Buffer.Write(v) // error is always nil or panic
}
func (b *PgcopyBytes) String() string {
	var v []byte = b.Buffer.Bytes()
	return fmt.Sprintf("%q", v)
}

type PgcopyBytesN struct {
	val   PgcopyBytes
	valid bool
}

func (b *PgcopyBytesN) IsNull() bool { return !b.valid }
func (b *PgcopyBytesN) Reset() {
	b.val.Reset()
	b.valid = false
}

func (b *PgcopyBytesN) WriteTo(w ValueWriter) error {
	switch b.IsNull() {
	case true:
		return w.WriteBytes(nil)
	default:
		return w.WriteBytes(b.val.Bytes())
	}
}

func (b *PgcopyBytesN) AsValue() Value { return b }
func (b *PgcopyBytesN) String() string {
	var s string = b.val.String()
	switch b.valid {
	case true:
		return fmt.Sprintf("Some(%s)", s)
	default:
		return "None"
	}
}
func (b *PgcopyBytesN) SetBytes(v []byte) {
	b.val.SetBytes(v)
	b.valid = true
}
