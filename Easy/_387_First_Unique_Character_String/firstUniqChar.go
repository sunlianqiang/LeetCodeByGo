package _387_First_Unique_Character_String

import "fmt"


func FirstUniqChar(s string) int {
	len1 := len(s)
	if 0 == len1 {
		return -1
	}
	var charMap map[byte]int
	charMap = make(map[byte]int)


	for i := 0; i < len1; i++ {
		_, ok := charMap[s[i]]

		if ok {
			charMap[s[i]] = len1
		} else {
			charMap[s[i]] = i
		}
	}

	fmt.Printf("map:%+v\n", charMap)

	var minPos int
	minPos = len1
	for _, v := range charMap {
		if minPos > v {
			minPos = v
		}
	}

	if len1 == minPos {
		return -1
	}
	return minPos
}