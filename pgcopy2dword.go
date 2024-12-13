package pgcopy2sql

import (
	"encoding/binary"
	"math"
)

func BytesToDWORD(b []byte) (uint32, error) {
	if 4 != len(b) {
		return 0, ErrInvalidDword
	}
	var buf [4]byte
	copy(buf[:], b)
	return binary.BigEndian.Uint32(buf[:]), nil
}

func IntFromDWORD(d uint32) (int32, error) { return int32(d), nil }

func FloatFromDWORD(d uint32) (float32, error) {
	return math.Float32frombits(d), nil
}

var IntFromBytes func([]byte) (int32, error) = ComposeErr(
	BytesToDWORD,
	IntFromDWORD,
)

var FloatFromBytes func([]byte) (float32, error) = ComposeErr(
	BytesToDWORD,
	FloatFromDWORD,
)

func PgColumnSizeFromRawBytes(b [4]byte) PgColumnSize {
	raw, _ := IntFromBytes(b[:])
	return PgColumnSize{raw}
}
