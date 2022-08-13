package maps

func Alt[K comparable, V any](a map[K]V) func(b map[K]V) map[K]V {
	return func(b map[K]V) map[K]V {
		if len(a) == 0 {
			return b
		}
		return a
	}
}
