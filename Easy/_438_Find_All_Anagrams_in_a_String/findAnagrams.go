package _438_Find_All_Anagrams_in_a_String

import "fmt"

func FindAnagrams(s string, p string) []int {
	lenP := len(p)
	var ret []int
	lenS := len(s)
	if lenS < lenP {
		return []int{}
	}

	var charMap map[byte]int
	charMap = make(map[byte]int)
	for i := 0; i < lenP; i++ {
		charMap[p[i]]++
	}

	var sMap map[byte]int
	sMap = make(map[byte]int)
	for i := 0; i < lenP; i++ {
		sMap[s[i]]++
	}

	for i := 0; i < lenS - lenP + 1; i++ {
		fmt.Printf("%+v\n", s[i:i+lenP])
		isAnagrams := true

		var tmpMap map[byte]int
		tmpMap = make(map[byte]int)
		for k, v := range charMap {
			tmpMap[k] = v
		}

		if i > 0 {
			sMap[s[i+lenP-1]]++
			sMap[s[i-1]]--
		}

		for k, v := range sMap {
			if 0 == v {
				continue
			}
			_, ok := tmpMap[k]
			if !ok {
				isAnagrams = false
				break
			}
			tmpMap[k] -= v
		}
		if isAnagrams {
			for _, v := range tmpMap {
				if v > 0 {
					isAnagrams = false
					break
				}
			}
		}

		if isAnagrams {
			ret = append(ret, i)
		}
	}

	return ret
}