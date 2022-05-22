package main

import "fmt"

func main() {
	for i := 0; i < 7; i++ {
		fmt.Printf("第%d题: \n",i+1)
	}
}
func islandPerimeter(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return dfs(i,j,grid)
			}
		}
	}
	return 0
}
func dfs(x, y int, grid [][]int) int {
	if x < 0 || x > len(grid)-1 || y < 0 || y > len(grid[0])-1 {
		return 1
	}
	if grid[x][y] == 0 {
		return 1
	}
	if grid[x][y] == 2 {
		return 0
	}
	grid[x][y] = 2
	return dfs(x+1,y,grid)+dfs(x-1,y,grid)+dfs(x,y+1,grid)+dfs(x,y-1,grid)
}