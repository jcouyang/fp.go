package slices

import (
	"fmt"
	"strings"
)

func ExampleFlatMap() {
	splitWords := func(w string) []string {
		return strings.Split(w, " ")
	}
	fmt.Println(strings.Join(
		FlatMap(splitWords)([]string{"hello world", "this is meowesome"}),
		","))

	// Output: hello,world,this,is,meowesome
}
