package _38_Count_and_Say

import "strconv"

func countAndSay(n int) string {
	if 1 == n {
		return "1"
	}

	str1 := countAndSay(n-1)
	len1 := len(str1)

	var ret string
	var c byte
	c = str1[0]
	count := 1
	for i := 1; i < len1; i++ {
		if str1[i] == c {
			count++
			continue
		}
		ret = ret + strconv.Itoa(count) + string(c)
		c = str1[i]
		count = 1
	}
	ret = ret + strconv.Itoa(count) + string(c)
	return ret
}