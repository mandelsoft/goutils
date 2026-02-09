package funcs

import "github.com/mandelsoft/goutils/generics"

// First returns the first result of a sequence of multiple function results.
func First[T any](v T, rest ...any) T {
	return v
}

// Second returns the second result of a sequence of multiple function results..
func Second[T any](a any, v T, rest ...any) T {
	return v
}

// Third returns the first result of a sequence of multiple function results.
func Third[T any](a, b any, v T, rest ...any) T {
	return v
}

// Fourth returns the first result of a sequence of multiple function results.
func Fourth[T any](a, b, c any, v T, rest ...any) T {
	return v
}

// Last returns the last result of a sequence of multiple function results.
func Last[T any](args ...any) T {
	return generics.Cast[T](args[len(args)-1])
}
