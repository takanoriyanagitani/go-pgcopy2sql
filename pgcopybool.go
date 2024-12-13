package pgcopy2sql

// This is generated go file using go run pgcopytype.go. NEVER EDIT.

import (
	"fmt"
)

func (p *PgcopyBool) RawValue() bool {
	return p.value
}

func (p *PgcopyBool) IsNull() bool { return false }
func (p *PgcopyBool) Reset() {
	p.value = false
}

func (p *PgcopyBool) SetValue(v bool) {
	p.value = v
}

func (p *PgcopyBool) String() string {
	var raw bool = p.value
	return fmt.Sprintf("%v", raw)
}

func (p *PgcopyBool) WriteTo(writer ValueWriter) error {
	var raw bool = p.value
	return writer.WriteBool(raw)
}

func (p *PgcopyBool) AsValue() Value { return p }

func BoolToPgcopy(v bool) (*PgcopyBool, error) {
	return &PgcopyBool{value: v}, nil
}
