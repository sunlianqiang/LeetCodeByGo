package _447_Number_Boomerangs

import "fmt"

func NumberOfBoomerangs(points [][]int) int {
	var num int
	len1 := len(points)
	for i := 0; i < len1; i++ {
		var distanceMap map[int]int
		distanceMap = make(map[int]int)
		for j := 0; j < len1; j++ {
			if i == j {
				continue
			}
			a := points[i][0] - points[j][0]
			b := points[i][1] - points[j][1]
			distance := a * a + b * b
			distanceMap[distance]++
		}

		fmt.Printf("i:%+v, map:%+v\n", i, distanceMap)
		for _, v := range distanceMap {
			num += v * (v - 1)
		}
	}

	return num
}