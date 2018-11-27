package dp

import "testing"

func TestLongestIncreasingSubsequence1(t *testing.T) {
	cases := []struct {
		numsIn []int
		want   int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{4, 5, 1, 2, 0}, 2},
	}

	for _, c := range cases {
		got := longestIncreasingSubsequence1(c.numsIn)

		if got != c.want {
			t.Error("Wrong answer!")
		}
	}
}
