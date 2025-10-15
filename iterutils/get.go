package iterutils

import "iter"

func Get[T any](seq iter.Seq[T]) []T {
	list := []T{}
	for e := range seq {
		list = append(list, e)
	}
	return list
}
