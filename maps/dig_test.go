package maps

import "fmt"

func ExampleDig() {
	type c struct {
		d string
	}
	data := map[string]any{"a": []c{
		{d: "e"},
		{d: "f"},
	},
	}
	res1, _ := Dig[c](data, "a", 0)
	fmt.Println(res1.d)

	// Output: e
}
