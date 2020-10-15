package utils

import "math"

func Swap(arr []int, i int, j int) {
	var tmp = arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func Max(arr []int) int {
	var max = arr[0]
	for _, value := range arr {
		if max < value {
			max = value
		}
	}
	return max
}

func GetDigitAtIndex(number int, digitIndex int) int {
	return int(number /  int(math.Pow(10, float64(digitIndex - 1))) % 10)
}