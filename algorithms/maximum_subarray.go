package algorithms

import "galgo/algorithms/utils"

func MaximumSubarray(elements []int) int {
	if len(elements) == 0 {
		return 0
	}

	maxEndingHere := elements[0]
	maxSoFar := elements[0]

	for i := 1; i < len(elements); i++ {
		maxEndingHere = utils.MaxInt(elements[i], maxEndingHere+elements[i])
		maxSoFar = utils.MaxInt(maxSoFar, maxEndingHere)
	}

	return maxSoFar
}
