package pgcopy2sql

import (
	"encoding/binary"
)

func BytesToWORD(b []byte) (uint16, error) {
	if 2 != len(b) {
		return 0, ErrInvalidWord
	}
	var buf [2]byte
	copy(buf[:], b)
	return binary.BigEndian.Uint16(buf[:]), nil
}

func ShortFromWORD(u uint16) (int16, error) { return int16(u), nil }

var ShortFromBytes func([]byte) (int16, error) = ComposeErr(
	BytesToWORD,
	ShortFromWORD,
)

func PgColumnCountFromRawBytes(b [2]byte) PgColumnCount {
	raw, _ := ShortFromBytes(b[:])
	return PgColumnCount{raw}
}
