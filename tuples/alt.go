package tuples

import "github.com/jcouyang/fp.go"

func Alt[A any](a A, err error) func(b A, err error) (A, error) {
	if err != nil {
		return fp.Identity2[A, error]
	}
	return fp.Const2(a, err)
}
