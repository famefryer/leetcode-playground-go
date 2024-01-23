package easy

func RemoveElement(nums []int, val int) int {
	size := len(nums)

	if size == 1 {
		if nums[0] == val {
			return 0
		}
		return 1
	}

	head := 0
	tail := size - 1
	for (tail - head) > 0 {
		if nums[tail] != val && nums[head] == val {
			temp := nums[tail]
			nums[tail] = val
			nums[head] = temp
			tail--
			head++

			continue
		}

		if nums[tail] == val {
			tail--
		}

		if nums[head] != val {
			head++
		}
	}

	noDup := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			noDup++
		} else {
			break
		}
	}

	return noDup
}

// Better Solution : function will loop through array with 2 index,
// i will be used to get every element of nums and k will be used to assign value back to the array
func removeElement2(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
