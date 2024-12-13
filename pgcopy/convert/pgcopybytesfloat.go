package convert

// This is generated go file using go run pgcopytype.go. NEVER EDIT.

import (
	"context"

	ps "github.com/takanoriyanagitani/go-pgcopy2sql"
	util "github.com/takanoriyanagitani/go-pgcopy2sql/util"
)

func BytesToFloatNew() func([]byte) (*ps.PgcopyFloat, error) {
	var buf ps.PgcopyFloat
	return func(raw []byte) (*ps.PgcopyFloat, error) {
		buf.Reset()
		i, e := ps.FloatFromBytes(raw)
		if nil == e {
			buf.SetValue(i)
		}
		return &buf, nil
	}
}

func ConfigToConverterFloat(
	_ ConvertConfig,
) func([]byte) util.IO[ps.Value] {
	var conv func(
		[]byte,
	) (*ps.PgcopyFloat, error) = BytesToFloatNew()
	return func(raw []byte) util.IO[ps.Value] {
		return func(_ context.Context) (ps.Value, error) {
			return conv(raw)
		}
	}
}

// Slow converter which may use heap allocation.
func (t PgcopyBytesFloat) Convert() (*ps.PgcopyFloat, error) {
	var conv func(
		[]byte,
	) (*ps.PgcopyFloat, error) = BytesToFloatNew()
	return conv(t)
}

func BytesToFloatNewN() func([]byte) (*ps.PgcopyFloatN, error) {
	var buf ps.PgcopyFloatN
	return func(raw []byte) (*ps.PgcopyFloatN, error) {
		buf.Reset()
		if nil == raw {
			return &buf, nil
		}

		i, e := ps.FloatFromBytes(raw)
		if nil == e {
			buf.SetValue(i)
		}
		return &buf, nil
	}
}

func ConfigToConverterFloatN(
	_ ConvertConfig,
) func([]byte) util.IO[ps.Value] {
	var conv func(
		[]byte,
	) (*ps.PgcopyFloatN, error) = BytesToFloatNewN()
	return func(raw []byte) util.IO[ps.Value] {
		return func(_ context.Context) (ps.Value, error) {
			return conv(raw)
		}
	}
}
