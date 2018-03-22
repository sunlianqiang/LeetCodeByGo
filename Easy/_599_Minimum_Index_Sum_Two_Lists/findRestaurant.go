package _599_Minimum_Index_Sum_Two_Lists


func findRestaurant(list1 []string, list2 []string) []string {
	var restaurants []string

	var restaurantMap1 map[string]int
	restaurantMap1 = make(map[string]int)

	len1 := len(list1)
	for i := 0; i < len1; i++ {
		restaurantMap1[list1[i]] = i
	}

	minIndex := 3000

	len2 := len(list2)
	for i := 0; i < len2; i++ {
		tmp, ok := restaurantMap1[list2[i]]
		if !ok {
			continue
		} else {
			if minIndex > tmp + i {
				minIndex = tmp + i
				restaurants = []string{list2[i]}
			} else if minIndex == tmp + i {
				restaurants = append(restaurants, list2[i])
			}

		}
	}


	return restaurants
}