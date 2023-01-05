package main

// we use sorting by counting, because number of different symbols is small
// => time complexity = O(nm)

func sortingByCounting(str string) string {
	var res []byte
	arrOfSymbols := make([]int, 26)

	for i := 0; i < len(str); i++ {
		arrOfSymbols[int(str[i]-'a')] += 1
	}

	for id, symbol := range arrOfSymbols {
		for i := 0; i < symbol; i++ {
			res = append(res, byte(id+'a'))
		}
	}

	return string(res)
}

func groupAnagrams(strs []string) [][]string {
	var ans [][]string
	anagramsGroup := make(map[string][]string)

	for _, str := range strs {
		sStr := sortingByCounting(str)

		anagramsGroup[sStr] = append(anagramsGroup[sStr], str)
	}

	for _, str := range anagramsGroup {
		ans = append(ans, str)
	}

	return ans
}
