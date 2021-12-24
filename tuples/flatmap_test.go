package tuples

import (
	"fmt"
	"testing"
)

func proc1(good bool) (string, error) {
	if good {
		return "a string", nil
	}
	return "", fmt.Errorf("error proc1")
}
func proc2(input string) (int, error) {
	return len(input), nil
}
func TestFlatMap(t *testing.T) {
	got, _ := FlatMap(proc2)(proc1(true))
	if got !=8  {
		t.Errorf("want 8 got %d", got)
	}

	_, err := FlatMap(proc2)(proc1(false))
	if err == nil  {
		t.Errorf("want error got nil")
	}
}

