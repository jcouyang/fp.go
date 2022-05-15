package maps

// Concatenation of two maps, if any key duplicated, returns the value of the later map
func Concat[K comparable, V any](a map[K]V, b map[K]V) (c map[K]V) {
	c = make(map[K]V, len(a)+len(b))
	for k1, v1 := range a {
		c[k1] = v1
	}
	for k2, v2 := range b {
		c[k2] = v2
	}
	return
}
