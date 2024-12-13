package pgcopy2sql

// This is generated go file using go run pgcopynull.go. NEVER EDIT.

import (
	"fmt"
)

func (p *PgcopyDoubleN) RawValue() float64 {
	return p.value.Float64
}
func (p *PgcopyDoubleN) IsValid() bool { return p.value.Valid }
func (p *PgcopyDoubleN) IsNull() bool  { return !p.IsValid() }
func (p *PgcopyDoubleN) Reset() {
	var zero PgcopyDoubleN
	p.value = zero.value
}

func (p *PgcopyDoubleN) SetValue(v float64) {
	p.value.Float64 = v
	p.value.Valid = true
}

func (p *PgcopyDoubleN) String() string {
	var raw float64 = p.RawValue()
	switch p.IsValid() {
	case true:
		return fmt.Sprintf("Some(%v)", raw)
	default:
		return "None"
	}
}

func (p *PgcopyDoubleN) WriteTo(writer ValueWriter) error {
	return writer.WriteNullDouble(p.value)
}

func (p *PgcopyDoubleN) AsValue() Value { return p }

func DoubleNToPgcopy(v float64) (*PgcopyDoubleN, error) {
	var p PgcopyDoubleN
	p.SetValue(v)
	return &p, nil
}
