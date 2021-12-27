package slices

// Take first n elements of a slice. if len(a) < n return len(a) elements
func Take[A any](n int) func (a []A) []A {
	return func(a []A) []A {
		if len(a) < n {
			return a[:]
		}
		return a[:n]
	}
	
}

// Take the longest prefix (possibly empty) of elements that satisfy fn:
func TakeWhile[A any](fn func(i A) bool) func (a []A) []A {
	return func(a []A) []A {
		cursor := 0
		for _,i := range a {
			if fn(i) {
				cursor++
			}
		}
		return a[:cursor]
	}
}
