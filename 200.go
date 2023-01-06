package main

func dfsOfIslands(grid [][]byte, x, y int) {
	grid[x][y] = '0'
	rows := len(grid)
	cols := len(grid[0])

	if x != 0 && grid[x-1][y] == '1' {
		dfsOfIslands(grid, x-1, y)
	}

	if y != 0 && grid[x][y-1] == '1' {
		dfsOfIslands(grid, x, y-1)
	}

	if x != rows-1 && grid[x+1][y] == '1' {
		dfsOfIslands(grid, x+1, y)
	}

	if y != cols-1 && grid[x][y+1] == '1' {
		dfsOfIslands(grid, x, y+1)
	}
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	ans := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				ans++
				dfsOfIslands(grid, i, j)
			}
		}
	}

	return ans
}
