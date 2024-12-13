package pgcopy2sql

import (
	"database/sql"
	"fmt"
	"time"
)

type PgcopyTime struct{ val time.Time }

func (t *PgcopyTime) Reset()                      { t.val = time.Time{} }
func (t *PgcopyTime) WriteTo(w ValueWriter) error { return w.WriteTime(t.val) }
func (t *PgcopyTime) IsNull() bool                { return false }
func (t *PgcopyTime) SetValue(v time.Time)        { t.val = v }
func (t *PgcopyTime) AsValue() Value              { return t }
func (t *PgcopyTime) String() string              { return t.val.String() }

func TimeToPgcopy(t time.Time) (*PgcopyTime, error) {
	return &PgcopyTime{val: t}, nil
}

type PgcopyTimeN struct{ val sql.NullTime }

func (t *PgcopyTimeN) Reset() {
	var z sql.NullTime
	t.val = z
}
func (t *PgcopyTimeN) WriteTo(w ValueWriter) error {
	return w.WriteNullTime(t.val)
}
func (t *PgcopyTimeN) SetValue(v time.Time) {
	t.val.Time = v
	t.val.Valid = true
}
func (t *PgcopyTimeN) String() string {
	switch t.IsValid() {
	case true:
		return fmt.Sprintf("Some(%v)", t.val.Time)
	default:
		return "None"
	}
}
func (t *PgcopyTimeN) IsValid() bool  { return t.val.Valid }
func (t *PgcopyTimeN) IsNull() bool   { return !t.IsValid() }
func (t *PgcopyTimeN) AsValue() Value { return t }

func TimeToPgcopyN(t time.Time) (*PgcopyTimeN, error) {
	var z PgcopyTimeN
	z.val.Time = t
	z.val.Valid = true
	return &z, nil
}
