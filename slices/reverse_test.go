package slices

import "fmt"

func ExampleReverse() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(
		Reverse(a), a,
	)
	// Output: [5 4 3 2 1] [1 2 3 4 5]
}

func ExampleReverseInline() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(
		ReverseInline(a), a,
	)
	// Output: [5 4 3 2 1] [5 4 3 2 1]
}
