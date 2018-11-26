package search

import "testing"

func TestMatrix(t *testing.T) {
	cases := []struct {
		matrixIn [][]int
		targetIn int
		want     bool
	}{
		{[][]int{{1}, {2}, {3}}, 2, true},
		{[][]int{{1, 2, 3}}, 3, true},
		{[][]int{{1, 2, 3}}, 2, true},
		{[][]int{{1, 2, 3}}, 1, true},
		{[][]int{{1, 2, 3}}, 0, false},
		{[][]int{{}}, 50, false},
		{[][]int{}, 50, false},
		{[][]int{{1, 4}, {2, 5}, {3, 6}}, 3, true},
	}

	for _, c := range cases {
		got := sortedMatrix(c.matrixIn, c.targetIn)

		if got != c.want {
			t.Errorf("Wrong answer!")
		}
	}
}
