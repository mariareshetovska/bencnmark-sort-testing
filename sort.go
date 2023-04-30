package sort

import "sort"

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

// Sort
func Sort(arr []int) {
	sort.Ints(arr)
}
