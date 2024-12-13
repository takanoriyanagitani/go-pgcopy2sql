package main

import (
	"context"
	"fmt"
	"iter"
	"log"
	"os"

	ps "github.com/takanoriyanagitani/go-pgcopy2sql"
	util "github.com/takanoriyanagitani/go-pgcopy2sql/util"

	pc "github.com/takanoriyanagitani/go-pgcopy2sql/pgcopy/convert"
	ph "github.com/takanoriyanagitani/go-pgcopy2sql/pgcopy/header"
	pr "github.com/takanoriyanagitani/go-pgcopy2sql/pgcopy/reader2rows"
)

var EnvVarByKey func(string) util.IO[string] = util.Lift(
	func(key string) (string, error) {
		val, found := os.LookupEnv(key)
		switch found {
		case true:
			return val, nil
		default:
			return "", fmt.Errorf("env var %s missing", key)
		}
	},
)

var typInfo util.IO[pc.TypeInfo] = func(
	ctx context.Context,
) (pc.TypeInfo, error) {
	return pc.TypeInfoFromEnvVarDefault(ctx, EnvVarByKey), nil
}

var converterIx util.IO[pc.ConverterIx] = util.Bind(
	typInfo,
	util.Lift(func(t pc.TypeInfo) (pc.ConverterIx, error) {
		return t.ToConverterIxDefault(), nil
	}),
)

var header util.IO[ph.SimpleHeader] = ph.SimpleHeaderStdinDefault

var rows util.IO[iter.Seq2[[]ps.Value, error]] = util.Bind(
	header,
	func(_ ph.SimpleHeader) util.IO[iter.Seq2[[]ps.Value, error]] {
		return util.Bind(
			converterIx,
			pr.ConverterToStdinToRows,
		)
	},
)

func rowsSinkDummy(i iter.Seq2[[]ps.Value, error]) util.IO[util.Void] {
	return func(ctx context.Context) (util.Void, error) {
		for val, e := range i {
			select {
			case <-ctx.Done():
				return util.Empty, ctx.Err()
			default:
			}

			if nil != e {
				return util.Empty, e
			}

			for _, col := range val {
				fmt.Printf("col: %v\n", col)
			}
		}
		return util.Empty, nil
	}
}

var rows2sink util.IO[util.Void] = util.Bind(
	rows,
	rowsSinkDummy,
)

var sub util.IO[util.Void] = func(ctx context.Context) (util.Void, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	return rows2sink(ctx)
}

func main() {
	_, e := sub(context.Background())
	if nil != e {
		log.Printf("%v\n", e)
	}
}
