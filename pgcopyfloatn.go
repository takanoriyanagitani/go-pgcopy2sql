package pgcopy2sql

import (
	"fmt"
)

func (p *PgcopyFloatN) RawValue() float32 {
	return p.value.V
}
func (p *PgcopyFloatN) IsValid() bool { return p.value.Valid }
func (p *PgcopyFloatN) IsNull() bool  { return !p.IsValid() }
func (p *PgcopyFloatN) Reset() {
	var zero PgcopyFloatN
	p.value = zero.value
}

func (p *PgcopyFloatN) SetValue(v float32) {
	p.value.V = v
	p.value.Valid = true
}

func (p *PgcopyFloatN) String() string {
	var raw float32 = p.RawValue()
	switch p.IsValid() {
	case true:
		return fmt.Sprintf("Some(%v)", raw)
	default:
		return "None"
	}
}

func (p *PgcopyFloatN) WriteTo(writer ValueWriter) error {
	return writer.WriteNullFloat(p.value)
}

func (p *PgcopyFloatN) AsValue() Value { return p }

func FloatNToPgcopy(v float32) (*PgcopyFloatN, error) {
	var p PgcopyFloatN
	p.SetValue(v)
	return &p, nil
}
