package convert

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"iter"
	"log"
	"maps"
	"strings"

	ps "github.com/takanoriyanagitani/go-pgcopy2sql"
	util "github.com/takanoriyanagitani/go-pgcopy2sql/util"
)

var (
	ErrUnsupportedType   error = errors.New("unsupported type")
	ErrConverterNotFound error = errors.New("converter missing")
)

type ConvertConfig struct {
	stringChecker func(string) error
}

func (c ConvertConfig) WithStringChecker(
	checker func(string) error,
) ConvertConfig {
	c.stringChecker = checker
	return c
}

var ConvertConfigDefault ConvertConfig = ConvertConfig{}.
	WithStringChecker(CheckStringUtf8)

type Input struct {
	ps.ColumnType
	Raw []byte
}

type CfgToConv func(ConvertConfig) func([]byte) util.IO[ps.Value]

//go:generate go run internal/gen/byte2pgcopy.go Bool
//go:generate gofmt -s -w pgcopybytesbool.go
type PgcopyBytesBool []byte

//go:generate go run internal/gen/byte2pgcopy.go Short
//go:generate gofmt -s -w pgcopybytesshort.go
type PgcopyBytesShort []byte

//go:generate go run internal/gen/byte2pgcopy.go Int
//go:generate gofmt -s -w pgcopybytesint.go
type PgcopyBytesInt []byte

//go:generate go run internal/gen/byte2pgcopy.go Long
//go:generate gofmt -s -w pgcopybyteslong.go
type PgcopyBytesLong []byte

//go:generate go run internal/gen/byte2pgcopy.go Float
//go:generate gofmt -s -w pgcopybytesfloat.go
type PgcopyBytesFloat []byte

//go:generate go run internal/gen/byte2pgcopy.go Double
//go:generate gofmt -s -w pgcopybytesdouble.go
type PgcopyBytesDouble []byte

var cmap map[ps.ColumnType]CfgToConv = map[ps.ColumnType]CfgToConv{
	ps.ColTypBool:  ConfigToConverterBool,
	ps.ColTypBoolN: ConfigToConverterBoolN,

	ps.ColTypShort:  ConfigToConverterShort,
	ps.ColTypShortN: ConfigToConverterShortN,

	ps.ColTypInt:  ConfigToConverterInt,
	ps.ColTypIntN: ConfigToConverterIntN,

	ps.ColTypLong:  ConfigToConverterLong,
	ps.ColTypLongN: ConfigToConverterLongN,

	ps.ColTypFloat:  ConfigToConverterFloat,
	ps.ColTypFloatN: ConfigToConverterFloatN,

	ps.ColTypDouble:  ConfigToConverterDouble,
	ps.ColTypDoubleN: ConfigToConverterDoubleN,

	ps.ColTypUuid: func(_ ConvertConfig) func([]byte) util.IO[ps.Value] {
		var conv func([]byte) (*ps.PgcopyUuid, error) = UuidFromBytesNew()
		return func(b []byte) util.IO[ps.Value] {
			return func(_ context.Context) (ps.Value, error) {
				return conv(b)
			}
		}
	},

	ps.ColTypUuidN: func(_ ConvertConfig) func([]byte) util.IO[ps.Value] {
		var conv func([]byte) (*ps.PgcopyUuidN, error) = UuidFromBytesNewN()
		return func(b []byte) util.IO[ps.Value] {
			return func(_ context.Context) (ps.Value, error) {
				return conv(b)
			}
		}
	},

	ps.ColTypTime: func(_ ConvertConfig) func([]byte) util.IO[ps.Value] {
		var conv func([]byte) (*ps.PgcopyTime, error) = TimeFromBytesNew()
		return func(b []byte) util.IO[ps.Value] {
			return func(_ context.Context) (ps.Value, error) {
				return conv(b)
			}
		}
	},

	ps.ColTypTimeN: func(_ ConvertConfig) func([]byte) util.IO[ps.Value] {
		var conv func([]byte) (*ps.PgcopyTimeN, error) = TimeFromBytesNewN()
		return func(b []byte) util.IO[ps.Value] {
			return func(_ context.Context) (ps.Value, error) {
				return conv(b)
			}
		}
	},

	ps.ColTypString: func(cfg ConvertConfig) func([]byte) util.IO[ps.Value] {
		var conv func([]byte) (*ps.PgcopyString, error) = StringFromBytesNew(
			cfg.stringChecker,
		)
		return func(b []byte) util.IO[ps.Value] {
			return func(_ context.Context) (ps.Value, error) {
				return conv(b)
			}
		}
	},

	ps.ColTypStringN: func(cfg ConvertConfig) func([]byte) util.IO[ps.Value] {
		var conv func(
			[]byte,
		) (*ps.PgcopyStringN, error) = NullStringFromBytesNew(
			cfg.stringChecker,
		)
		return func(b []byte) util.IO[ps.Value] {
			return func(_ context.Context) (ps.Value, error) {
				return conv(b)
			}
		}
	},

	ps.ColTypBytes: func(_ ConvertConfig) func([]byte) util.IO[ps.Value] {
		var conv func([]byte) (*ps.PgcopyBytes, error) = BytesFromBytesNew()
		return func(b []byte) util.IO[ps.Value] {
			return func(_ context.Context) (ps.Value, error) {
				return conv(b)
			}
		}
	},

	ps.ColTypBytesN: func(_ ConvertConfig) func([]byte) util.IO[ps.Value] {
		var conv func(
			[]byte,
		) (*ps.PgcopyBytesN, error) = NullBytesFromBytesNew()
		return func(b []byte) util.IO[ps.Value] {
			return func(_ context.Context) (ps.Value, error) {
				return conv(b)
			}
		}
	},
}

func TypeToConverter(
	typ ps.ColumnType,
	cfg ConvertConfig,
) func([]byte) util.IO[ps.Value] {
	cfg2conv, found := cmap[typ]
	if !found {
		return func(_ []byte) util.IO[ps.Value] {
			return util.Err[ps.Value](
				fmt.Errorf("%w: %v", ErrUnsupportedType, typ),
			)
		}
	}
	return cfg2conv(cfg)
}

type TypeInfo map[int16]ps.ColumnType

type ConverterIx func(int16, []byte) util.IO[ps.Value]

func (t TypeInfo) ToConverterIx(cfg ConvertConfig) ConverterIx {
	i := func(yield func(int16, func([]byte) util.IO[ps.Value]) bool) {
		for ix, typ := range maps.All(t) {
			var conv func([]byte) util.IO[ps.Value] = TypeToConverter(
				typ,
				cfg,
			)
			yield(ix, conv)
		}
	}
	var m map[int16]func([]byte) util.IO[ps.Value] = maps.Collect(i)
	return func(ix int16, raw []byte) util.IO[ps.Value] {
		conv, found := m[ix]
		if !found {
			return util.Err[ps.Value](
				fmt.Errorf("%w: ix=%v", ErrConverterNotFound, ix),
			)
		}
		return conv(raw)
	}
}

func (t TypeInfo) ToConverterIxDefault() ConverterIx {
	return t.ToConverterIx(ConvertConfigDefault)
}

func TypeInfoFromIter(
	i iter.Seq2[int16, ps.ColumnType],
) TypeInfo {
	return maps.Collect(i)
}

type IndexedTypeInfo struct {
	Index int16  `json:"index"`
	Type  string `json:"type"`
}

func TypeInfoFromIndexIter(
	i iter.Seq[IndexedTypeInfo],
) TypeInfo {
	var mapd iter.Seq2[int16, ps.ColumnType] = func(
		yield func(int16, ps.ColumnType) bool,
	) {
		for it := range i {
			var ix int16 = it.Index
			var typ string = it.Type
			parsed, e := ps.StringToColumnType(typ)
			if nil == e {
				yield(ix, parsed)
			}
		}
	}
	return TypeInfoFromIter(mapd)
}

// Gets type info from env vars and keys of env vars.
//
// # Arguments
//   - getenv: Gets env var.
//   - keys: env var keys for vars which contains type string.
func TypeInfoFromEnvVarAndKeys(
	ctx context.Context,
	getenv func(string) util.IO[string],
	keys []string,
) TypeInfo {
	var i iter.Seq[IndexedTypeInfo] = func(
		yield func(IndexedTypeInfo) bool,
	) {
		for i, key := range keys {
			var ix int16 = int16(i)
			val, e := getenv(key)(ctx)
			if nil != e {
				log.Printf("%v\n", e)
				continue
			}

			tinfo := IndexedTypeInfo{
				Index: ix,
				Type:  val,
			}

			yield(tinfo)
		}
	}
	return TypeInfoFromIndexIter(i)
}

// Gets type info from env vars.
//
// # Arguments
//   - getenv: Gets env var.
//   - keysKey: env var key of keys of env vars.
func TypeInfoFromEnvVar(
	ctx context.Context,
	getenv func(string) util.IO[string],
	keysKey string,
) TypeInfo {
	keys, e := getenv(keysKey)(ctx)
	if nil != e {
		return map[int16]ps.ColumnType{}
	}
	var splited []string = strings.Split(keys, ",")
	return TypeInfoFromEnvVarAndKeys(ctx, getenv, splited)
}

const (
	KeysKeyDefault string = "ENV_KEYS_KEY"
)

// Gets type info from env vars using default env vars key.
//
// # Arguments
//   - getenv: Gets env var.
func TypeInfoFromEnvVarDefault(
	ctx context.Context,
	getenv func(string) util.IO[string],
) TypeInfo {
	return TypeInfoFromEnvVar(ctx, getenv, KeysKeyDefault)
}

// Gets type info from strings(JSONs).
//
// JSON sample:
//
//	{"index":0, "type": "string"}
//	{"index":1, "type": "string-null"}
//	{"index":2, "type": "int"}
//	{"index":3, "type": "int-null"}
func TypeInfoFromJsonIter(
	i iter.Seq[string],
) TypeInfo {
	var mapd iter.Seq[IndexedTypeInfo] = func(
		yield func(IndexedTypeInfo) bool,
	) {
		var info IndexedTypeInfo
		var buf bytes.Buffer
		for jstr := range i {
			buf.Reset()
			_, _ = buf.WriteString(jstr) // error is always nil or panic
			e := json.Unmarshal(buf.Bytes(), &info)
			if nil == e {
				yield(info)
			}
		}
	}
	return TypeInfoFromIndexIter(mapd)
}

// Gets type info from the reader(JSON lines).
func TypeInfoFromJsonReadable(
	rdr io.Reader,
) TypeInfo {
	var s *bufio.Scanner = bufio.NewScanner(rdr)
	var i iter.Seq[string] = func(yield func(string) bool) {
		for s.Scan() {
			var line string = s.Text()
			yield(line)
		}
	}
	return TypeInfoFromJsonIter(i)
}
