package maps

import (
	"oyanglul.us/fp.go"
)

// Flatten a nested map
func Flatten[K1, K2 comparable, A any](a map[K1]map[K2]A) (b map[K2]A) {
	b = make(map[K2]A)
	for _, j := range a {
		for k, l := range j {
			b[k] = l
		}
	}
	return
}

// Apply a function to each element of map, and flatten the nested map
func FlatMap[K1, K2 comparable, V1, V2 any](fn func(k K1, a V1) map[K2]V2) func(a map[K1]V1) map[K2]V2 {
	return fp.Compose(Flatten[K1, K2, V2], Map(fn))
}
