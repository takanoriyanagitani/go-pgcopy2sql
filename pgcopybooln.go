package pgcopy2sql

// This is generated go file using go run pgcopynull.go. NEVER EDIT.

import (
	"fmt"
)

func (p *PgcopyBoolN) RawValue() bool {
	return p.value.Bool
}
func (p *PgcopyBoolN) IsValid() bool { return p.value.Valid }
func (p *PgcopyBoolN) IsNull() bool  { return !p.IsValid() }
func (p *PgcopyBoolN) Reset() {
	var zero PgcopyBoolN
	p.value = zero.value
}

func (p *PgcopyBoolN) SetValue(v bool) {
	p.value.Bool = v
	p.value.Valid = true
}

func (p *PgcopyBoolN) String() string {
	var raw bool = p.RawValue()
	switch p.IsValid() {
	case true:
		return fmt.Sprintf("Some(%v)", raw)
	default:
		return "None"
	}
}

func (p *PgcopyBoolN) WriteTo(writer ValueWriter) error {
	return writer.WriteNullBool(p.value)
}

func (p *PgcopyBoolN) AsValue() Value { return p }

func BoolNToPgcopy(v bool) (*PgcopyBoolN, error) {
	var p PgcopyBoolN
	p.SetValue(v)
	return &p, nil
}
