package slices

import (
	"github.com/jcouyang/fp.go/arrow"
)

// Flatten a nested slice
func Flatten[A any](a [][]A) (b []A) {
	for _, i := range a {
		for _, j := range i {
			b = append(b, j)
		}
	}
	return
}

// Map a function over all the elements of a slice and concatenate the resulting slice
func FlatMap[A any, B any](fn func(a A) []B) func(a []A) []B {
	return arrow.Compose(Flatten[B], Map(fn))
}

// Map a function over all the elements of a slice and concatenate the resulting slice
func FlatMapE[A any, B any](fn func(a A) ([]B, error)) func(a []A) ([]B, error) {
	return func(a []A) ([]B, error) {
		return arrow.First[[][]B, []B, error](Flatten[B])(MapE(fn)(a))
	}
}
