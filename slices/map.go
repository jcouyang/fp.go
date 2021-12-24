package slices

func Map[A any, B any](fn func (a A) B) func (a []A) []B {
	return func (a []A) (b []B) {
		b = make([]B, len(a))
		for i,j := range(a) {
			b[i] = fn(j)
		}
		return
	}
	
}
