package convert

// This is generated go file using go run pgcopytype.go. NEVER EDIT.

import (
	"context"

	ps "github.com/takanoriyanagitani/go-pgcopy2sql"
	util "github.com/takanoriyanagitani/go-pgcopy2sql/util"
)

func BytesToIntNew() func([]byte) (*ps.PgcopyInt, error) {
	var buf ps.PgcopyInt
	return func(raw []byte) (*ps.PgcopyInt, error) {
		buf.Reset()
		i, e := ps.IntFromBytes(raw)
		if nil == e {
			buf.SetValue(i)
		}
		return &buf, nil
	}
}

func ConfigToConverterInt(
	_ ConvertConfig,
) func([]byte) util.IO[ps.Value] {
	var conv func(
		[]byte,
	) (*ps.PgcopyInt, error) = BytesToIntNew()
	return func(raw []byte) util.IO[ps.Value] {
		return func(_ context.Context) (ps.Value, error) {
			return conv(raw)
		}
	}
}

// Slow converter which may use heap allocation.
func (t PgcopyBytesInt) Convert() (*ps.PgcopyInt, error) {
	var conv func(
		[]byte,
	) (*ps.PgcopyInt, error) = BytesToIntNew()
	return conv(t)
}

func BytesToIntNewN() func([]byte) (*ps.PgcopyIntN, error) {
	var buf ps.PgcopyIntN
	return func(raw []byte) (*ps.PgcopyIntN, error) {
		buf.Reset()
		if nil == raw {
			return &buf, nil
		}

		i, e := ps.IntFromBytes(raw)
		if nil == e {
			buf.SetValue(i)
		}
		return &buf, nil
	}
}

func ConfigToConverterIntN(
	_ ConvertConfig,
) func([]byte) util.IO[ps.Value] {
	var conv func(
		[]byte,
	) (*ps.PgcopyIntN, error) = BytesToIntNewN()
	return func(raw []byte) util.IO[ps.Value] {
		return func(_ context.Context) (ps.Value, error) {
			return conv(raw)
		}
	}
}
