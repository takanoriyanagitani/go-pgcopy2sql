package convert

// This is generated go file using go run pgcopytype.go. NEVER EDIT.

import (
	"context"

	ps "github.com/takanoriyanagitani/go-pgcopy2sql"
	util "github.com/takanoriyanagitani/go-pgcopy2sql/util"
)

func BytesToLongNew() func([]byte) (*ps.PgcopyLong, error) {
	var buf ps.PgcopyLong
	return func(raw []byte) (*ps.PgcopyLong, error) {
		buf.Reset()
		i, e := ps.LongFromBytes(raw)
		if nil == e {
			buf.SetValue(i)
		}
		return &buf, nil
	}
}

func ConfigToConverterLong(
	_ ConvertConfig,
) func([]byte) util.IO[ps.Value] {
	var conv func(
		[]byte,
	) (*ps.PgcopyLong, error) = BytesToLongNew()
	return func(raw []byte) util.IO[ps.Value] {
		return func(_ context.Context) (ps.Value, error) {
			return conv(raw)
		}
	}
}

// Slow converter which may use heap allocation.
func (t PgcopyBytesLong) Convert() (*ps.PgcopyLong, error) {
	var conv func(
		[]byte,
	) (*ps.PgcopyLong, error) = BytesToLongNew()
	return conv(t)
}

func BytesToLongNewN() func([]byte) (*ps.PgcopyLongN, error) {
	var buf ps.PgcopyLongN
	return func(raw []byte) (*ps.PgcopyLongN, error) {
		buf.Reset()
		if nil == raw {
			return &buf, nil
		}

		i, e := ps.LongFromBytes(raw)
		if nil == e {
			buf.SetValue(i)
		}
		return &buf, nil
	}
}

func ConfigToConverterLongN(
	_ ConvertConfig,
) func([]byte) util.IO[ps.Value] {
	var conv func(
		[]byte,
	) (*ps.PgcopyLongN, error) = BytesToLongNewN()
	return func(raw []byte) util.IO[ps.Value] {
		return func(_ context.Context) (ps.Value, error) {
			return conv(raw)
		}
	}
}
