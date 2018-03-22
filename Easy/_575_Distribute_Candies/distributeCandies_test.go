package _575_Distribute_Candies

import "testing"

func TestDistributeCandies(t *testing.T) {
	candies1 := []int{1,1,2,2,3,3}
	kinds1 := DistributeCandies(candies1)

	want1 := 3
	if kinds1 != want1 {
		t.Errorf("fail, want %+v, get %+v\n", want1, kinds1)
	}

	candies2 := []int{1,1,2,3}
	kinds2 := DistributeCandies(candies2)

	want2 := 2
	if kinds2 != want2 {
		t.Errorf("fail, want %+v, get %+v\n", want2, kinds2)
	}
}