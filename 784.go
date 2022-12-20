package main

// A1b2 -> [a1b2, A1B2, a1B2]

func isLetter(ch byte) bool {
	return ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z'
}

func letterCasePermutation(s string) []string {
	countOfLetters := 0
	counter := 1

	for i := 0; i < len(s); i++ {
		if isLetter(s[i]) {
			countOfLetters++
		}
	}

	for i := 0; i < countOfLetters; i++ {
		counter *= 2
	}

	res := make([][]byte, counter)
	for i := 0; i < counter; i++ {
		res[i] = make([]byte, len(s))
	}

	for i := 0; i < counter; i++ {
		count := 0

		for j := 0; j < len(s); j++ {
			if isLetter(s[j]) {
				temp := uint32(i)
				temp <<= 32 - count - 1
				temp >>= 31

				if temp == 1 {
					if s[j] >= 'a' && s[j] <= 'z' {
						res[i][j] = s[j] - 'a' + 'A'
					} else {
						res[i][j] = s[j] + 'a' - 'A'
					}
				} else {
					res[i][j] = s[j]
				}

				count++
			} else {
				res[i][j] = s[j]
			}
		}

	}

	retRes := make([]string, counter)
	for i := 0; i < counter; i++ {
		retRes[i] = string(res[i])
	}

	return retRes
}

//func main() {
//	fmt.Println(letterCasePermutation("a1b2"))
//}
