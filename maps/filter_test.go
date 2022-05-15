package maps

import "fmt"

func ExampleFilter() {
	animalList := map[string]int{
		"cat":  1,
		"dog":  2,
		"fish": 3,
	}
	moreThan2 := Filter(func(k string, v int) bool { return v > 2 })
	fmt.Println(
		moreThan2(animalList),
	)
	// Output: map[fish:3]
}
