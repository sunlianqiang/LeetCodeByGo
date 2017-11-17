package _121_Best_Time2Buy_and_Sell_Stock

import "testing"

func TestMaxProfit(t *testing.T) {
	var tests = []struct{
		input []int
		output int
	} {
		{[]int{}, 0},
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, v := range tests {
		ret := MaxProfit(v.input)

		if ret == v.output {
			t.Logf("pass")
		} else {
			t.Errorf("fail, want %+v, get %+v\n", v.output, ret)
		}
	}
}
