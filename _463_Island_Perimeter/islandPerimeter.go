package _463_Island_Perimeter

func getPerimeter(grid [][]int, len1, len2, i, j int) (perimeter int) {
	//up
	if 0 == i || 0 == grid[i-1][j] {
		perimeter++
	}
	//left
	if 0 == j || 0 == grid[i][j-1] {
		perimeter++
	}

	//right
	if (len1-1) == i || 0 == grid[i+1][j] {
		perimeter++
	}
	//down
	if (len2-1) == j || 0 == grid[i][j+1] {
		perimeter++
	}

	return perimeter
}
func IslandPerimeter(grid [][]int) int {
	 len1 := len(grid)
	 len2 := len(grid[0])

	 var perimiter int
	 for i := 0; i < len1; i++ {
		 for j := 0; j < len2; j++ {
			 if 0 == grid[i][j] {
				 continue
			 }
			 per := getPerimeter(grid, len1, len2, i, j)

			 perimiter += per
		 }
	 }

	return perimiter
}