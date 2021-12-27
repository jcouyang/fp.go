package maps

import "fmt"

func ExampleMap() {
	animalList := map[string]int{
		"cat": 1,
		"dog": 2,
		"fish": 3,
	}
	eachPlus1 := Map(func(_ string, v int) int {return v+1})
	fmt.Println(
		eachPlus1(animalList),
	)
	// Output: map[cat:2 dog:3 fish:4]
}
