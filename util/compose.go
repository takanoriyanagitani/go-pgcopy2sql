package util

import (
	ps "github.com/takanoriyanagitani/go-pgcopy2sql"
)

func ComposeErr[T, U, V any](
	f func(T) (U, error),
	g func(U) (V, error),
) func(T) (V, error) {
	return ps.ComposeErr(f, g)
}
