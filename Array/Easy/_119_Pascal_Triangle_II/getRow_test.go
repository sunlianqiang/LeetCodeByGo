package _119_Pascal_Triangle_II

import "testing"

func TestGetRow(t *testing.T) {
	var tests = []struct{
		row int
		ret []int
	}{
		{2, []int{1,2,1}},
	}

	for _, v := range tests {
		GetRow(v.row)

	}
}