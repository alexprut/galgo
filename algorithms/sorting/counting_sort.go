package sorting

import "galgo/algorithms/utils"

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

func CountingSortDigitIndex(arr []int, digitIndex int) []int {
	var result = make([]int, len(arr))
	c := [10]int{}
	for _, element := range arr {
		c[utils.GetDigitAtIndex(element, digitIndex)]++
	}
	for i := 1; i < len(c); i++ {
		c[i] += c[i - 1]
	}
	for j := len(arr) - 1; j >= 0; j-- {
		result[c[utils.GetDigitAtIndex(arr[j], digitIndex)] - 1] = arr[j]
		c[utils.GetDigitAtIndex(arr[j], digitIndex)]--
	}

	return result
}