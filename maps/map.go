package maps

import "constraints"

func Map[K constraints.Ordered, A any,  B any](fn func (a A) B) func (a  map[K]A) map[K]B {
	return func (a map[K]A) (b map[K]B) {
		b = make(map[K]B)
		for i,j := range a {
			b[i] = fn(j)
		}
		return
	}
}
