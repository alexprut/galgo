package sorting

import (
	"log"
	"math"
)

func MergeSort(arr []int, p int, r int) {
	if p < r {
		var q = (p + r) / 2
		MergeSort(arr, p, q)
		MergeSort(arr, q+1, r)
		merge(arr, p, q, r)
	}
}

func merge(arr []int, p int, q int, r int) {
	log.Println(arr)
	left := make([]int, q-p+2)
	right := make([]int, r-q+1)

	left[len(left)-1] = math.MaxInt32
	right[len(right)-1] = math.MaxInt32

	for i := 0; i < len(left) - 1; i++ {
		left[i] = arr[i+p]
	}

	for i := 0; i < len(right) - 1; i++ {
		right[i] = arr[i+q+1]
	}

	i := 0
	j := 0
	for k := p; k <= r; k++ {
		if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
	}
}
