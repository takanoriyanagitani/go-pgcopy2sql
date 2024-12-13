package pgcopy2sql

// This is generated go file using go run pgcopytype.go. NEVER EDIT.

import (
	"fmt"
)

func (p *PgcopyFloat) RawValue() float32 {
	return p.value
}

func (p *PgcopyFloat) IsNull() bool { return false }
func (p *PgcopyFloat) Reset() {
	p.value = 0.0
}

func (p *PgcopyFloat) SetValue(v float32) {
	p.value = v
}

func (p *PgcopyFloat) String() string {
	var raw float32 = p.value
	return fmt.Sprintf("%v", raw)
}

func (p *PgcopyFloat) WriteTo(writer ValueWriter) error {
	var raw float32 = p.value
	return writer.WriteFloat(raw)
}

func (p *PgcopyFloat) AsValue() Value { return p }

func FloatToPgcopy(v float32) (*PgcopyFloat, error) {
	return &PgcopyFloat{value: v}, nil
}
