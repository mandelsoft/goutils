package atomic

import (
	"sync/atomic"

	"github.com/mandelsoft/goutils/generics"
)

type Value[T any] struct {
	atomic.Value
}

func (v *Value[T]) Load() T {
	return generics.Cast[T](v.Value.Load())
}

func (v *Value[T]) Store(new T) {
	v.Value.Store(new)
}

func (v *Value[T]) Swap(new T) T {
	return generics.Cast[T](v.Value.Swap(new))
}

// CompareAndSwap swaps for matching old value.
// We have to deal with real nil and typed nil for parameter
// old.
// Therefore, we handle typed nil pointers, separately.
// We cannot use type T for old, because
// then it is not possible to pass a real nil for an initial Value.CompareAndSwap.
func (v *Value[T]) CompareAndSwap(old any, new T) bool {
	if b := v.Value.CompareAndSwap(old, new); b || old != nil {
		return b
	}
	var _nil T
	old = _nil
	return v.Value.CompareAndSwap(old, new)
}
