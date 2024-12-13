package pgcopy2sql

// This is generated go file using go run pgcopytype.go. NEVER EDIT.

import (
	"fmt"
)

func (p *PgcopyInt) RawValue() int32 {
	return p.value
}

func (p *PgcopyInt) IsNull() bool { return false }
func (p *PgcopyInt) Reset() {
	p.value = 0
}

func (p *PgcopyInt) SetValue(v int32) {
	p.value = v
}

func (p *PgcopyInt) String() string {
	var raw int32 = p.value
	return fmt.Sprintf("%v", raw)
}

func (p *PgcopyInt) WriteTo(writer ValueWriter) error {
	var raw int32 = p.value
	return writer.WriteInt(raw)
}

func (p *PgcopyInt) AsValue() Value { return p }

func IntToPgcopy(v int32) (*PgcopyInt, error) {
	return &PgcopyInt{value: v}, nil
}
