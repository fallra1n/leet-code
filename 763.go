package main

import "math"

func partitionLabels(s string) []int {
	size := len(s)
	var ans []int

	lastPos := make(map[byte]int)
	for i := 0; i < size; i++ {
		lastPos[s[i]] = i
	}

	anchor := 0
	last := 0

	for i := 0; i < size; i++ {
		last = int(math.Max(float64(last), float64(lastPos[s[i]])))
		if i == last {
			ans = append(ans, i-anchor)
			anchor = i + 1
		}
	}

	return ans
}
