package maps

// Returns the map of those elements that satisfy the predicate
func Filter[K comparable, A any](fn func (k K,v A) bool) func (a  map[K]A) map[K]A {
	return func (a map[K]A) (b map[K]A) {
		b = make(map[K]A)
		for i,j := range a {
			if fn(i, j) {
				b[i] = j
			}
		}
		return
	}
}
