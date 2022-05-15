package slices

// Returns the slice of those elements that satisfy the predicate
func Filter[A any](fn func(a A) bool) func(a []A) []A {
	return func(a []A) (b []A) {
		b = []A{}
		for _, i := range a {
			if fn(i) {
				b = append(b, i)
			}
		}
		return
	}
}
