package main

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := 0

	lC := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := lC[nums[i]]; !ok {
			leftLC := lC[nums[i]-1]
			rightLC := lC[nums[i]+1]

			curLS := leftLC + rightLC + 1
			lC[nums[i]] = curLS

			if curLS > res {
				res = curLS
			}
			lC[nums[i]-leftLC] = curLS
			lC[nums[i]+rightLC] = curLS
		}
	}

	return res
}
