package pgcopy2sql

// This is generated go file using go run pgcopynull.go. NEVER EDIT.

import (
	"fmt"
)

func (p *PgcopyLongN) RawValue() int64 {
	return p.value.Int64
}
func (p *PgcopyLongN) IsValid() bool { return p.value.Valid }
func (p *PgcopyLongN) IsNull() bool  { return !p.IsValid() }
func (p *PgcopyLongN) Reset() {
	var zero PgcopyLongN
	p.value = zero.value
}

func (p *PgcopyLongN) SetValue(v int64) {
	p.value.Int64 = v
	p.value.Valid = true
}

func (p *PgcopyLongN) String() string {
	var raw int64 = p.RawValue()
	switch p.IsValid() {
	case true:
		return fmt.Sprintf("Some(%v)", raw)
	default:
		return "None"
	}
}

func (p *PgcopyLongN) WriteTo(writer ValueWriter) error {
	return writer.WriteNullLong(p.value)
}

func (p *PgcopyLongN) AsValue() Value { return p }

func LongNToPgcopy(v int64) (*PgcopyLongN, error) {
	var p PgcopyLongN
	p.SetValue(v)
	return &p, nil
}
