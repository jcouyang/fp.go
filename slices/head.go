package slices

import "fmt"

// Get the pointer of the first element of a slice, error when empty.
func Head[A any](a []A) (*A, error) {
	if len(a) > 0 {
		return &a[0], nil
	}
	return nil, fmt.Errorf("slice with len %d has no head", len(a))
}
