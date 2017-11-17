package _383_Ransom_Note


func CanConstruct(ransomNote string, magazine string) bool {
	var magazineMap map[rune]int
	magazineMap = make(map[rune]int)
	magazineRune := []rune(magazine)

	for _, v := range magazineRune {
		magazineMap[v]++
	}

	ransomNoteRune := []rune(ransomNote)
	for _, letter := range ransomNoteRune {
		tmp, ok := magazineMap[letter]
		if !ok || tmp <= 0 {
			return false
		} else {
			magazineMap[letter]--
		}
	}

	return true
}
