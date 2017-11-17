package _598_Range_Addition_II

func maxCount(m int, n int, ops [][]int) int {
	var row, col int
	row = m
	col = n

	for _, op := range ops {
		if op[0] < row {
			row = op[0]
		}
		if op[1] < col {
			col = op[1]
		}
	}

	return row * col
}
