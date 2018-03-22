package _242_Valid_Anagram

import "sort"

type RuneSlice []rune

func (s RuneSlice) Len() int           { return len(s) }
func (s RuneSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s RuneSlice) Less(i, j int) bool { return s[i] < s[j] }

func isAnagram(s string, t string) bool {
	s1 := []rune(s)
	t1 := []rune(t)

	len1 := len(s1)
	len2 := len(t1)

	if len1 != len2 {
		return false
	}

	sort.Sort(RuneSlice(s1))
	sort.Sort(RuneSlice(t1))
	for i := 0; i < len1; i++{
		if s1[i] != t1[i] {
			return false
		}
	}

	return true
}
