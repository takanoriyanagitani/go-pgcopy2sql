package pgcopy2sql

// This is generated go file using go run pgcopynull.go. NEVER EDIT.

import (
	"fmt"
)

func (p *PgcopyShortN) RawValue() int16 {
	return p.value.Int16
}
func (p *PgcopyShortN) IsValid() bool { return p.value.Valid }
func (p *PgcopyShortN) IsNull() bool  { return !p.IsValid() }
func (p *PgcopyShortN) Reset() {
	var zero PgcopyShortN
	p.value = zero.value
}

func (p *PgcopyShortN) SetValue(v int16) {
	p.value.Int16 = v
	p.value.Valid = true
}

func (p *PgcopyShortN) String() string {
	var raw int16 = p.RawValue()
	switch p.IsValid() {
	case true:
		return fmt.Sprintf("Some(%v)", raw)
	default:
		return "None"
	}
}

func (p *PgcopyShortN) WriteTo(writer ValueWriter) error {
	return writer.WriteNullShort(p.value)
}

func (p *PgcopyShortN) AsValue() Value { return p }

func ShortNToPgcopy(v int16) (*PgcopyShortN, error) {
	var p PgcopyShortN
	p.SetValue(v)
	return &p, nil
}
