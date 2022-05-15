package fp

func Compose[A any, B any, C any](fn1 func(b B) C, fn2 func(a A) B) func(a A) C {
	return func(a A) C {
		return fn1(fn2(a))
	}
}

func Compose2[A any, B any, C any](fn1 func(b B) C, fn2 func(a A, d ...any) B) func(a A, d ...any) C {
	return func(a A, d ...any) C {
		return fn1(fn2(a, d...))
	}
}

func Pipe[A any, B any, C any](fn1 func(a A) B, fn2 func(b B) C) func(a A) C {
	return func(a A) C {
		return fn2(fn1(a))
	}
}

func Eq[A comparable](a A) func(b A) bool {
	return func(b A) bool {
		return a == b
	}
}

// Send the first component of the input through the argument arrow, and copy the rest unchanged to the output.
func First[A, B, C any](fn func(a A) B) func(a A, c C) (B, C) {
	return func(a A, c C) (B, C) {
		return fn(a), c
	}
}

// Send the second component of the input through the argument arrow, and copy the rest unchanged to the output.
func Second[A, B, C any](fn func(a A) B) func(c C, a A) (C, B) {
	return func(c C, a A) (C, B) {
		return c, fn(a)
	}
}

func Identity[A any](a A) A { return a }
