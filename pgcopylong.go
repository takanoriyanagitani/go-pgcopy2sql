package pgcopy2sql

// This is generated go file using go run pgcopytype.go. NEVER EDIT.

import (
	"fmt"
)

func (p *PgcopyLong) RawValue() int64 {
	return p.value
}

func (p *PgcopyLong) IsNull() bool { return false }
func (p *PgcopyLong) Reset() {
	p.value = 0
}

func (p *PgcopyLong) SetValue(v int64) {
	p.value = v
}

func (p *PgcopyLong) String() string {
	var raw int64 = p.value
	return fmt.Sprintf("%v", raw)
}

func (p *PgcopyLong) WriteTo(writer ValueWriter) error {
	var raw int64 = p.value
	return writer.WriteLong(raw)
}

func (p *PgcopyLong) AsValue() Value { return p }

func LongToPgcopy(v int64) (*PgcopyLong, error) {
	return &PgcopyLong{value: v}, nil
}
