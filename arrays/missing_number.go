package array

// Given an array of 1 to n. Find the missing number.
//
func missingNumber(nums []int) int {
	sum := 0

	for _, c := range nums {
		sum += c
	}

	return ((len(nums) + 1) * (len(nums) + 2) / 2) - sum
}
