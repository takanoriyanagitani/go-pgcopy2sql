package pgcopy2sql

// This is generated go file using go run pgcopytype.go. NEVER EDIT.

import (
	"fmt"
)

func (p *PgcopyShort) RawValue() int16 {
	return p.value
}

func (p *PgcopyShort) IsNull() bool { return false }
func (p *PgcopyShort) Reset() {
	p.value = 0
}

func (p *PgcopyShort) SetValue(v int16) {
	p.value = v
}

func (p *PgcopyShort) String() string {
	var raw int16 = p.value
	return fmt.Sprintf("%v", raw)
}

func (p *PgcopyShort) WriteTo(writer ValueWriter) error {
	var raw int16 = p.value
	return writer.WriteShort(raw)
}

func (p *PgcopyShort) AsValue() Value { return p }

func ShortToPgcopy(v int16) (*PgcopyShort, error) {
	return &PgcopyShort{value: v}, nil
}
