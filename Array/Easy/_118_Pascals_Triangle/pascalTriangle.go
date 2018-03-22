package _118_Pascals_Triangle

import "fmt"

func Generate(numRows int) [][]int {
	var triangle [][]int
	if 0 == numRows {
		return triangle
	}
	triangle = append(triangle, []int{1})
	if 1 == numRows {
		return triangle
	}

	i := 1
	for ; i < numRows; i++ {
		var row []int
		row = append(row, 1)

		j := 1
		for ; j < i/2 + 1; j++ {
			tmp := triangle[i-1][j-1] + triangle[i-1][j]
			row = append(row, tmp)
		}

		j = (i-1)/2
		for ; j >=0; j-- {
			row = append(row, row[j])
		}
		fmt.Printf("row:%+v", row)

		triangle = append(triangle, row)
	}

	return triangle
}
