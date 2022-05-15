package slices

import "fmt"

func ExampleFoldl() {
	sum := Foldl(func(b int, a int) int {
		return b + a
	}, 0)

	// (((1 + 2) + 3) + 4) + 5
	fmt.Println(sum([]int{1, 2, 3, 4, 5}))
	// Output: 15
}

func ExampleFoldr() {
	foldSub := Foldr(func(a int, b int) int {
		return a - b
	}, 0)

	// 1 - (2 - (3 - (4 - (5 - 0))))
	fmt.Println(foldSub([]int{1, 2, 3, 4, 5}))
	// Output: 3
}
