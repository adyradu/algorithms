package array

import (
	"testing"
)

func TestMissingNumber(t *testing.T) {
	cases := []struct {
		numsIn []int
		want   int
	}{
		{[]int{2, 1, 3}, 4},
		{[]int{3, 1, 4, 5}, 2},
		{[]int{2}, 1},
	}

	for _, c := range cases {
		got := missingNumber(c.numsIn)

		if got != c.want {
			t.Error("Wrong answer!")
		}
	}
}
