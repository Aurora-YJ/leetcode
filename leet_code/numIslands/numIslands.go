package main

import (
	"fmt"
)

func main() {
	// Example grid
	grid := [][]byte{
		{1, 1, 0, 0},
		{1, 0, 1, 0},
		{0, 1, 1, },
		{0, 0, 1, 0},
	}

	// Call the function to count the connected components
	result := numIslands(grid)
	
	fmt.Println("Number of connected components:", result)
}
func numIslands(grid [][]byte) int {
	count := 0
for i := 0; i < len(grid); i++ {
	for j := 0; j < len(grid[i]); j++ {
		if grid[i][j] == '1'  {
			count++
			dfs(grid, i , j)
		}
	}
}
return count
}


func dfs(g [][]byte , i, j int) {
if i < 0 || j < 0 || i >= len(g) ||  j >= len(g[0]) || g[i][j] != '1' {
	return
}
g[i][j] = '2'
dfs(g , i +1 , j)
dfs(g , i -1 , j)
dfs(g , i ,j+1)
dfs(g , i ,j-1)
}