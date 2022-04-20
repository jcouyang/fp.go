package chans

func Traverse[A any, B any](fn func(a A) chan B) func (a []A) chan []B {
	return func(a []A) chan []B{
		b := make([]B, len(a))
		res := make(chan []B)
		for i,j := range(a) {
			b[i] = <-fn(j)
		}
		res <- b
		return res
	}
}
