package iterutils

import "iter"

// Convert transforms an iterator of type I into an iterator of type O by applying a type assertion during iteration.
func Convert[O, I any](in iter.Seq[I]) iter.Seq[O] {
	return func(yield func(O) bool) {
		in(func(v I) bool {
			var o any = v
			return yield(o.(O))
		})
	}
}
