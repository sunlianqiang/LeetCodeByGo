package _566_Reshape_Matrix

import "testing"

func TestMatrixReshape(t *testing.T) {
	input := [][]int{
		{1, 2},
		{3, 4},
	}
	ret := MatrixReshape(input, 1, 4)

	t.Logf("%+v", ret)
}
