package convert

// This is generated go file using go run pgcopytype.go. NEVER EDIT.

import (
	"context"

	ps "github.com/takanoriyanagitani/go-pgcopy2sql"
	util "github.com/takanoriyanagitani/go-pgcopy2sql/util"
)

func BytesToShortNew() func([]byte) (*ps.PgcopyShort, error) {
	var buf ps.PgcopyShort
	return func(raw []byte) (*ps.PgcopyShort, error) {
		buf.Reset()
		i, e := ps.ShortFromBytes(raw)
		if nil == e {
			buf.SetValue(i)
		}
		return &buf, nil
	}
}

func ConfigToConverterShort(
	_ ConvertConfig,
) func([]byte) util.IO[ps.Value] {
	var conv func(
		[]byte,
	) (*ps.PgcopyShort, error) = BytesToShortNew()
	return func(raw []byte) util.IO[ps.Value] {
		return func(_ context.Context) (ps.Value, error) {
			return conv(raw)
		}
	}
}

// Slow converter which may use heap allocation.
func (t PgcopyBytesShort) Convert() (*ps.PgcopyShort, error) {
	var conv func(
		[]byte,
	) (*ps.PgcopyShort, error) = BytesToShortNew()
	return conv(t)
}

func BytesToShortNewN() func([]byte) (*ps.PgcopyShortN, error) {
	var buf ps.PgcopyShortN
	return func(raw []byte) (*ps.PgcopyShortN, error) {
		buf.Reset()
		if nil == raw {
			return &buf, nil
		}

		i, e := ps.ShortFromBytes(raw)
		if nil == e {
			buf.SetValue(i)
		}
		return &buf, nil
	}
}

func ConfigToConverterShortN(
	_ ConvertConfig,
) func([]byte) util.IO[ps.Value] {
	var conv func(
		[]byte,
	) (*ps.PgcopyShortN, error) = BytesToShortNewN()
	return func(raw []byte) util.IO[ps.Value] {
		return func(_ context.Context) (ps.Value, error) {
			return conv(raw)
		}
	}
}
