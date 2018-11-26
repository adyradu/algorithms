package search

// Given a matrix n x m where all rows and columns are sorted
// in ascending order, return true if target is in the matrix
// or false otherwise.
//
func sortedMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	for r, c := 0, len(matrix[0])-1; r < len(matrix) && c >= 0; {
		if target == matrix[r][c] {
			return true
		}

		if target > matrix[r][c] {
			r++
		} else {
			c--
		}
	}

	return false
}
