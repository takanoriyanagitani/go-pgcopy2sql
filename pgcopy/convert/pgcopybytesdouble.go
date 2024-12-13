package convert

// This is generated go file using go run pgcopytype.go. NEVER EDIT.

import (
	"context"

	ps "github.com/takanoriyanagitani/go-pgcopy2sql"
	util "github.com/takanoriyanagitani/go-pgcopy2sql/util"
)

func BytesToDoubleNew() func([]byte) (*ps.PgcopyDouble, error) {
	var buf ps.PgcopyDouble
	return func(raw []byte) (*ps.PgcopyDouble, error) {
		buf.Reset()
		i, e := ps.DoubleFromBytes(raw)
		if nil == e {
			buf.SetValue(i)
		}
		return &buf, nil
	}
}

func ConfigToConverterDouble(
	_ ConvertConfig,
) func([]byte) util.IO[ps.Value] {
	var conv func(
		[]byte,
	) (*ps.PgcopyDouble, error) = BytesToDoubleNew()
	return func(raw []byte) util.IO[ps.Value] {
		return func(_ context.Context) (ps.Value, error) {
			return conv(raw)
		}
	}
}

// Slow converter which may use heap allocation.
func (t PgcopyBytesDouble) Convert() (*ps.PgcopyDouble, error) {
	var conv func(
		[]byte,
	) (*ps.PgcopyDouble, error) = BytesToDoubleNew()
	return conv(t)
}

func BytesToDoubleNewN() func([]byte) (*ps.PgcopyDoubleN, error) {
	var buf ps.PgcopyDoubleN
	return func(raw []byte) (*ps.PgcopyDoubleN, error) {
		buf.Reset()
		if nil == raw {
			return &buf, nil
		}

		i, e := ps.DoubleFromBytes(raw)
		if nil == e {
			buf.SetValue(i)
		}
		return &buf, nil
	}
}

func ConfigToConverterDoubleN(
	_ ConvertConfig,
) func([]byte) util.IO[ps.Value] {
	var conv func(
		[]byte,
	) (*ps.PgcopyDoubleN, error) = BytesToDoubleNewN()
	return func(raw []byte) util.IO[ps.Value] {
		return func(_ context.Context) (ps.Value, error) {
			return conv(raw)
		}
	}
}
