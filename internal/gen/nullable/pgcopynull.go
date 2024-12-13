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
var primitiveName util.IO[string] = ArgByIndex(2)
var nullTypeName util.IO[string] = ArgByIndex(3)

var typeName util.IO[string] = util.Bind(
	typeHint,
	util.Lift(func(s string) (string, error) {
		return "Pgcopy" + s + "N", nil
	}),
)
var writeMethod util.IO[string] = util.Bind(
	typeHint,
	util.Lift(func(s string) (string, error) {
		return "WriteNull" + s, nil
	}),
)

type TypeConfig struct {
	TypeName        string
	PrimitiveName   string
	WriteMethodName string
	NullTypeName    string

	Prefix string
}

func (t TypeConfig) ToPrefix() string {
	return strings.TrimPrefix(t.TypeName, "Pgcopy")
}
func (t TypeConfig) PopulatePrefix() TypeConfig {
	t.Prefix = t.ToPrefix()
	return t
}

var typeConfig util.IO[TypeConfig] = util.Bind(
	util.All([]util.IO[string]{
		typeName,
		primitiveName,
		writeMethod,
		nullTypeName,
	}),
	util.Lift(func(vals []string) (TypeConfig, error) {
		return TypeConfig{
			TypeName:        vals[0],
			PrimitiveName:   vals[1],
			WriteMethodName: vals[2],
			NullTypeName:    vals[3],
		}.PopulatePrefix(), nil
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
	"./internal/gen/nullable/pgcopynull.tmpl",
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
