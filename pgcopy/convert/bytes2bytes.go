package convert

import (
	ps "github.com/takanoriyanagitani/go-pgcopy2sql"
)

func BytesFromBytesNew() func([]byte) (*ps.PgcopyBytes, error) {
	var buf ps.PgcopyBytes
	return func(b []byte) (*ps.PgcopyBytes, error) {
		buf.SetBytes(b)
		return &buf, nil
	}
}

func NullBytesFromBytesNew() func([]byte) (*ps.PgcopyBytesN, error) {
	var buf ps.PgcopyBytesN
	return func(b []byte) (*ps.PgcopyBytesN, error) {
		if nil == b {
			buf.Reset()
			return &buf, nil
		}
		buf.SetBytes(b)
		return &buf, nil
	}
}
