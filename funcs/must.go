package funcs

import (
	"fmt"
)

// Must expect a result to be provided without error.
func Must[T any](o T, err error) T {
	if err != nil {
		panic(fmt.Errorf("expected a %T, but got error %w", o, err))
	}
	return o
}

func ErrorFrom[T any](t T, err error) error {
	return err
}

func ErrorFrom2[T, U any](t T, u U, err error) error {
	return err
}

func ErrorFrom3[T, U, V any](t T, u U, v V, err error) error {
	return err
}

// Error provides the last value, which must be of type error.
// It is expected to be called on the result of a function call providing
// an error return, whose results should be ignored beside the provided error.
// e.g.: Error(f())
func Error(args ...any) error {
	return Last[error](args...)
}
