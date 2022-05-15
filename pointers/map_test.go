package pointers

import "fmt"

func ExampleMap() {
	type Vector struct {
		x int
		y int
	}

	tryPlusYBy1 := Map(func(a *Vector) *Vector {
		a.y = a.y + 1
		return a
	})

	tryCopyPlusYBy1 := Map(func(a *Vector) *Vector {
		b := *a
		b.y = b.y + 1
		return &b
	})
	aVector := new(Vector)
	bVector := tryCopyPlusYBy1(aVector)
	fmt.Println(aVector, bVector)
	fmt.Println(tryPlusYBy1(aVector), aVector, bVector)
	fmt.Println(tryPlusYBy1(nil))
	// Output: &{0 0} &{0 1}
	// &{0 1} &{0 1} &{0 1}
	// <nil>
}
