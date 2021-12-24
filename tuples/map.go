package tuples

func Map[A any, B any](a A, e error) (func (fn func(a A) B) (B, error)) {
	return func(fn func(a A) B) (B, error){
		if e != nil {
			return *new(B), e
		}
		return fn(a), e
	}
}
