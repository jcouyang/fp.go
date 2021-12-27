package tuples

import (
	"fmt"
)

func ExampleFlatMap() {
	proc1 := func(input string) (string, error) {
		if len(input) > 0{
			return input, nil
		}
		return "", fmt.Errorf("error from proc1")
	}
	proc2 := func(input string) (int, error) {
		if len(input) > 5 {
			return len(input), nil
		} else {
			return 0, fmt.Errorf("error from proc2")
		}
	}

	fmt.Println(
		FlatMap(proc2)(proc1("bad")),
	)

	fmt.Println(
		FlatMap(proc2)(proc1("")),
	)

	fmt.Println(
		FlatMap(proc2)(proc1("good string")),
	)
	// Output: 0 error from proc2
	// 0 error from proc1
	// 11 <nil>
}
