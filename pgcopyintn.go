package pgcopy2sql

// This is generated go file using go run pgcopynull.go. NEVER EDIT.

import (
	"fmt"
)

func (p *PgcopyIntN) RawValue() int32 {
	return p.value.Int32
}
func (p *PgcopyIntN) IsValid() bool { return p.value.Valid }
func (p *PgcopyIntN) IsNull() bool  { return !p.IsValid() }
func (p *PgcopyIntN) Reset() {
	var zero PgcopyIntN
	p.value = zero.value
}

func (p *PgcopyIntN) SetValue(v int32) {
	p.value.Int32 = v
	p.value.Valid = true
}

func (p *PgcopyIntN) String() string {
	var raw int32 = p.RawValue()
	switch p.IsValid() {
	case true:
		return fmt.Sprintf("Some(%v)", raw)
	default:
		return "None"
	}
}

func (p *PgcopyIntN) WriteTo(writer ValueWriter) error {
	return writer.WriteNullInt(p.value)
}

func (p *PgcopyIntN) AsValue() Value { return p }

func IntNToPgcopy(v int32) (*PgcopyIntN, error) {
	var p PgcopyIntN
	p.SetValue(v)
	return &p, nil
}
