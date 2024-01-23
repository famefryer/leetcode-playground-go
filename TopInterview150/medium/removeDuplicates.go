package medium

func removeDuplicates(nums []int) int {
	size := len(nums)
	i := 0
	dupCount := 1
	for j := 1; j < size; j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
			dupCount = 1
		} else {
			if dupCount < 2 {
				i++
				dupCount++
				nums[i] = nums[j]
			}
		}
	}

	return i + 1
}

// Better Solution : We used two pointer to loop through nums array
// First pointer (index) used to point at the 2 prior elements of the second pointer (i)
// Second pointer (i) used to point and loop through every element if array and check with 2 prior element that is it duplicate or not
// If no we will assign it to the index position and increase index to assign the next value
// Short Summary: If it duplicated move on, if not assign it to the next space
func removeDuplicates2(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	index := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[index-2] {
			nums[index] = nums[i]
			index++
		}
	}

	return index
}
