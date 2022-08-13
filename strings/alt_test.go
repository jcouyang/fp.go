package strings

import (
	"fmt"
)

func ExampleAlt() {
	fmt.Println(
		Alt("")("default"),
		Alt("not empty")("default"),
	)
	// Output: default not empty
}
