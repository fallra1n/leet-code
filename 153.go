package main

func findMin(nums []int) int {
	if (len(nums) == 1) || (nums[0] < nums[len(nums)-1]) {
		return nums[0]
	}

	left := 0
	right := len(nums) - 1
	mid := (left + right) / 2

	for left != right-1 {
		if nums[left] < nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}

		mid = (left + right) / 2
	}

	return nums[right]
}
