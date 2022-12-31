package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	rowsSize := len(matrix)
	colsSize := len(matrix[0])

	upRow := 0
	downRow := rowsSize
	midRow := (upRow + downRow) / 2

	for upRow != downRow-1 {
		if target < matrix[midRow][0] {
			downRow = midRow
		} else {
			upRow = midRow
		}

		midRow = (upRow + downRow) / 2
	}

	targetRow := upRow

	leftCol := 0
	rightCol := colsSize
	midCol := (leftCol + rightCol) / 2

	for leftCol != rightCol-1 {
		if target < matrix[targetRow][midCol] {
			rightCol = midCol
		} else {
			leftCol = midCol
		}

		midCol = (leftCol + rightCol) / 2
	}

	return target == matrix[targetRow][leftCol]
}
