package dp

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

// Solution for longest increasing subsequence in O(N^2).
//
func longestIncreasingSubsequence1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var sol = make([]int, len(nums))
	maxAns := 1

	for i := 1; i < len(nums); i++ {
		maxSol := 0

		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				maxSol = max(maxSol, sol[j])
			}
		}

		sol[i] = maxSol + 1
		maxAns = max(maxAns, sol[i])
	}

	return maxAns
}
