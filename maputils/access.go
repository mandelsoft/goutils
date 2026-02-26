package maputils

func AndGet[MA ~map[A]B, A comparable, B any](m MA, key A) B {
	var _nil B
	if m == nil {
		return _nil
	}
	return m[key]
}

// Set(m)(a, AndSet(b, Value(c)))

func Value[A any](a A) func(A) A {
	return func(_ A) A {
		return a
	}
}

func Set[MA ~map[A]B, A comparable, B any](m MA) func(A, func(B) B) MA {
	return func(a A, f func(B) B) MA {
		return AndSet[MA](a, f)(m)
	}
}

func AndSet[MA ~map[A]B, A comparable, B any](a A, f func(B) B) func(MA) MA {
	return func(m MA) MA {
		if m == nil {
			m = MA{}
		}
		m[a] = f(m[a])
		return m
	}
}
