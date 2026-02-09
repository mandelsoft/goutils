package funcs

import (
	"fmt"
	"github.com/mandelsoft/goutils/general"
)

// Option represents an optional resuölt of a function.
// If can represent a result value (of type T) or no result value
// potentially with an error.
type Option[T any] interface {
	// IsNone indicates whether a value is provided or not.
	IsNone() bool
	// Value provides the value of an Option.
	// If no value is provided, a panic is called.
	// It may only be called if IsNone() is false.
	Value() T
	// Error provides an error is IsNone is false.
	Error() error
}

type some[T any] struct {
	value T
}

func (_ some[T]) IsNone() bool {
	return false
}

func (o some[T]) Value() T {
	return o.value
}

func (o some[T]) Error() error {
	return nil
}

// Some provides an Option for a given value.
func Some[T any](v T) Option[T] {
	return some[T]{v}
}

type none[T any] struct {
	err error
}

func (o none[T]) IsNone() bool {
	return true
}

func (o none[T]) Value() T {
	panic("called Option.Value of error value")
}

func (o none[T]) Error() error {
	return o.err
}

// None provides an Option without a value,
// but optionally with an explicit error.
func None[T any](err ...error) Option[T] {
	e := general.Optional(err...)
	if e == nil {
		e = fmt.Errorf("no value provided")
	}
	return none[T]{err: e}
}

// AndThen chains an operation on a given Option result,
// for example, if f and g are two operations providing an optional result
// the AndThen(f(), g) provides a result if f and g provide a result, whereas
// g is only called and fed with the result of f if f provided a result.
// Option together with AndThen enables chaining of function calls providing
// Nil results or error results.
func AndThen[A, B any](a Option[A], f func(A) Option[B]) Option[B] {
	if a.IsNone() {
		return None[B](a.Error())
	}
	return f(a.Value())
}
