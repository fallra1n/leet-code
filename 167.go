package main

func twoSum(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return nil
	}

	res := make([]int, 2)
	left := 0
	right := len(numbers) - 1

	for {
		if numbers[left]+numbers[right] > target {
			right--
			continue
		}

		if numbers[left]+numbers[right] < target {
			left++
			continue
		}

		res[0] = left + 1
		res[1] = right + 1
		break
	}

	return res
}
