package chans

func Map[A any, B any](fn func (a A) B) func (a chan A) chan B {
	return func (a chan A) (b chan B) {
		b = make(chan B)
		for i := range(a) {
			b <- fn(i)
		}
		return
	}

}

func MapAsync[A any, B any](fn func (a A) B) func (a chan A) chan B {
	return func (a chan A) (b chan B) {
		b = make(chan B)
		for i := range(a) {
			go func (i A) {
				b <- fn(i)
			}(i)
		}
		return
	}
}
