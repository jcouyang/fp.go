package fp

func Compose[A any, B any, C any](fn1 func(b B) C, fn2 func(a A) B) func(a A) C {
	return func(a A) C {
		return fn1(fn2(a))
	}
}

func Pipe[A any, B any, C any](fn1 func(a A) B, fn2 func(b B) C) func(a A) C {
	return func(a A) C {
		return fn2(fn1(a))
	}
}
