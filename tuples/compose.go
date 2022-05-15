package tuples

func Compose[A, B, C any](fnB func(a B) (C, error), fnA func(a A) (B, error)) func(a A) (C, error) {
	return func(a A) (C, error) {
		return FlatMap(fnB)(fnA(a))
	}
}

func Pipe[A, B, C any](fnA func(a A) (B, error), fnB func(a B) (C, error)) func(a A) (C, error) {
	return Compose(fnB, fnA)
}

func Pipe3[A, B, C, D any](fnA func(a A) (B, error), fnB func(a B) (C, error), fnC func(c C) (D, error)) func(a A) (D, error) {
	return Compose(fnC, Compose(fnB, fnA))
}

func Pipe4[A, B, C, D, E any](fnA func(a A) (B, error), fnB func(a B) (C, error), fnC func(c C) (D, error), fnD func(d D) (E, error)) func(a A) (E, error) {
	return Compose(fnD, Compose(fnC, Compose(fnB, fnA)))
}

func Pipe5[A, B, C, D, E, F any](fnA func(a A) (B, error), fnB func(a B) (C, error), fnC func(c C) (D, error), fnD func(d D) (E, error), fnE func(e E) (F, error)) func(a A) (F, error) {
	return Compose(fnE, Compose(fnD, Compose(fnC, Compose(fnB, fnA))))
}

func Pipe6[A, B, C, D, E, F, G any](fnA func(a A) (B, error), fnB func(a B) (C, error), fnC func(c C) (D, error), fnD func(d D) (E, error), fnE func(e E) (F, error), fnF func(e F) (G, error)) func(a A) (G, error) {
	return Compose(fnF, Compose(fnE, Compose(fnD, Compose(fnC, Compose(fnB, fnA)))))
}

func Pipe7[A, B, C, D, E, F, G, H any](fnA func(a A) (B, error), fnB func(a B) (C, error), fnC func(c C) (D, error), fnD func(d D) (E, error), fnE func(e E) (F, error), fnF func(e F) (G, error), fnG func(e G) (H, error)) func(a A) (H, error) {
	return Compose(fnG, Compose(fnF, Compose(fnE, Compose(fnD, Compose(fnC, Compose(fnB, fnA))))))
}
