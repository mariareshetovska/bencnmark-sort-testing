package main

import (
	"fmt"
	_ "net/http/pprof"
	"sort"
)

func BubbleSort(arr []int) []int {
	keepWorking := true
	for keepWorking {
		keepWorking = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				keepWorking = true
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
	return arr
}

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[0]
	var left, right []int
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	left = QuickSort(left)
	right = QuickSort(right)
	arr = append(append(left, pivot), right...)
	return arr
}

func Sort(arr []int) []int {
	sort.Ints(arr)
	return arr
}

func main() {
	array := []int{1, 2, 54, 32, 44}

	fmt.Println(BubbleSort(array))
	fmt.Println(QuickSort(array))
	fmt.Println(Sort(array))
}
