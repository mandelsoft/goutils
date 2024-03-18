package atomic

import (
	"sync/atomic"
)

type Value[T any] struct {
	atomic.Value
}

func (v *Value[T]) Load() T {
	return v.Value.Load().(T)
}

func (v *Value[T]) Store(new T) {
	v.Value.Store(new)
}

func (v *Value[T]) Swap(new T) T {
	return v.Value.Swap(new).(T)
}

func (v *Value[T]) CompareAndSwap(old, new T) bool {
	return v.Value.CompareAndSwap(old, new)
}
