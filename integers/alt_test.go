package integers

import (
	"fmt"
)

func ExampleAlt() {
	fmt.Println(
		Alt(1)(2),
		Alt(0)(2),
	)
	// Output: 1 2
}
