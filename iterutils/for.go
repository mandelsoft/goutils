package iterutils

import "iter"

func For[T any](v ...T) iter.Seq[T] {
	return func(yield func(v T) bool) {
		for _, e := range v {
			if !yield(e) {
				return
			}
		}
	}
}
