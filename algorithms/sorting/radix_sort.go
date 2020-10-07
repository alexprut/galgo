package sorting

func RadixSort(arr []int, maxDigits int) {
	for i := 1; i <= maxDigits; i++ {
		CountingSortDigitIndex(arr, i)
	}
}