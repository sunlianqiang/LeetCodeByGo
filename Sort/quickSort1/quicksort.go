package main

import "fmt"

func Partition(arr []int, l, r int) int {
	var i, j int = l, r
	x := arr[l]

	for i < j {
		// 从右向左找小于x的数
		for ; i < j && arr[j] >= x; j-- {
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}

		// 从左到右找大于x的数
		for ; i < j && arr[i] <= x; i++ {
		}

		if i < j {
			arr[j] = arr[i]
			j--
		}
	}

	arr[i] = x
	return i
}
func QuickSort(arr []int, start int, end int) {
	if start < end {
		index := Partition(arr, start, end)
		fmt.Printf("arr:%v\n", arr)
		if start < index {
			QuickSort(arr, start, index-1)
		}
		if index < end {
			QuickSort(arr, index+1, end)
		}
	}
}
func main() {
	var arr []int = []int{4, 2, 1, 3}

	QuickSort(arr, 0, 3)

	fmt.Printf("arr:%v\n", arr)
}
