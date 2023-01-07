package main

import "math"

func maxArea(height []int) int {
	var ans int
	var l int
	var r int

	ans = 0
	l = 0
	r = len(height) + 1

	for l != r {
		if int(math.Min(float64(height[l]), float64(height[r])))*(r-l) > ans {
			ans = int(math.Min(float64(height[l]), float64(height[r]))) * (r - l)
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return ans
}

func main() {
	height := []int{1, 2, 1}
	maxArea(height)
}
