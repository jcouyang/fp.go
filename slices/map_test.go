package slices

import (
	"fmt"
)

func ExampleMap() {
	strLen := func (a1 string) int {
			return len(a1)
		}
	fmt.Println(
		Map(strLen)([]string{"cat", "dog"}),
	)
	// Output: [3 3]
}
