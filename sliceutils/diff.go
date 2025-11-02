package sliceutils

import (
	"github.com/mandelsoft/goutils/general"
	"slices"
)

func Diff[S ~[]E, E comparable](a, b S) S {
	var diff S
	for _, o := range a {
		if !slices.Contains(b, o) {
			diff = append(diff, o)
		}
	}
	return diff
}

func DiffFunc[S ~[]E, E comparable](a, b S, eq general.EqualsFunc[E]) S {
	var diff S
	for _, o := range a {
		if !slices.ContainsFunc(b, func(e E) bool { return eq(o, e) }) {
			diff = append(diff, o)
		}
	}
	return diff
}
