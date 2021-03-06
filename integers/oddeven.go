package integers

import "golang.org/x/exp/constraints"

func Odd[A constraints.Integer](a A) bool {
	return a%2 != 0
}

func Even[A constraints.Integer](a A) bool {
	return a%2 == 0
}
