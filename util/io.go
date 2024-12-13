package util

import (
	"context"
)

type IO[T any] func(context.Context) (T, error)

func Of[T any](t T) IO[T] {
	return func(_ context.Context) (T, error) {
		return t, nil
	}
}

func Err[T any](err error) IO[T] {
	return func(_ context.Context) (t T, e error) {
		return t, err
	}
}

func Lift[T, U any](
	pure func(T) (U, error),
) func(T) IO[U] {
	return func(t T) IO[U] {
		return func(_ context.Context) (U, error) {
			return pure(t)
		}
	}
}

func Bind[T, U any](
	i IO[T],
	f func(T) IO[U],
) IO[U] {
	return func(ctx context.Context) (u U, e error) {
		t, e := i(ctx)
		if nil != e {
			return u, e
		}
		return f(t)(ctx)
	}
}

type Void struct{}

var Empty Void = struct{}{}

func All[T any](ios []IO[T]) IO[[]T] {
	return func(ctx context.Context) ([]T, error) {
		var ret []T = make([]T, 0, len(ios))
		for _, i := range ios {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			default:
			}

			t, e := i(ctx)
			if nil != e {
				return nil, e
			}

			ret = append(ret, t)
		}
		return ret, nil
	}
}
