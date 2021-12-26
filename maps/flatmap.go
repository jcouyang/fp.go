package maps

func Flatten[K comparable, A any] (a map[K]map[K]A) (b map[K]A) {
	b = make(map[K]A)
	for _,j := range a {
		for k,l := range j {
			b[k] = l
		}
	}
	return
}

func Flatmap[K comparable, A any,  B any](fn func (a A) map[K]B) func (a  map[K]A) map[K]B {
	return func (a map[K]A) map[K]B {
		b := make(map[K]map[K]B)
		for i,j := range a {
			b[i] = fn(j)
		}
		return Flatten(b)
	}
}
