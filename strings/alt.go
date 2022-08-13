package strings

func Alt(a string) func(b string) string {
	return func(b string) string {
		if a == "" {
			return b
		}
		return a
	}
}
