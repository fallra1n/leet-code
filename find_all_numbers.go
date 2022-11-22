package main

// import "fmt"

func findDisappearedNumbers(nums []int) []int {
	if len(nums) == 0 {
		return []int{1}
	}

	max := nums[0]
	flags := make(map[int]bool)

	for _, num := range nums {
		flags[num] = true
		if num > max {
			max = num
		}
	}

	res := make([]int, 0, 0)

	for i := 1; i <= max; i++ {
		if flags[i] == false {
			res = append(res, i)
		}
	}

	if len(flags) != len(nums) {
		for i := max + 1; i <= len(nums); i++{
			res = append(res, i)
		} 
	}

	return res
}

// func main() {
// 	a := []int{1, 1}
// 	fmt.Println(findDisappearedNumbers((a)))
// }
