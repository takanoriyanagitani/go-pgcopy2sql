package convert

import (
	ps "github.com/takanoriyanagitani/go-pgcopy2sql"
)

func TimeFromBytesNew() func([]byte) (*ps.PgcopyTime, error) {
	var buf ps.PgcopyTime
	return func(b []byte) (*ps.PgcopyTime, error) {
		buf.Reset()
		i, e := ps.TimeFromBytes(b)
		if nil == e {
			buf.SetValue(i)
		}
		return &buf, e
	}
}

func TimeFromBytesNewN() func([]byte) (*ps.PgcopyTimeN, error) {
	var buf ps.PgcopyTimeN
	return func(b []byte) (*ps.PgcopyTimeN, error) {
		buf.Reset()
		if nil == b {
			return &buf, nil
		}

		i, e := ps.TimeFromBytes(b)
		if nil == e {
			buf.SetValue(i)
		}
		return &buf, e
	}
}
