package tuples

import "fmt"

func ExampleMap() {
	proc := func(good string) (string, error) {
		if good == "good" {
			return "a string", nil
		}
		return "", fmt.Errorf("error proc1")
	}
	strLen := func(input string) int {
		return len(input)
	}

	fmt.Println(
		Map(strLen)(proc("good")),
	)

	fmt.Println(
		Map(strLen)(proc("bad")),
	)
	// Output: 8 <nil>
	// 0 error proc1
}
