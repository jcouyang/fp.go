package slices

func Alt[A any](a []A) func(b []A) []A {
	return func(b []A) []A {
		if len(a) == 0 {
			return b
		}
		return a
	}
}
