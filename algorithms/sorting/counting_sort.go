package sorting

import (
	"../utils"
)

func CountingSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	var result = make([]int, len(arr))
	var max = utils.Max(arr)

	var c = make([]int, max+1)
	for _, value := range arr {
		c[value]++
	}
	for i := 1; i < len(c); i++ {
		c[i] += c[i-1]
	}
	for j := len(arr) - 1; j >= 0; j-- {
		result[c[arr[j]]-1] = arr[j]
		c[arr[j]]--
	}

	return result
}
