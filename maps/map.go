package maps

// Apply a function to each element of map
func Map[K comparable, A any,  B any](fn func (k K, a A) B) func (a  map[K]A) map[K]B {
	return func (a map[K]A) (b map[K]B) {
		b = make(map[K]B)
		for i,j := range a {
			b[i] = fn(i, j)
		}
		return
	}
}
