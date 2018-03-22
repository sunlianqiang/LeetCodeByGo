package _500_Keyboard_Row

import (
	"fmt"
	"strings"
)

var alphabetMap map[string]int

func init() {
	alphabetMap = make(map[string]int)
	row1 := []string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p"}
	row2 := []string{"a", "s", "d", "f", "g", "h", "j", "k", "l"}
	for _, v := range row1 {
		alphabetMap[v] = 1
	}
	for _, v := range row2 {
		alphabetMap[v] = 2
	}
}
func findWord(word string) (ret bool) {
	word = strings.ToLower(word)
	var allRow int
	allRow = alphabetMap[word[:1]]
	length := len(word)
	for i := 1; i < length; i++ {
		row := alphabetMap[word[i:i+1]]

		if allRow != row {
			return false
		}
	}

	return true
}

func FindWords(words []string) []string {
	var ret []string

	for _, v := range words {
		ok := findWord(v)
		if ok {
			ret = append(ret, v)
		}
	}
	fmt.Printf("ret:%+v\n", ret)
	return ret
}
