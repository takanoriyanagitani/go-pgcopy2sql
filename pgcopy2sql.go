package pgcopy2sql

import (
	"database/sql"
	"errors"
	"fmt"
	"maps"
	"time"
)

var (
	ErrInvalidBool error = errors.New("invalid bool")

	ErrInvalidWord  error = errors.New("invalid word")
	ErrInvalidDword error = errors.New("invalid dword")
	ErrInvalidQword error = errors.New("invalid qword")

	ErrInvalidShort error = errors.New("invalid short")
	ErrInvalidInt   error = errors.New("invalid int")
	ErrInvalidLong  error = errors.New("invalid long")

	ErrInvalidFloat  error = errors.New("invalid float")
	ErrInvalidDouble error = errors.New("invalid double")

	ErrInvalidUuid error = errors.New("invalid uuid")

	ErrInvalidTime error = errors.New("invalid time")
)

type ColumnType string

const (
	ColTypUnknown ColumnType = "UNKNOWN"

	ColTypString ColumnType = "string"
	ColTypBytes  ColumnType = "bytes"
	ColTypShort  ColumnType = "short"
	ColTypInt    ColumnType = "int"
	ColTypLong   ColumnType = "long"
	ColTypBool   ColumnType = "boolean"
	ColTypFloat  ColumnType = "float"
	ColTypDouble ColumnType = "double"
	ColTypTime   ColumnType = "time"
	ColTypNull   ColumnType = "null"
	ColTypUuid   ColumnType = "uuid"

	ColTypStringN ColumnType = "string-null"
	ColTypShortN  ColumnType = "short-null"
	ColTypIntN    ColumnType = "int-null"
	ColTypLongN   ColumnType = "long-null"
	ColTypBoolN   ColumnType = "boolean-null"
	ColTypFloatN  ColumnType = "float-null"
	ColTypDoubleN ColumnType = "double-null"
	ColTypTimeN   ColumnType = "time-null"
	ColTypUuidN   ColumnType = "uuid-null"
	ColTypBytesN  ColumnType = "bytes-null"
)

//go:generate go run internal/gen/pgcopytype.go PgcopyBool bool false
//go:generate gofmt -s -w pgcopybool.go
type PgcopyBool struct{ value bool }

//go:generate go run internal/gen/pgcopytype.go PgcopyFloat float32 0.0
//go:generate gofmt -s -w pgcopyfloat.go
type PgcopyFloat struct{ value float32 }

//go:generate go run internal/gen/pgcopytype.go PgcopyDouble float64 0.0
//go:generate gofmt -s -w pgcopydouble.go
type PgcopyDouble struct{ value float64 }

//go:generate go run internal/gen/pgcopytype.go PgcopyLong int64 0
//go:generate gofmt -s -w pgcopylong.go
type PgcopyLong struct{ value int64 }

//go:generate go run internal/gen/pgcopytype.go PgcopyInt int32 0
//go:generate gofmt -s -w pgcopyint.go
type PgcopyInt struct{ value int32 }

//go:generate go run internal/gen/pgcopytype.go PgcopyShort int16 0
//go:generate gofmt -s -w pgcopyshort.go
type PgcopyShort struct{ value int16 }

//go:generate go run internal/gen/nullable/pgcopynull.go Bool bool Bool
//go:generate gofmt -s -w pgcopybooln.go
type PgcopyBoolN struct{ value sql.NullBool }

type PgcopyFloatN struct{ value sql.Null[float32] }

//go:generate go run internal/gen/nullable/pgcopynull.go Double float64 Float64
//go:generate gofmt -s -w pgcopydoublen.go
type PgcopyDoubleN struct{ value sql.NullFloat64 }

//go:generate go run internal/gen/nullable/pgcopynull.go Short int16 Int16
//go:generate gofmt -s -w pgcopyshortn.go
type PgcopyShortN struct{ value sql.NullInt16 }

//go:generate go run internal/gen/nullable/pgcopynull.go Int int32 Int32
//go:generate gofmt -s -w pgcopyintn.go
type PgcopyIntN struct{ value sql.NullInt32 }

//go:generate go run internal/gen/nullable/pgcopynull.go Long int64 Int64
//go:generate gofmt -s -w pgcopylongn.go
type PgcopyLongN struct{ value sql.NullInt64 }

func ComposeErr[T, U, V any](
	f func(T) (U, error),
	g func(U) (V, error),
) func(T) (V, error) {
	return func(t T) (v V, e error) {
		u, e := f(t)
		if nil != e {
			return v, e
		}
		return g(u)
	}
}

var coltyp2string map[ColumnType]string = maps.Collect(
	func(yield func(ColumnType, string) bool) {
		yield(ColTypString, string(ColTypString))
		yield(ColTypBytes, string(ColTypBytes))
		yield(ColTypBytesN, string(ColTypBytesN))
		yield(ColTypInt, string(ColTypInt))
		yield(ColTypLong, string(ColTypLong))
		yield(ColTypBool, string(ColTypBool))
		yield(ColTypFloat, string(ColTypFloat))
		yield(ColTypDouble, string(ColTypDouble))
		yield(ColTypTime, string(ColTypTime))
		yield(ColTypNull, string(ColTypNull))
		yield(ColTypUuid, string(ColTypUuid))
		yield(ColTypStringN, string(ColTypStringN))
		yield(ColTypIntN, string(ColTypIntN))
		yield(ColTypLongN, string(ColTypLongN))
		yield(ColTypBoolN, string(ColTypBoolN))
		yield(ColTypFloatN, string(ColTypFloatN))
		yield(ColTypDoubleN, string(ColTypDoubleN))
		yield(ColTypTimeN, string(ColTypTimeN))
		yield(ColTypUuidN, string(ColTypUuidN))
	},
)

var string2coltyp map[string]ColumnType = maps.Collect(
	func(yield func(string, ColumnType) bool) {
		for typ, s := range maps.All(coltyp2string) {
			yield(s, typ)
		}
	},
)

func (t ColumnType) String() string {
	s, found := coltyp2string[t]
	if found {
		return s
	}
	return "UNKNOWN COLUMN TYPE(BUG)"
}

func StringToColumnType(s string) (ColumnType, error) {
	typ, found := string2coltyp[s]
	switch found {
	case true:
		return typ, nil
	default:
		return ColTypUnknown, fmt.Errorf("invalid type: %s", s)
	}
}

func BoolFromBytes(b []byte) (bool, error) {
	if 1 != len(b) {
		return false, ErrInvalidBool
	}
	var u uint8 = b[0]
	var notTrue bool = 0 == u
	return !notTrue, nil
}

func UuidFromBytes(b []byte) ([16]byte, error) {
	if 16 != len(b) {
		return [16]byte{}, ErrInvalidUuid
	}
	var buf [16]byte
	copy(buf[:], b)
	return buf, nil
}

type ValueWriter interface {
	WriteString(s string) error
	WriteBytes(b []byte) error
	WriteShort(i int16) error
	WriteInt(i int32) error
	WriteLong(i int64) error
	WriteBool(b bool) error
	WriteFloat(f float32) error
	WriteDouble(f float64) error
	WriteTime(t time.Time) error
	WriteUuid(u [16]byte) error
	WriteNull() error

	WriteNullString(s sql.NullString) error
	WriteNullShort(i sql.NullInt16) error
	WriteNullInt(i sql.NullInt32) error
	WriteNullLong(l sql.NullInt64) error
	WriteNullBool(b sql.NullBool) error
	WriteNullFloat(f sql.Null[float32]) error
	WriteNullDouble(d sql.NullFloat64) error
	WriteNullUuid(u sql.Null[[16]byte]) error
	WriteNullTime(t sql.NullTime) error
}

type Value interface {
	WriteTo(writer ValueWriter) error
	Reset()
	IsNull() bool
	fmt.Stringer
}

type PgColumnCount struct{ raw int16 }

func (p PgColumnCount) Count() int16  { return p.raw }
func (p PgColumnCount) LastRow() bool { return p.raw < 0 }

type PgColumnSize struct{ raw int32 }

func (s PgColumnSize) Size() int32  { return s.raw }
func (s PgColumnSize) IsNull() bool { return s.raw < 0 }
