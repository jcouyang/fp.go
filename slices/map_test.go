package slices

import (
	"fmt"
	"testing"
)

func FuzzMap(f *testing.F) {
	f.Add("cat", "dog")
	f.Add("catac", "dogod")
	f.Fuzz(func(t *testing.T, a string, b string) {
		got := Map(func (a1 string) int {
			return len(a1)
		})([]string{a, b})

		if fmt.Sprint(got) != fmt.Sprint([]int{len(a), len(b)}) {
			t.Errorf("failed map slices, input %s, %s, got %v", a, b, got)
		}
	})	
}
