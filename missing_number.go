package main

import "fmt"

func missingNumber(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	
	max := nums[0]
	flags := make(map[int]bool)

	for _, num := range nums {
		flags[num] = true
		if num > max {
			max = num
		}
	}

	for i := 0; i <= max; i++ {
		if flags[i] == false {
			return i
		}
	}

	return max + 1
}

func main() {
	a := []int{1, 2, 3, 4};
	fmt.Println(missingNumber((a)))	
}
