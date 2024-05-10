package general

type CompareFunc[E any] func(a, b E) int

func EqualsForCompareFuncFor[T any](cmp CompareFunc[T]) func(a, b T) bool {
	return func(a, b T) bool {
		return cmp(a, b) == 0
	}
}

func ContainsForCompareFuncFor[T any](a T, cmp CompareFunc[T]) func(b T) bool {
	return func(b T) bool {
		return cmp(a, b) == 0
	}
}

type Equals[T any] interface {
	Equals(a T) bool
}

func EqualsFuncFor[E Equals[E]]() func(a, b E) bool {
	return func(a, b E) bool {
		return a.Equals(b)
	}
}

func ContainsFuncFor[E Equals[E]](a E) func(b E) bool {
	return func(b E) bool {
		return a.Equals(b)
	}
}
