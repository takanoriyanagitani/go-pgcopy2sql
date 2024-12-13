package pgcopy2sql

import (
	"encoding/binary"
	"math"
	"time"
)

func BytesToQWORD(b []byte) (uint64, error) {
	if 8 != len(b) {
		return 0, ErrInvalidQword
	}
	var buf [8]byte
	copy(buf[:], b)
	return binary.BigEndian.Uint64(buf[:]), nil
}

func LongFromQWORD(u uint64) (int64, error) { return int64(u), nil }

func DoubleFromQWORD(u uint64) (float64, error) {
	return math.Float64frombits(u), nil
}

const PgTimeToUnixtimeUs int64 = 10957 * 86400 * 1000 * 1000

func TimeFromQWORD(u uint64) (time.Time, error) {
	var pgMicroTime int64 = int64(u)
	var unixtimeMicros int64 = pgMicroTime + PgTimeToUnixtimeUs
	return time.UnixMicro(unixtimeMicros), nil
}

var LongFromBytes func([]byte) (int64, error) = ComposeErr(
	BytesToQWORD,
	LongFromQWORD,
)

var DoubleFromBytes func([]byte) (float64, error) = ComposeErr(
	BytesToQWORD,
	DoubleFromQWORD,
)

var TimeFromBytes func([]byte) (time.Time, error) = ComposeErr(
	BytesToQWORD,
	TimeFromQWORD,
)
