package convert

import (
	ps "github.com/takanoriyanagitani/go-pgcopy2sql"
)

func UuidFromBytesNew() func([]byte) (*ps.PgcopyUuid, error) {
	var buf ps.PgcopyUuid
	return func(b []byte) (*ps.PgcopyUuid, error) {
		buf.Reset()
		i, e := ps.UuidFromBytes(b)
		if nil == e {
			buf.SetValue(i)
		}
		return &buf, e
	}
}

func UuidFromBytesNewN() func([]byte) (*ps.PgcopyUuidN, error) {
	var buf ps.PgcopyUuidN
	return func(b []byte) (*ps.PgcopyUuidN, error) {
		buf.Reset()
		if nil == b {
			return &buf, nil
		}

		i, e := ps.UuidFromBytes(b)
		if nil == e {
			buf.SetValue(i)
		}
		return &buf, e
	}
}
