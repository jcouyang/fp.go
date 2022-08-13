package arrow

import (
	"errors"
	"fmt"

	"oyanglul.us/fp.go"
)

func handleErrorIfAny(err error) error {
	if err != nil {
		// handle exception
	}
	return nil
}

func ExampleFirst() {
	doFirstIgnoreSecond := First[string, string, error](fp.Identity[string])
	fmt.Println(
		doFirstIgnoreSecond("first", errors.New("second")),
	)
	// Output: first second
}

func ExampleSecond() {
	doSecondIgnoreFirst := Second[error, error, string](handleErrorIfAny)
	fmt.Println(
		doSecondIgnoreFirst("first", errors.New("second")),
	)
	// Output: first <nil>
}

func ExampleSplit() {
	processBothFirstAndSecond := Split(fp.Identity[string], handleErrorIfAny)
	fmt.Println(
		processBothFirstAndSecond("first", errors.New("second")),
	)
	// Output: first <nil>
}

func ExampleFanout() {
	applyBothFuncToSameInput := Fanout(fp.Identity[string], func (a string) int {return len(a)})
	fmt.Println(applyBothFuncToSameInput("how long is this string"))
	// Output: how long is this string 23
}
