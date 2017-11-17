package _205_Isomorphic_Strings

func IsIsomorphic(s string, t string) bool {
	var byteMap map[byte]byte
	byteMap = make(map[byte]byte)
	var tMap map[byte]bool
	tMap = make(map[byte]bool)

	len1 := len(s)

	for i := 0; i < len1; i++ {
		tmp, ok := byteMap[s[i]]
		if ok {
			if tmp != t[i] {
				return false
			}
		} else {
			if s[i] != t[i] {
				_, ok1 := tMap[t[i]]
				if ok1 {
					return false
				}
			}

			byteMap[s[i]] = t[i]
			tMap[t[i]] = true
		}
	}

	return true
}
