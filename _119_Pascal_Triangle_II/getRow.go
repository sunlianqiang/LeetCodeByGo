package _119_Pascal_Triangle_II

import "fmt"

func GetRow(rowIndex int) []int {
	row := []int{1}

	if 0 == rowIndex {
		return row
	}
	row = append(row, 1)
	if 1 == rowIndex {
		return row
	}

	var newRow []int
	for i := 2; i <= rowIndex; i++ {
		newRow = []int{1}

		for j := 1; j < i/2 + 1; j++ {
			tmp := row[j-1] + row[j]
			newRow = append(newRow, tmp)
		}
		fmt.Printf("1newRow:%+v\n", newRow)
		for k := (i-1)/2; k >= 0; k-- {
			newRow = append(newRow, newRow[k])
		}
		fmt.Printf("2newRow:%+v\n", newRow)
		row = newRow
	}

	return row
}
