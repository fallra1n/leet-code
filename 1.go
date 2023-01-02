package main

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	retVals := make([]int, 2)

	for id, num := range nums {
		dict[num] = id
	}

	for id, num := range nums {
		if i, ok := dict[target-num]; ok && i != id {
			retVals[0] = id
			retVals[1] = i
			break
		}
	}

	return retVals
}
