package pgcopy2sql

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type PgcopyUuid struct{ val [16]byte }

func (u *PgcopyUuid) Reset()                      { u.val = [16]byte{} }
func (u *PgcopyUuid) WriteTo(w ValueWriter) error { return w.WriteUuid(u.val) }
func (u *PgcopyUuid) IsNull() bool                { return false }
func (u *PgcopyUuid) SetValue(v [16]byte)         { u.val = v }
func (u *PgcopyUuid) AsValue() Value              { return u }
func (u *PgcopyUuid) String() string {
	var v uuid.UUID = uuid.UUID(u.val)
	return v.String()
}

func UuidToPgcopy(u [16]byte) (*PgcopyUuid, error) {
	return &PgcopyUuid{val: u}, nil
}

type PgcopyUuidN struct{ sql.Null[[16]byte] }

func (u *PgcopyUuidN) WriteTo(w ValueWriter) error {
	return w.WriteNullUuid(u.Null)
}

func (u *PgcopyUuidN) Reset()         { u.Null = sql.Null[[16]byte]{} }
func (u *PgcopyUuidN) IsNull() bool   { return !u.Null.Valid }
func (u *PgcopyUuidN) AsValue() Value { return u }
func (u *PgcopyUuidN) SetValue(v [16]byte) {
	u.Null.Valid = true
	u.Null.V = v
}
func (u *PgcopyUuidN) String() string {
	var v [16]byte = u.Null.V
	if !u.Null.Valid {
		return "None"
	}
	nn := &PgcopyUuid{val: v}
	var s string = nn.String()
	return fmt.Sprintf("Some(%s)", s)
}

func UuidToPgcopyUuidN(u [16]byte) *PgcopyUuidN {
	return &PgcopyUuidN{Null: sql.Null[[16]byte]{V: u, Valid: true}}
}
