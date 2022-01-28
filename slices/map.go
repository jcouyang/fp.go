package slices

// Apply a function to each element of slice
func Map[A, B any](fn func (a A) B) func (a []A) []B {
	return func (a []A) (b []B) {
		b = make([]B, len(a))
		for i,j := range(a) {
			b[i] = fn(j)
		}
		return
	}
}

// Apply an effectful function to each element of slice, return when error happen
func MapE[A, B any](fn func(a A) (B, error)) func (a []A) ([]B, error) {
	return func (a []A) (b []B, e error) {
		b = make([]B, len(a))
		for i,j := range(a) {
			b[i], e = fn(j)
			if e != nil { return }
		}
		return
	}	
}
