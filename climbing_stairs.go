package main

// import "fmt"

var data map[int]int

func solution(n int) int {
    if n == 1 {
		data[1] = 1
		return 1
	}

	if n == 2 {
		data[2] = 2
		return 2;
	}

	if _, ok := data[n - 2]; !ok {
		solution(n - 2)
	}

	if _, ok := data[n - 1]; !ok {
		solution(n - 1)
	}

	data[n] = data[n - 2] + data[n - 1]
	return data[n]
}

func climbStairs(n int) int {
	data = make(map[int]int)
	return solution(n)
}


// func main() {
// 	a := 3
// 	fmt.Println(climbStairs((a)))	
// }