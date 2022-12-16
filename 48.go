package main

func rotate(matrix [][]int) {
	n := len(matrix)
	curLen := len(matrix)
	x := 0
	y := 0
	oldX := 0
	oldY := 0
	oldVal := 0

	for i := 0; i < len(matrix)/2; i++ {
		oldX = x
		oldY = x

		for k := 0; k < curLen-1; k++ {
			oldVal = matrix[x][y]
			for j := 0; j < 4; j++ {
				tempVal := matrix[y][n-1-x]
				matrix[y][n-1-x] = oldVal
				oldVal = tempVal
				x, y = y, n-1-x
			}
			x++
		}

		x = oldX + 1
		y = oldY + 1
		curLen -= 2
	}
}

//func main() {
//	a := make([][]int, 3)
//	a[0] = append(a[0], 1, 2, 3)
//	a[1] = append(a[1], 4, 5, 6)
//	a[2] = append(a[2], 7, 8, 9)
//	rotate(a)
//	fmt.Print(a)
//}
