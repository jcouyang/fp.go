package fp

func Eq[A comparable](a A) func(b A) bool {
	return func(b A) bool {
		return a == b
	}
}

func Identity[A any](a A) A { return a }

func Identity2[A, B any](a A, b B) (A, B) { return a, b }

func Const[A any](a A) func(b A) A { return func(b A) A { return a } }

func Const2[A, B any](a A, b B) func(aa A, bb B) (A, B) {
	return func(aa A, bb B) (A, B) { return a, b }
}
