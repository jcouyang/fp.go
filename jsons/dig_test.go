package jsons

import "fmt"

func ExampleDig() {
	data := `{"a": [{"b": {"c": 1}}]}`
	res1, _ := Dig[map[string]any]([]byte(data), "a", 0, "b")
	res2, _ := res1["c"].(float64)
	fmt.Println(res1, int(res2))
	

	res3,_ := Dig[float64]([]byte(data), "a", 0, "b", "c")
	
	fmt.Println(res3)
	
	_, err1 := Dig[float64]([]byte(data), "a", 0, "not-exits", "c")
	fmt.Println(err1)
	
	// Output: map[c:1] 1
	// 1
	// expect map[string]any got <nil> at path [a 0 not-exits]
}
