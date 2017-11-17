package _455_Assign_Cookies

import (
	"sort"
)

func findContentChildren(g []int, s []int) int {
	var ret int
	sort.Ints(g)
	sort.Ints(s)

	len1 := len(g)
	len2 := len(s)
	for i, j := 0, 0; i < len1 && j < len2; {
		if g[i] <= s[j] {
			i++
			j++
			ret++
		} else {
			j++
		}
	}

	return ret
}