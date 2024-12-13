package pgcopy2sql

// This is generated go file using go run pgcopytype.go. NEVER EDIT.

import (
	"fmt"
)

func (p *PgcopyDouble) RawValue() float64 {
	return p.value
}

func (p *PgcopyDouble) IsNull() bool { return false }
func (p *PgcopyDouble) Reset() {
	p.value = 0.0
}

func (p *PgcopyDouble) SetValue(v float64) {
	p.value = v
}

func (p *PgcopyDouble) String() string {
	var raw float64 = p.value
	return fmt.Sprintf("%v", raw)
}

func (p *PgcopyDouble) WriteTo(writer ValueWriter) error {
	var raw float64 = p.value
	return writer.WriteDouble(raw)
}

func (p *PgcopyDouble) AsValue() Value { return p }

func DoubleToPgcopy(v float64) (*PgcopyDouble, error) {
	return &PgcopyDouble{value: v}, nil
}
