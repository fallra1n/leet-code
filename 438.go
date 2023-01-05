package main

import "fmt"

func TernaryOperator(expr bool, val *int) {
	if expr == true {
		*val--
	} else {
		*val++
	}
}

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}

	var ans []int
	compare := make([]int, 26)
	sum := 0

	for i := 0; i < len(p); i++ {
		oldPi := compare[p[i]-'a']
		oldSi := compare[s[i]-'a']
		compare[p[i]-'a']++
		compare[s[i]-'a']--

		if p[i] != s[i] {
			TernaryOperator(oldPi < 0, &sum)
			TernaryOperator(oldSi > 0, &sum)
		}
	}

	for i := 0; i < len(s)-len(p); i++ {
		if sum == 0 {
			ans = append(ans, i)
		}

		oldCh := compare[s[i]-'a']
		newCh := compare[s[i+len(p)]-'a']
		compare[s[i]-'a']++
		compare[s[i+len(p)]-'a']--

		if s[i] != s[i+len(p)] {
			TernaryOperator(oldCh < 0, &sum)
			TernaryOperator(newCh > 0, &sum)
		}
	}

	if sum == 0 {
		ans = append(ans, len(s)-len(p))
	}

	return ans
}

func main() {
	fmt.Println(findAnagrams("baa", "aa"))
}
