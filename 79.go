package main

type coord struct {
	x int32
	y int32
}

func searchSymbol(board [][]byte, target byte) []coord {
	res := make([]coord, 0, 0)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == target {
				res = append(res, coord{x: int32(i), y: int32(j)})
			}
		}
	}

	return res
}

func dfsOfWord(board [][]byte, word string, curChar int, flags [][]bool, x, y int) bool {
	if curChar+1 == len(word) {
		return true
	}

	flags[x][y] = true

	if y != 0 {
		if (board[x][y-1] == word[curChar+1]) && (flags[x][y-1] == false) {
			if dfsOfWord(board, word, curChar+1, flags, x, y-1) == true {
				return true
			}
		}
	}

	if x != 0 {
		if (board[x-1][y] == word[curChar+1]) && (flags[x-1][y] == false) {
			if dfsOfWord(board, word, curChar+1, flags, x-1, y) == true {
				return true
			}
		}
	}

	if y != len(board[0])-1 {
		if (board[x][y+1] == word[curChar+1]) && (flags[x][y+1] == false) {
			if dfsOfWord(board, word, curChar+1, flags, x, y+1) == true {
				return true
			}
		}
	}

	if x != len(board)-1 {
		if (board[x+1][y] == word[curChar+1]) && (flags[x+1][y] == false) {
			if dfsOfWord(board, word, curChar+1, flags, x+1, y) == true {
				return true
			}
		}
	}

	flags[x][y] = false
	return false
}

func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return len(word) == 0
	}

	if len(word) == 0 {
		return true
	}

	firstSymbol := searchSymbol(board, word[0])

	for i := 0; i < len(firstSymbol); i++ {
		flags := make([][]bool, len(board))
		for j := 0; j < len(board); j++ {
			flags[j] = make([]bool, len(board[0]))
		}

		if dfsOfWord(board, word, 0, flags, int(firstSymbol[i].x), int(firstSymbol[i].y)) == true {
			return true
		}
	}

	return false
}

//func main() {
//	flags := make([][]byte, 3)
//	for j := 0; j < len(flags); j++ {
//		flags[j] = make([]byte, 4)
//	}
//
//	flags[0][0] = 'A'
//	flags[0][1] = 'B'
//	flags[0][2] = 'C'
//	flags[0][3] = 'E'
//	flags[1][0] = 'S'
//	flags[1][1] = 'F'
//	flags[1][2] = 'E'
//	flags[1][3] = 'S'
//	flags[2][0] = 'A'
//	flags[2][1] = 'D'
//	flags[2][2] = 'E'
//	flags[2][3] = 'E'
//	fmt.Println(exist(flags, string("ABCESEEEFS")))
//}
