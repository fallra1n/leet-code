package main

// import "fmt"

func singleNumber(nums []int) int {
	if (len(nums) == 0) {
		return 0
	}

    ans := nums[0]
    for i := 1; i < len(nums); i++ {
        ans = ans ^ nums[i]
    }

    return ans
}

// func main() {
// 	a := []int{1, 2, 1};
// 	fmt.Println(singleNumber((a)))	
// }