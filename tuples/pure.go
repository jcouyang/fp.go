package tuples

func Pure[A any](a A) (A, error) {
	return a, nil
}

func Error[A any](e error) (*A, error) {
	return nil, e
}
