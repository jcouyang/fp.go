package tuples

func Traverse[A any, B any](fn func(a A) (B, error)) func (a []A) ([]B, error) {
	return func(a []A) ([]B, error){
		b := make([]B, len(a))
		for i,j := range(a) {
			res, err := fn(j)
			if err != nil {
				return b, err
			} else {
				b[i] = res
			}
		}
		return b, nil
	}
}
