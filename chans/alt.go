package chans

func Alt[A any](a chan A) func(b chan A) chan A {
	return func(b chan A) chan A {
		c := make(chan A, len(a))
		go func(a chan A, b chan A) {
			select {
			case av := <-a:
				c <- av
			case bv := <-b:
				c <- bv
			}
		}(a, b)
		return c
	}
}
