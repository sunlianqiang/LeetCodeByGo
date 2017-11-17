package _409_Longest_Palindrome

func longestPalindrome(s string) int {
	var length int

	len1 := len(s)
	if len1 < 2 {
		return len1
	}
	sRune := []rune(s)
	var charMap map[rune]int
	charMap = make(map[rune]int)

	for _, v := range sRune {
		charMap[v]++
	}

	var middle bool
	for _, v := range charMap {
		if 0 != v % 2 {
			middle = true
		}
		length += v - v%2
	}

	if middle {
		length++
	}
	return length
}
