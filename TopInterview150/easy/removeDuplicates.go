package easy

// Solution, We used two pointers and loop through sorted array.
// First pointer (i) start with 0 and second pointer (j) start with 1. And we used j for our loop condition
// And in every loop we check that if value that i point is not equals to value that j point then we increase i and assign that value to nums[i]
func removeDuplicates(nums []int) int {
	size := len(nums)
	i := 0
	for j := 1; j < size; j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}
