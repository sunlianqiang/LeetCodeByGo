package main

import (
	"fmt"

)

func longestCommonPrefixTwo(str1 string, str2 string) (prefix string){
	str1Len := len(str1)
	str2Len := len(str2)

	var prefixLen int
	if str1Len > str2Len {
		prefixLen = str2Len
	}else {
		prefixLen = str1Len
	}

	prefix = str1[:prefixLen]
	for prefix != str2[:prefixLen] {
		prefixLen--
		prefix = str1[:prefixLen]
	}

	return prefix
}
func longestCommonPrefix1(strs []string) (commonPrefix string) {
	if 0 == len(strs)  {
		return ""
	}
	strsLen := len(strs)
	commonPrefix = strs[0]

	for i := 1; i< strsLen; i++ {
		commonPrefix = longestCommonPrefixTwo(commonPrefix, strs[i])
	}
		return commonPrefix
}

func longestCommonPrefix(strs []string) (commonPrefix string) {
	if 0 == len(strs)  {
		return ""
	}
	strsLen := len(strs)
	commonPrefix = strs[0]

	for i := 1; i< strsLen; i++ {
		str1Len := len(commonPrefix)
		str2 := strs[i]
		str2Len := len(str2)

		var prefixLen int
		if str1Len > str2Len {
			prefixLen = str2Len
		}else {
			prefixLen = str1Len
		}

		commonPrefix = commonPrefix[:prefixLen]
		for commonPrefix != str2[:prefixLen] {
			prefixLen--
			commonPrefix = commonPrefix[:prefixLen]
		}


	}
	return commonPrefix
}
func main() {
	//strs := []string{
	//	"122345",
	//	"123456",
	//	"1234567",
	//	"12345678"}
	//
	//commonStr := longestCommonPrefix(strs)
	strs := []string{}
	commonStr := longestCommonPrefix1(strs)

	fmt.Printf("commonStr:%+v\n", commonStr)
}


