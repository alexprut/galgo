package sorting

import "../utils"

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		var j = i
		for j > 0 && arr[j-1] > arr[j] {
			utils.Swap(arr, j, j-1)
			j--
		}
	}
	return arr
}
