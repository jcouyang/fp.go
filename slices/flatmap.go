package slices

import ("github.com/jcouyang/fizpop")

// Flatten a nested slice
func Flatten[A any](a [][]A) (b[]A) {
	for _, i := range(a) {
		for _, j := range(i) {
			b = append(b, j)
		}
	}
	return
}

// Map a function over all the elements of a slice and concatenate the resulting slice
func FlatMap[A any, B any](fn func (a A) []B) func (a []A) []B {
	return fp.Compose(Flatten[B], Map(fn))
}
