package pointers

func Alt[A any](a *A) func(b *A) *A {
	return func(b *A) *A {
		if a != nil {
			return a
		}
		return b
	}
}
