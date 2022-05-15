package slices

import (
	"fmt"
	"oyanglul.us/fp.go/integers"
)

func ExampleFilter() {
	filterOddInt := Filter(integers.Odd[int])
	filterOddUint := Filter(integers.Odd[uint])
	fmt.Println(
		filterOddInt([]int{-2, -1, 0, 1, 2}),
		filterOddUint([]uint{1, 2, 3, 4, 5}),
	)
	// Output: [-1 1] [1 3 5]
}
