package _345_Reverse_Vowels_String

import "fmt"

func isVowels(r rune) bool {
	if r == rune('a') || r == rune('e') || r == rune('i') || r == rune('o') || r == rune('u') || r == rune('A') || r == rune('E') || r == rune('I') || r == rune('O') || r == rune('U'){
		return true
	}
	return false
}

func ReverseVowels(s string) string {
	len1 := len(s)
	runeS := []rune(s)
	for i, j := 0, len1-1; i < j; i, j = i+1, j-1{
		for ;!isVowels(runeS[i]) && i < len1-1; i++ {}
		fmt.Printf("i:%+v\n", i)
		for ;!isVowels(runeS[j]) && j > 0; j-- {}

		if i < j {
			runeS[i], runeS[j] = runeS[j], runeS[i]
		}
	}

	ret := string(runeS)

	fmt.Printf("ret:%+v\n", ret)
	return ret
}
