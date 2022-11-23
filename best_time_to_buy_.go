package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	res := 0
	min := prices[0]

	for i := 1; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
			continue
		}

		if res < prices[i]-min {
			res = prices[i] - min
		}
	}

	return res
}

func main() {
	a := []int{7, 1, 5, 0, 6, 4}
	fmt.Println(maxProfit((a)))
}
