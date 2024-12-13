package pgcopy2sql

import (
	"database/sql"
	"fmt"
	"strings"
)

type PgcopyString struct{ strings.Builder }

func (s *PgcopyString) Reset()         { s.Builder.Reset() }
func (s *PgcopyString) String() string { return s.Builder.String() }

func (s *PgcopyString) WriteTo(w ValueWriter) error {
	return w.WriteString(s.Builder.String())
}

func (s *PgcopyString) SetBytes(
	b []byte,
	checker func(string) error,
) error {
	s.Reset()
	_, _ = s.Builder.Write(b) // error is always nil or OOM
	var val string = s.Builder.String()
	return checker(val)
}

func (s *PgcopyString) IsNull() bool   { return false }
func (s *PgcopyString) AsValue() Value { return s }

type PgcopyStringN struct {
	val   PgcopyString
	valid bool
}

func (s *PgcopyStringN) Reset() {
	s.val.Reset()
	s.valid = false
}

func (s *PgcopyStringN) SetBytes(
	b []byte,
	checker func(string) error,
) error {
	e := s.val.SetBytes(b, checker)
	s.valid = nil == e
	return e
}

func (s *PgcopyStringN) ToNullString() sql.NullString {
	return sql.NullString{
		String: s.val.String(),
		Valid:  s.valid,
	}
}

func (s *PgcopyStringN) WriteTo(w ValueWriter) error {
	return w.WriteNullString(s.ToNullString())
}

func (s *PgcopyStringN) IsNull() bool   { return !s.valid }
func (s *PgcopyStringN) AsValue() Value { return s }
func (s *PgcopyStringN) String() string {
	var v string = s.val.String()
	switch s.valid {
	case true:
		return fmt.Sprintf("Some(%s)", v)
	default:
		return "None"
	}
}
