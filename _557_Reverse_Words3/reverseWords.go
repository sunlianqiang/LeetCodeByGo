package _557_Reverse_Words3

import "strings"

func ReverseWord(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func ReverseWords(s string) string {
	wordArr := strings.Split(s, " ")
	var ret string
	length := len(wordArr)
	for i := 0; i < length; i++ {
		reverseWord := ReverseWord(wordArr[i])
		if 0 == i {
			ret = ret + reverseWord
		} else {
			ret = ret + " " + reverseWord
		}

	}

	return ret
}