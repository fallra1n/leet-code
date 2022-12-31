package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums)
	mid := (left + right) / 2

	for left < right {
		if target == nums[mid] {
			return mid
		}

		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}

		mid = (left + right) / 2
	}

	return -1
}
