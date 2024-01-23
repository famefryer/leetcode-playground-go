package easy

// My code
func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		curNum := nums2[i]
		for j := 0; j < m+n; j++ {
			if j >= m+i {
				nums1[j] = curNum
				break
			}

			if curNum < nums1[j] {
				temp := curNum
				curNum = nums1[j]
				nums1[j] = temp
			}
		}
	}
}

// Better solution, Since both array already sorted, so we can compare from the back of both array
// and continuously place it at the end of nums1
func merge2(nums1 []int, m int, nums2 []int, n int) {
	tail := m + n - 1
	i := m - 1
	j := n - 1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[tail] = nums1[i]
			tail--
			i--
		} else {
			nums1[tail] = nums2[j]
			tail--
			j--
		}
	}
}
