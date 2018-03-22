package _506_Relative_Ranks

import (
	"sort"
	"strconv"
	"fmt"
)
type IntSlice []int

func (s IntSlice) Len() int { return len(s) }
func (s IntSlice) Swap(i, j int){ s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] > s[j] }

func FindRelativeRanks(nums []int) []string {
	len1 := len(nums)
	if 0 == len1 {
		return []string{}
	}else if 1 == len1{
		return []string{"Gold Medal"}
	}

	var rank []string

	var numsSorted []int
	numsSorted = make([]int , 0, len1)
	for _, v := range nums {
		numsSorted = append(numsSorted, v)
	}
	sort.Sort(IntSlice(numsSorted))
	fmt.Printf("sort:%+v\n", numsSorted)
	fmt.Printf("nums:%+v\n", nums)

	var rankMap map[int]string
	rankMap = make(map[int]string)

	if len1 >= 2 {
		rankMap[numsSorted[0]] = "Gold Medal"
		rankMap[numsSorted[1]] = "Silver Medal"
	}

	if len1 >= 3 {
		rankMap[numsSorted[2]] = "Bronze Medal"
		for i := 3; i < len1; i++ {
			rankMap[numsSorted[i]] = strconv.Itoa(i+1)
		}
	}

	fmt.Printf("rankMap:%+v\n", rankMap)


	for i := 0; i < len1; i++ {
		rank = append(rank, rankMap[nums[i]])
	}

	return rank
}