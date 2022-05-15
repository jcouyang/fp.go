package slices

// Fold a slice from left
func Foldl[A any, B any](fold func(b B, a A) B, dflt B) func(a []A) B {
	return func(a []A) (b B) {
		if len(a) > 0 {
			b = dflt
			for _, i := range a {
				b = fold(b, i)
			}
			return
		}
		return dflt
	}
}

// Fold a slice from right
func Foldr[A any, B any](fold func(a A, b B) B, dflt B) func(a []A) B {
	return func(a []A) (b B) {
		if len(a) > 0 {
			b = dflt
			for i := len(a) - 1; i >= 0; i-- {
				b = fold(a[i], b)
			}
			return
		}
		return dflt
	}
}
