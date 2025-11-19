package maputils

import (
	"cmp"
	"maps"
	"slices"
)

// FoldV provides an aggregation of all elements of a map
// using an incremental aggregation function, which
// combines a map element with an intermediate aggregation
// result starting with a given initial value.
func FoldV[M ~map[K]E, K comparable, E any, F any](m M, init F, consume func(e E, acc F) F) F {
	acc := init
	for _, e := range m {
		acc = consume(e, acc)
	}
	return acc
}

// FoldKV provides an aggregation of all elements of a map
// using an incremental aggregation function, which
// combines a slice element and its key with an intermediate aggregation
// result starting with a given initial value.
func FoldKV[M ~map[K]E, K comparable, E any, F any](m M, init F, consume func(k K, e E, acc F) F) F {
	acc := init
	for k, e := range m {
		acc = consume(k, e, acc)
	}
	return acc
}

// FoldSV provides an aggregation of all elements of a map
// using an incremental aggregation function, which
// combines a map element with an intermediate aggregation
// result starting with a given initial value.
// The entry order is determined by the ordered values.
func FoldSV[M ~map[K]E, K cmp.Ordered, E any, F any](m M, init F, consume func(e E, acc F) F) F {
	keys := slices.Sorted(maps.Keys(m))
	acc := init
	for _, k := range keys {
		acc = consume(m[k], acc)
	}
	return acc
}

// FoldSKV provides an aggregation of all elements of a map
// using an incremental aggregation function, which
// combines a slice element and its key with an intermediate aggregation
// result starting with a given initial value.
// The entry order is determined by the ordered key values.
func FoldSKV[M ~map[K]E, K cmp.Ordered, E any, F any](m M, init F, consume func(k K, e E, acc F) F) F {
	keys := slices.Sorted(maps.Keys(m))
	acc := init
	for _, k := range keys {
		acc = consume(k, m[k], acc)
	}
	return acc
}

// FoldSVFunc provides an aggregation of all elements of a map
// using an incremental aggregation function, which
// combines a map element with an intermediate aggregation
// result starting with a given initial value.
// The entry order is determined by the sort function on the
// key values.
func FoldSVFunc[M ~map[K]E, K cmp.Ordered, E any, F any](m M, init F, consume func(e E, acc F) F, cmp CompareFunc[K]) F {
	keys := slices.SortedFunc(maps.Keys(m), cmp)
	acc := init
	for _, k := range keys {
		acc = consume(m[k], acc)
	}
	return acc
}

// FoldSKVFunc provides an aggregation of all elements of a map
// using an incremental aggregation function, which
// combines a slice element and its key with an intermediate aggregation
// result starting with a given initial value.
// The entry order is determined by the sort function on the
// key values.
func FoldSKVFunc[M ~map[K]E, K cmp.Ordered, E any, F any](m M, init F, consume func(k K, e E, acc F) F, cmp CompareFunc[K]) F {
	keys := slices.SortedFunc(maps.Keys(m), cmp)
	acc := init
	for _, k := range keys {
		acc = consume(k, m[k], acc)
	}
	return acc
}
