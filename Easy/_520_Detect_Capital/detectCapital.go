package _520_Detect_Capital

import "strings"

func DetectCapitalUse(word string) bool {
	//case 1
	upper := strings.ToUpper(word)
	if upper == word {
		return true
	} else {
		//case 2
		lower := strings.ToLower(word)
		if lower == word {
			return true
		}

		//case 3
		lowerRune := []rune(lower)
		wordRune := []rune(word)

		lowerRune[0] = lowerRune[0] + ('A' - 'a')
		lower2 := string(lowerRune)
		word2 := string(wordRune)
		if lower2 == word2 {
			return true
		}
	}

	return false
}
