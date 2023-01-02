package main

func search(nums []int, target int) bool {
	res := false

	left := 0
	right := len(nums) - 1
	mid := (left + right) / 2

	for left != right {
		mid = (left + right) / 2
		if nums[mid] == target {
			res = true
			break
		}

		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
			continue
		}

		// first segment increase
		if nums[left] <= nums[mid] {
			// target is included in first segment
			if target <= nums[mid] && target >= nums[left] {
				right = mid - 1
				continue
			}

			// target is included in second segment
			left = mid + 1
			continue
		}

		// if target is included in second segment && second segment increase
		if target <= nums[right] && target >= nums[mid] {
			left = mid + 1
			continue
		}

		// if target is included in first segment && second segment increase
		right = mid - 1
	}

	return res
}
