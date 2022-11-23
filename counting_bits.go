package main

// import "fmt"

func countBits(n int) []int {
	ans := make([]int, n + 1, n + 1)

	ans[0] = 0
	for i := 1; i < n + 1; i++ {
		ans[i] = ans[i / 2]
		ans[i] += (i % 2)
	}

	return ans
}


// func main() {
// 	fmt.Println(countBits(10))
// }
