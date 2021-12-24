package slices

// Concatenation of two slices
func Concat[A any](a []A, b []A) (c []A) {
	c = make([]A, len(a) + len(b))
	copy(c, a)
	copy(c[len(a):], b)
	return
}
