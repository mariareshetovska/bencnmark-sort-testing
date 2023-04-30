package sort

import (
	"testing"
)

func BenchmarkBubbleSort(b *testing.B) {
	arr := []int{3, 6, 4, 56, 34}

	for i := 0; i < b.N; i++ {
		BubbleSort(arr)
	}
}

func BenchmarkSort(b *testing.B) {
	arr := []int{3, 6, 4, 56, 34}

	for i := 0; i < b.N; i++ {
		Sort(arr)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	arr := []int{3, 6, 4, 56, 34}

	for i := 0; i < b.N; i++ {
		QuickSort(arr)
	}
}

func BenchmarkBubbleSort100K(b *testing.B) {
	arr := getarr(100000)

	for i := 0; i < b.N; i++ {
		BubbleSort(arr)
	}
}

func BenchmarkSort100K(b *testing.B) {
	arr := getarr(100000)

	for i := 0; i < b.N; i++ {
		Sort(arr)
	}
}

func BenchmarkQuickSort100K(b *testing.B) {
	arr := []int{3, 6, 4, 56, 34}

	for i := 0; i < b.N; i++ {
		QuickSort(arr)
	}
}

func getarr(n int) []int {
	result := make([]int, n)
	j := 0
	for i := n - 1; i > 0; i-- {
		result[j] = i
		j++
	}
	return result
}
