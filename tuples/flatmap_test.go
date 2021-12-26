package tuples

import (
	"fmt"
)

func ExampleFlatMap() {
	proc1 := func(good bool) (string, error) {
		if good {
			return "a string", nil
		}
		return "", fmt.Errorf("error proc1")
	}
	proc2 := func(input string) (int, error) {
		return len(input), nil
	}

	fmt.Println(
		FlatMap(proc2)(proc1(true)),
	)

	fmt.Println(
		FlatMap(proc2)(proc1(false)),
	)
	// Output: 8 <nil>
	// 0 error proc1
}
