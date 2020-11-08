package sorting

import "galgo/algorithms/utils"

func QuickSort(arr []int, start int, end int) {
	if start < end {
		pivot := partition(arr, start, end)
		QuickSort(arr, start, pivot-1)
		QuickSort(arr, pivot+1, end)
	}
}

func partition(arr []int, start int, end int) int {
	x := arr[end] // Pivot element
	i := start - 1
	for j := start; j < end; j++ {
		if arr[j] < x {
			i++
			utils.Swap(arr, i, j)
		}
	}
	utils.Swap(arr, i+1, end)
	return i + 1
}