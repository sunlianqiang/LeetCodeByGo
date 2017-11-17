package _447_Number_Boomerangs

import "testing"

func TestNumberOfBoomerangs(t *testing.T) {
	var tests = []struct{
		intput [][]int
		output int
	} {
		{
			[][]int{{0,0},{1,0},{2,0}}, 2,
		},
		{
			[][]int{{0,0}, {1,0}, {-1,0}, {0,1}, {0,-1}}, 20,
		},
	}

	for _, test := range tests {
		ret := NumberOfBoomerangs(test.intput)

		if ret == test.output {
			t.Logf("pass")
		} else {
			t.Errorf("fail, want %+v, get %+v", test.output, ret)
		}
	}
}
