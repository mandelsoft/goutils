package atomic

import (
	"sync/atomic"

	"github.com/mandelsoft/goutils/generics"
)

type value[T any] struct {
	value T
}

// InterfaceValue can be used for interface types
// to be able to store nil values.
// It uses an additional indirecttion to be able to
// hold a nil value. If this is not required, use Value.
type InterfaceValue[T any] struct {
	atomic.Value
}

func (v *InterfaceValue[T]) Load() T {
	var _nil T
	val := v.Value.Load()
	if val == nil {
		return _nil
	}
	return generics.Cast[value[T]](val).value
}

func (v *InterfaceValue[T]) Store(new T) {
	v.Value.Store(value[T]{new})
}

func (v *InterfaceValue[T]) Swap(new T) T {
	var _nil T
	val := v.Value.Swap(value[T]{new})
	if val == nil {
		return _nil
	}
	return generics.Cast[value[T]](val).value
}

// CompareAndSwap swaps for matching old value.
// We have to deal with real nil and typed nil for parameter
// old.
// Therefore, we handle typed nil pointers, separately.
// We cannot use type T for old, because
// then it is not possible to pass a real nil for an initial Value.CompareAndSwap.
func (v *InterfaceValue[T]) CompareAndSwap(old any, new T) bool {
	n := value[T]{new}
	o := value[T]{}
	if old == nil {
		// first try intentional nil value
		if b := v.Value.CompareAndSwap(o, n); b {
			return b
		}
		// second try initial
		return v.Value.CompareAndSwap(old, n)
	}
	// old value is non nil, always use ff value
	o.value = old.(T)
	return v.Value.CompareAndSwap(o, n)
}
