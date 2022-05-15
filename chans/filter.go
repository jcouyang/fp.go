package chans

func Filter[A any](fn func(a A) bool) func(a <-chan A) chan A {
	return func(a <-chan A) (b chan A) {
		b = make(chan A, len(a))
		go func() {
			for i := range a {
				if fn(i) {
					b <- i
				}
			}
		}()
		return
	}
}
