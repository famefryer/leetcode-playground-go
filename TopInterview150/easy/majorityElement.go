package easy

func majorityElement(nums []int) int {
	size := len(nums)
	if size == 1 {
		return nums[0]
	}
	for i := 0; i < size; i++ {
		count := 1
		for j := i + 1; j < size; j++ {
			if nums[i] == nums[j] {
				count++
			}
			if count > (size / 2) {
				return nums[i]
			}
		}
	}

	return -1
}

// Better Solution : Use Moore's algorithm
// The algorithm works on the basis of the assumption that the majority element occurs more than n/2 times in the array.
// This assumption guarantees that even if the count is reset to 0 by other elements, the majority element will eventually regain the lead.
func majorityElement2(nums []int) int {
	count := 1
	majority := nums[0]
	for i := 1; i < len(nums); i++ {

		if majority == nums[i] {
			count++
		} else {
			count--
		}
		if count == 0 {
			majority = nums[i]
			count = 1
		}
	}
	return majority
}
