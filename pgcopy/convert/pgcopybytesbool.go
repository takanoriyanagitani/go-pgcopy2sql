package convert

// This is generated go file using go run pgcopytype.go. NEVER EDIT.

import (
	"context"

	ps "github.com/takanoriyanagitani/go-pgcopy2sql"
	util "github.com/takanoriyanagitani/go-pgcopy2sql/util"
)

func BytesToBoolNew() func([]byte) (*ps.PgcopyBool, error) {
	var buf ps.PgcopyBool
	return func(raw []byte) (*ps.PgcopyBool, error) {
		buf.Reset()
		i, e := ps.BoolFromBytes(raw)
		if nil == e {
			buf.SetValue(i)
		}
		return &buf, nil
	}
}

func ConfigToConverterBool(
	_ ConvertConfig,
) func([]byte) util.IO[ps.Value] {
	var conv func(
		[]byte,
	) (*ps.PgcopyBool, error) = BytesToBoolNew()
	return func(raw []byte) util.IO[ps.Value] {
		return func(_ context.Context) (ps.Value, error) {
			return conv(raw)
		}
	}
}

// Slow converter which may use heap allocation.
func (t PgcopyBytesBool) Convert() (*ps.PgcopyBool, error) {
	var conv func(
		[]byte,
	) (*ps.PgcopyBool, error) = BytesToBoolNew()
	return conv(t)
}

func BytesToBoolNewN() func([]byte) (*ps.PgcopyBoolN, error) {
	var buf ps.PgcopyBoolN
	return func(raw []byte) (*ps.PgcopyBoolN, error) {
		buf.Reset()
		if nil == raw {
			return &buf, nil
		}

		i, e := ps.BoolFromBytes(raw)
		if nil == e {
			buf.SetValue(i)
		}
		return &buf, nil
	}
}

func ConfigToConverterBoolN(
	_ ConvertConfig,
) func([]byte) util.IO[ps.Value] {
	var conv func(
		[]byte,
	) (*ps.PgcopyBoolN, error) = BytesToBoolNewN()
	return func(raw []byte) util.IO[ps.Value] {
		return func(_ context.Context) (ps.Value, error) {
			return conv(raw)
		}
	}
}
