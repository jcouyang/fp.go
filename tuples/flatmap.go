package tuples

func FlatMap[A any, B any](fn func(a A) (B, error)) func(a A, e error) (B, error) {
	return func(a A, e error) (B, error) {
		if e != nil {
			return *new(B), e
		}
		return fn(a)
	}
}
