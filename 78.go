package main

func subsets(nums []int) [][]int {
	s := len(nums)
	counter := 1

	for i := 0; i < s; i++ {
		counter *= 2
	}

	res := make([][]int, counter)

	for i := 0; i < counter; i++ {
		for j := 0; j < s; j++ {
			temp := uint32(i)
			temp <<= 32 - j - 1
			temp >>= 31

			if temp == 1 {
				res[i] = append(res[i], nums[j])
			}
		}
	}

	return res
}
