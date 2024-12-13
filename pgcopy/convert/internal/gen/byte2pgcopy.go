//go:build ignore

package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"text/template"

	util "github.com/takanoriyanagitani/go-pgcopy2sql/util"
)

var argLen int = len(os.Args)

func ArgByIndex(ix int) util.IO[string] {
	return func(_ context.Context) (string, error) {
		switch ix < argLen {
		case true:
			return os.Args[ix], nil
		default:
			return "", fmt.Errorf("no argument with index %v", ix)
		}
	}
}

var typeHint util.IO[string] = ArgByIndex(1)

var typeName util.IO[string] = util.Bind(
	typeHint,
	util.Lift(func(s string) (string, error) { return "PgcopyBytes" + s, nil }),
)
var colTypName util.IO[string] = util.Bind(
	typeHint,
	util.Lift(func(s string) (string, error) { return "ColTyp" + s, nil }),
)
var pgcopyTypeName util.IO[string] = util.Bind(
	typeHint,
	util.Lift(func(s string) (string, error) { return "Pgcopy" + s, nil }),
)
var pgcopyNewName util.IO[string] = util.Bind(
	typeHint,
	util.Lift(func(s string) (string, error) { return s + "FromBytesNew", nil }),
)

type TypeConfig struct {
	TypeHint       string
	TypeName       string
	ColTypName     string
	PgcopyTypeName string
	PgcopyNewName  string
}

var typeConfig util.IO[TypeConfig] = util.Bind(
	util.All([]util.IO[string]{
		typeName,
		colTypName,
		pgcopyTypeName,
		pgcopyNewName,
		typeHint,
	}),
	util.Lift(func(vals []string) (TypeConfig, error) {
		return TypeConfig{
			TypeHint:       vals[4],
			TypeName:       vals[0],
			ColTypName:     vals[1],
			PgcopyTypeName: vals[2],
			PgcopyNewName:  vals[3],
		}, nil
	}),
)

var filename util.IO[string] = util.Bind(
	typeName,
	util.Lift(func(s string) (string, error) {
		return strings.ToLower(s) + ".go", nil
	}),
)

func Must[T any](t T, e error) T {
	if nil != e {
		panic(e)
	}
	return t
}

var tmpl *template.Template = Must(template.ParseFiles(
	"./internal/gen/byte2pgcopy.tmpl",
))

func ApplyTemplate(
	filename string,
	t *template.Template,
	c TypeConfig,
) error {
	var f *os.File = Must(os.Create(filename))
	defer f.Close()

	var bw *bufio.Writer = bufio.NewWriter(f)
	defer bw.Flush()

	e := t.Execute(bw, c)
	if nil != e {
		return e
	}

	return bw.Flush()
}

func configToTemplateToFile(
	filename string,
) func(TypeConfig) util.IO[util.Void] {
	return func(c TypeConfig) util.IO[util.Void] {
		return func(_ context.Context) (util.Void, error) {
			return util.Empty, ApplyTemplate(filename, tmpl, c)
		}
	}
}

var filename2config2template2file util.IO[util.Void] = util.Bind(
	filename,
	func(name string) util.IO[util.Void] {
		return util.Bind(
			typeConfig,
			configToTemplateToFile(name),
		)
	},
)

var sub util.IO[util.Void] = func(ctx context.Context) (util.Void, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	_, e := filename2config2template2file(ctx)
	return util.Empty, e
}

func main() {
	_, e := sub(context.Background())
	if nil != e {
		panic(e)
	}
}
