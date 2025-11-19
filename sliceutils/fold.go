package sliceutils

// FoldV provides an aggregation of all elements of a slice
// using an incremental aggregation function, which
// combines a slice element with an intermediate aggregation
// result starting with a given initial value.
func FoldV[S ~[]E, E any, F any](s S, init F, consume func(e E, acc F) F) F {
	acc := init
	for _, e := range s {
		acc = consume(e, acc)
	}
	return acc
}

// Fold provides an aggregation of all elements of a slice
// using an incremental aggregation function, which
// combines a slice element and its index with an intermediate aggregation
// result starting with a given initial value.
func FoldIV[S ~[]E, E any, F any](s S, init F, consume func(i int, e E, acc F) F) F {
	acc := init
	for i, e := range s {
		acc = consume(i, e, acc)
	}
	return acc
}
