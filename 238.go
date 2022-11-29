package main

// import "fmt"

func productExceptSelf(nums []int) []int {
    mul := 1
	number_of_zero := 0

	for _, num := range nums {
		if num == 0 {
			number_of_zero++
		} else {
			mul *= num
		}
	}

	ans := make([]int, len(nums))

	for i, num := range nums {
		if num == 0 {
			if number_of_zero > 1 {
				ans[i] = 0
			} else {
				ans[i] = mul
			}
		} else {
			if number_of_zero == 0 {
		    	ans[i] = mul / num
			} else {
				ans[i] = 0
			}
	    }
	}

	return ans
}

// func main() {
// 	data := []int{1, 2, 3, 4}
// 	fmt.Println(productExceptSelf(data))
// }