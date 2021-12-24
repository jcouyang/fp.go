package slices

import "fmt"

func ExampleConcat() {
	fmt.Println(
		Concat([]string{"cat", "is"}, []string{"cute", "!"}),
	)
	// Output: [cat is cute !]
}
