package slices

import (
	"fmt"
)

func ExampleAlt() {
	fmt.Println(
		Alt([]int{})([]int{1}),
		Alt([]int{2})([]int{}),
	)
	// Output: [1] [2]
}
