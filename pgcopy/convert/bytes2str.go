package convert

import (
	"errors"
	"unicode/utf8"

	ps "github.com/takanoriyanagitani/go-pgcopy2sql"
)

var (
	ErrInvalidUtf8 error = errors.New("invalid utf8")
)

func StringFromBytesNew(
	checker func(string) error,
) func([]byte) (*ps.PgcopyString, error) {
	var buf ps.PgcopyString
	return func(b []byte) (*ps.PgcopyString, error) {
		e := buf.SetBytes(b, checker)
		return &buf, e
	}
}

func CheckStringUtf8(s string) error {
	var valid bool = utf8.ValidString(s)
	if !valid {
		return ErrInvalidUtf8
	}
	return nil
}

func StringFromBytesNewDefault() func([]byte) (*ps.PgcopyString, error) {
	return StringFromBytesNew(CheckStringUtf8)
}

func NullStringFromBytesNew(
	checker func(string) error,
) func([]byte) (*ps.PgcopyStringN, error) {
	var buf ps.PgcopyStringN
	return func(b []byte) (*ps.PgcopyStringN, error) {
		if nil == b {
			buf.Reset()
			return &buf, nil
		}
		e := buf.SetBytes(b, checker)
		return &buf, e
	}
}

func NullStringFromBytesNewDefault() func([]byte) (*ps.PgcopyStringN, error) {
	return NullStringFromBytesNew(CheckStringUtf8)
}
