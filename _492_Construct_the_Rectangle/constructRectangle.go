package _492_Construct_the_Rectangle

import "math"

func constructRectangle(area int) []int {
	var ret []int
	sqrtArea := math.Sqrt(float64(area))
	var w int
	w = int(sqrtArea)

	for ; w > 0; w-- {
		if area % w == 0 {
			l := area / w
			ret = []int{l, w}
			break
		}
	}

	return ret
}