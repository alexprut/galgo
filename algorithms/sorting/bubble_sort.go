package sorting

import (
	"../utils"
)

func BubbleSort(arr []int) []int {
	var swapped = true
	for swapped {
		swapped = false
		for i := 1; i < len(arr); i++ {
			if arr[i-1] > arr[i] {
				utils.Swap(arr, i, i-1)
				swapped = true
			}
		}
	}
	return arr
}
