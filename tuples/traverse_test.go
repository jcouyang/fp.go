package tuples

import (
	"encoding/base64"
	"fmt"
)

func ExampleTraverse() {
	decoded, _ := Traverse(base64.StdEncoding.DecodeString)([]string{"aGVsbG8K", "d29ybGQK"})
	fmt.Printf("%s\n", decoded)

	_, err := Traverse(base64.StdEncoding.DecodeString)([]string{"aGVsbG8K", "123"})
	fmt.Println(err)

	// Output: [hello
	//  world
	// ]
	// illegal base64 data at input byte 0
}
