package transformer

// Transformer is an interface for a function mapping one element to another
// with potentially different types.
type Transformer[I, O any] = func(in I) O

func Identity[T any](in T) T {
	return in
}

func KeyToValue[M ~map[K]V, K comparable, V any](m M) Transformer[K, V] {
	return func(k K) V {
		return m[k]
	}
}

func IndexToValue[S ~[]V, V any](s S) Transformer[int, V] {
	return func(i int) V {
		return s[i]
	}
}

// OptimizedTransform transforms an element to another element with potentially
// different types using a Transformer. But the transformer is only called if
// the inbound element has not yet satisfied the outbound type
func OptimizedTransform[I, O any](in I, conv Transformer[I, O]) O {
	if c, ok := any(in).(O); ok {
		return c
	} else {
		return conv(in)
	}
}

// Optimized maps a regular transformer into an optimized transformer
// using OptimizedTransform.
func Optimized[I, O any](t Transformer[I, O]) Transformer[I, O] {
	return func(i I) O {
		return OptimizedTransform[I, O](i, t)
	}
}
