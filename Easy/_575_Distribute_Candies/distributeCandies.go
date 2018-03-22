package _575_Distribute_Candies

func DistributeCandies(candies []int) int {
	var kinds int
	length := len(candies)
	kinds = length / 2

	var candiesMap map[int]int
	candiesMap = make(map[int]int)
	for _, v := range candies {
		_, ok := candiesMap[v]
		if !ok {
			candiesMap[v] = 1
		} else {
			candiesMap[v]++
		}
	}

	mapLen := len(candiesMap)
	if kinds > mapLen {
		kinds = mapLen
	}

	return kinds
}