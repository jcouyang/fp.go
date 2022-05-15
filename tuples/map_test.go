package tuples

import "fmt"

func proc(good string) (string, error) {
	if good == "good" {
		return "a string", nil
	}
	return "", fmt.Errorf("error proc1")
}
func strLen(input string) int {
	return len(input)
}

func strLen2(i1 string, i2 string) int {
	return len(i1) + len(i2)
}
func ExampleMap() {

	fmt.Println(
		Map(strLen)(proc("good")),
	)

	fmt.Println(
		Map(strLen)(proc("bad")),
	)
	// Output: 8 <nil>
	// 0 error proc1
}

func ExampleMap2() {
	fmt.Println(
		Map2(strLen2)(proc("good"))(proc("bad")),
	)

	fmt.Println(
		Map(strLen)(proc("bad")),
	)
	// Output: 8 <nil>
	// 0 error proc1
}
