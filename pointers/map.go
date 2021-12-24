package pointers

func Map[A any, B any](fn func (a *A) *B) func (a  *A) *B {
	return func (a *A) *B {
		if a != nil {
			return fn(a)
		}
		return nil
	}
}
