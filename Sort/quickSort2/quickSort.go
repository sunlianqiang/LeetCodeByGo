package main

import "fmt"

func quickSort(arr []int, start, end int) {
	if start < end {
		part := partition(arr, start, end)
		quickSort(arr, start, part-1)
		quickSort(arr, part+1, end)
	}

}

func partition(arr []int, start, end int) (part int) {
	x := arr[start]
	i, j := start+1, end
	for i < j {
		for ; i < j && arr[j] >= x; j-- {
		}
		for ; i < j && arr[i] < x; i++ {
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[start] = arr[i]
	arr[i] = x
	fmt.Printf("arr sorting:%v, start:%v, end:%v, x:%v, i:%v\n", arr, start, end, x, i)
	return i
}

func main() {
	a := []int{10, 8, 4, 9, 3, 2, 1, 5, 7}
	fmt.Printf("a before sort:%v\n", a)
	quickSort(a, 0, len(a)-1)
	fmt.Printf("a after sort:%v\n", a)

}
