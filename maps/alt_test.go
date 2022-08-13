package maps

import (
	"fmt"
)

func ExampleAlt() {
	fmt.Println(
		Alt(map[string]string{"a": "b"})(map[string]string{"c": "d"}),
		Alt(map[string]string{})(map[string]string{"c": "d"}),
	)
	// Output: map[a:b] map[c:d]
}
