package main

func search(nums []int, target int) int {
	left := 0
	right := len(nums)
	mid := (left + right) / 2

	res := -1

	for left != right-1 {
		if nums[left] < nums[mid] {
			if target >= nums[left] && target < nums[mid] {
				right = mid
			} else {
				left = mid
			}

			mid = (right + left) / 2
			continue
		}

		if nums[mid] <= nums[right-1] {
			if target >= nums[mid] && target <= nums[right-1] {
				left = mid
			} else {
				right = mid
			}

			mid = (right + left) / 2
			continue
		}
	}

	if nums[left] == target {
		res = left
	}

	return res
}
