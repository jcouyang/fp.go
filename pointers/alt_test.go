package pointers

import (
	"fmt"
)

func ExampleAlt() {
	val1 := "val1"
	val2 := "val2"
	fmt.Println(
		*Alt[string](nil)(&val1),
		*Alt(&val2)(&val1),
	)
	// Output: val1 val2
}
