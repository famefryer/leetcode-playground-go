package medium

func Rotate(nums []int, k int) {
	size := len(nums)
	copyNums := make([]int, size)
	copy(copyNums, nums)
	for i := 0; i < size; i++ {
		newIndex := 0
		if (i + k) >= size {
			newIndex = (i + k) % size
		} else {
			newIndex = i + k
		}
		nums[newIndex] = copyNums[i]
	}
}

// Better solution, For go version1.21 or later
//func rotate2(nums []int, k int) {
//	if len(nums) <= 1 || k == 0 {
//		return
//	}
//	k = k % len(nums)
//	slices.Reverse(nums)
//	slices.Reverse(nums[0:k])
//	slices.Reverse(nums[k:])
//}
