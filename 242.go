package main

func isAnagram(s string, t string) bool {
	ans := true
	counting := make([]int, 26)

	for _, symbol := range s {
		counting[symbol-'a'] += 1
	}

	for _, symbol := range t {
		counting[symbol-'a'] -= 1
	}

	for _, num := range counting {
		if num != 0 {
			ans = false
			break
		}
	}

	return ans
}
