package slices

func Reverse[A any](a []A) (b []A) {
	b = make([]A, len(a))
	for i, j := 0, len(a)-1; j >= 0; i,j = i+1, j-1 {
		b[i] = a[j]
	}
	return b
}

func ReverseInline[A any](a []A) []A {
	for i, j := 0, len(a)-1; i < j; i,j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}
