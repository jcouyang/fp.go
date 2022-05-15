package maps

import "fmt"

func ExampleFlatMap() {
	animalList := map[string]int{
		"cat":  1,
		"dog":  2,
		"fish": 3,
	}

	flipKV := FlatMap(func(k string, v int) map[int]string {
		return map[int]string{
			v: k,
		}
	})
	fmt.Println(
		flipKV(animalList),
	)
	// Output: map[1:cat 2:dog 3:fish]
}
