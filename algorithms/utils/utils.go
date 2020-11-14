package utils

import (
	"math"
)

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

func Min(arr []int) int {
	var min = arr[0]
	for _, value := range arr {
		if min > value {
			min = value
		}
	}
	return min
}

func MaxInt(a int, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func MinInt(a int, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func GetDigitAtIndex(number int, digitIndex int) int {
	return number / int(math.Pow(10, float64(digitIndex-1))) % 10
}

func Reverse(arr []int, start int, end int) {
	for i := start; i < (start+end)/2; i++ {
		Swap(arr, i, end-i)
	}
}

func IsPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}

func LongestIncreasingSubsequence(arr []int) []int {
	prev := make([]int, len(arr))
	dp := make([]int, len(arr))
	maxLength := 0
	indexMaxEnd := -1
	for i := 0; i < len(arr); i++ {
		dp[i] = 1
		prev[i] = -1
		for j := i; j >= 0; j-- {
			if dp[j]+1 > dp[i] && arr[j] < arr[i] {
				dp[i] = dp[j] + 1
				prev[i] = j
			}
			if dp[i] > maxLength {
				maxLength = dp[i]
				indexMaxEnd = i
			}
		}
	}
	lis := make([]int, maxLength)
	for i := maxLength - 1; i >= 0; i-- {
		lis[i] = arr[indexMaxEnd]
		indexMaxEnd = prev[indexMaxEnd]
	}
	return lis
}

func LongestCommonSubsequence(a string, b string) int {
	dp := make([][]int, len(a)+1, len(b)+1)
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = MaxInt(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}
