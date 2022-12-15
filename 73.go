package main

func setZeroes(matrix [][]int) {
	if matrix == nil {
		return
	}

	first_col_is_zero := false
	first_width_is_zero := false

	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			first_col_is_zero = true
			break
		}
	}

	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			first_width_is_zero = true
			break
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if first_col_is_zero {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}

	if first_width_is_zero {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
}

//func main() {
//	a := make([][]int, 1, 1)
//	a[0] = append(a[0], 0)
//	a[0] = append(a[0], 1)
//	setZeroes(a)
//	print(a)
//}
