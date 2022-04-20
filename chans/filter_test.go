package chans

import (
	"fmt"
	"time"
	"oyanglul.us/fp.go"
)

func ExampleFilter() {

	// Sized channel
	tokenChan := make(chan string, 3)
	catChan := Filter(fp.Eq("cat"))(tokenChan)

	tokenChan <- "cat"
	tokenChan <- "catfish"
	tokenChan <- "caterpillar"

	go func() {
		for i := range catChan {
			fmt.Println(i)
		}
	}()

	time.Sleep(1 * time.Millisecond)

	// Output: cat
}
