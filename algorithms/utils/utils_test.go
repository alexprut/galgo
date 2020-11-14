package utils

import (
	"reflect"
	"testing"
)

func TestSwap(t *testing.T) {
	var arr = []int{1, 2, 0, 9, 5}
	var expected = []int{0, 2, 1, 9, 5}

	Swap(arr, 0, 2)

	AssertArrayIntEquals(t, expected, arr)
}

func TestMax(t *testing.T) {
	var arr = []int{1, 2, 0, 9, 5}

	AssertEquals(t, 9, Max(arr))
}

func TestShouldFindMin(t *testing.T) {
	arr := []int{1, 2, 0, 9, 5}
	AssertEquals(t, 0, Min(arr))
}

func TestShouldReverse(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	Reverse(arr, 0, len(arr)-1)
	AssertEquals(t, []int{5, 4, 3, 2, 1}, arr)
}

func TestShouldCheckPalindrome(t *testing.T) {
	AssertTrue(t, IsPalindrome(""))
	AssertTrue(t, IsPalindrome("s"))
	AssertTrue(t, IsPalindrome("asa"))
	AssertTrue(t, IsPalindrome("assa"))

	AssertFalse(t, IsPalindrome("as"))
	AssertFalse(t, IsPalindrome("asas"))
}

func TestShouldFindLongestIncreasingSubsequence(t *testing.T) {
	AssertArrayIntEquals(t, []int{3, 7, 40, 80}, LongestIncreasingSubsequence([]int{50, 3, 10, 7, 40, 80}))
	AssertArrayIntEquals(t, []int{1, 2}, LongestIncreasingSubsequence([]int{1, 2, 1}))
	AssertArrayIntEquals(t, []int{3}, LongestIncreasingSubsequence([]int{3, 2, 1}))
	AssertArrayIntEquals(t, []int{1, 2, 3}, LongestIncreasingSubsequence([]int{1, 2, 3}))
	AssertArrayIntEquals(t, []int{}, LongestIncreasingSubsequence([]int{}))
}

func TestShouldFindLongestCommonSubsequence(t *testing.T) {
	AssertEquals(t, 0, LongestCommonSubsequence("", ""))
	AssertEquals(t, 0, LongestCommonSubsequence("xyz", "abcad"))
	AssertEquals(t, 1, LongestCommonSubsequence("aaa", "abcd"))
	AssertEquals(t, 2, LongestCommonSubsequence("aaa", "abcad"))
}

func TestShouldGetDigitAtIndex(t *testing.T) {
	number := 32981
	AssertEquals(t, 1, GetDigitAtIndex(number, 1))
	AssertEquals(t, 8, GetDigitAtIndex(number, 2))
	AssertEquals(t, 9, GetDigitAtIndex(number, 3))
	AssertEquals(t, 2, GetDigitAtIndex(number, 4))
	AssertEquals(t, 3, GetDigitAtIndex(number, 5))
	AssertEquals(t, 0, GetDigitAtIndex(number, 6))
	AssertEquals(t, 0, GetDigitAtIndex(number, 7))

	AssertEquals(t, 9, GetDigitAtIndex(9, 1))
	AssertEquals(t, 0, GetDigitAtIndex(9, 2))
}

func AssertEquals(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %v, actual %v", expected, actual)
	}
}

func AssertTrue(t *testing.T, actual bool) {
	if !actual {
		t.Errorf("Expected %v, actual %v", true, actual)
	}
}

func AssertFalse(t *testing.T, actual bool) {
	if actual {
		t.Errorf("Expected %v, actual %v", false, actual)
	}
}

func AssertArrayIntEquals(t *testing.T, expected []int, actual []int) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, actual %v", expected, actual)
	}
}
