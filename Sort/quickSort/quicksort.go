package main

import "fmt"

func Partition(arr []int, start, end int) (index int) {
	var pivoit int = arr[end]
	index = start - 1
	for j := start; j < end; j++ {
		if arr[j] <= pivoit {
			index++
			arr[j], arr[index] = arr[index], arr[j]
		}
	}

	index++
	arr[index], arr[end] = arr[end], arr[index]
	return index
}
func QuickSort(arr []int, start int, end int) {
	if start < end {
		index := Partition(arr, start, end)

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
