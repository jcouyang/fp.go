package tuples

import "fmt"

func Map[A any, B any](fn func(a A) B) func (a A, e error) (B, error) {
	return func(a A, e error) (B, error){
		if e != nil {
			return *new(B), e
		}
		return fn(a), e
	}
}

func Map2[A any, B any, C any](fn func(a A, b B) C) func (a A, e error) func (B, error) (C, error) {
	return func(a A, e1 error) func (B, error) (C, error) {
		return func(b B, e2 error) (c C, e3 error) {
			if e1 != nil {
				e3 = e1
				if e2 != nil {
					e3 = fmt.Errorf("%v;%w", e1, e2)
				}
				return *new(C), e3
			}
			return fn(a, b), nil
		}
	}
}
