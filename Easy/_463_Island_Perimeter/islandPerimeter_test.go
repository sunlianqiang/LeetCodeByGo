package _463_Island_Perimeter

import "testing"

func TestIslandPerimeter(t *testing.T) {
	input := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0}}

	want := 16

	ret := IslandPerimeter(input)

	if ret == want {
		t.Logf("pass")
	} else {
		t.Errorf("fail, want %+v, get %+v", want, ret)
	}
}
