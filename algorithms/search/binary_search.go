package search

func BinarySearch(a []int, v int) int {
	start := 0
	end := len(a) - 1

	for start <= end {
		middle := start + (end-start)/2
		if a[middle] == v {
			return middle
		}

		if a[middle] > v {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}

	return -1
}

func BinarySearchFirstLowest(a []int, v int) int {
	start := 0
	end := len(a) - 1

	for start <= end {
		middle := start + (end-start)/2
		if a[middle] == v {
			return middle
		}

		if v < a[middle] {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}

	return end
}

func BinarySearchFirstHighest(a []int, v int) int {
	start := 0
	end := len(a) - 1

	for start <= end {
		middle := start + (end-start)/2
		if a[middle] == v {
			return middle
		}

		if v < a[middle] {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}

	return start
}
