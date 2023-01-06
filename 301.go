package main

func isPsp(s string) bool {
	i := 0

	for _, ch := range s {
		if ch == '(' {
			i++
		}

		if ch == ')' {
			i--
		}

		if i < 0 {
			return false
		}
	}

	return i == 0
}

func dfs(s string, ans *[]string, curPos, openFail, closeFail int) {
	if openFail == 0 && closeFail == 0 {
		if isPsp(s) {
			*ans = append(*ans, s)
			return
		}
	}

	for i := curPos; i < len(s); i++ {
		if i != curPos && s[i] == s[i-1] { // drop recurring cases
			continue
		}

		cur := s
		cur = cur[:i] + cur[i+1:]

		if closeFail > 0 && s[i] == ')' {
			dfs(cur, ans, i, openFail, closeFail-1)
		} else if openFail > 0 && s[i] == '(' {
			dfs(cur, ans, i, openFail-1, closeFail)
		}
	}
}

func removeInvalidParentheses(s string) []string {
	var ans []string

	openFail := 0
	closeFail := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			openFail++
		}

		if openFail == 0 && s[i] == ')' {
			closeFail++
		} else if s[i] == ')' {
			openFail--
		}
	}

	dfs(s, &ans, 0, openFail, closeFail)

	return ans
}

//func main() {
//	removeInvalidParentheses("()())()")
//}
