package slices

// Extract the elements after the head of a slice. Tail of empty slice is still empty slice
func Tail[A any](a []A) []A {
	if len(a) > 0 {
		return a[1:]
	}
	return a
}
