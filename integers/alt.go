package integers

import "golang.org/x/exp/constraints"

func Alt[A constraints.Integer](a A) func(b A) A {
	return func(b A) A {
		if a == 0 {
			return b
		}
		return a
	}
}
