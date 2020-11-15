package algorithms

import (
	"reflect"
	"testing"
)

func TestShouldSearchPattern(t *testing.T) {
	AssertEquals(t, 4, KnuthMorrisPratt("bacbababaabcbab", "ababa"))
	AssertEquals(t, 0, KnuthMorrisPratt("bacbab", "b"))
	AssertEquals(t, 0, KnuthMorrisPratt("bacbab", "ba"))
	AssertEquals(t, -1, KnuthMorrisPratt("bacbab", "x"))
	AssertEquals(t, -1, KnuthMorrisPratt("bacbab", "xay"))
	AssertEquals(t, 0, KnuthMorrisPratt("bacbab", "bacbab"))
	AssertEquals(t, 0, KnuthMorrisPratt("", ""))
}

func TestShouldCreatePrefixFunction(t *testing.T) {
	AssertArrayIntEquals(t, []int{0, 0, 1, 2, 3, 0, 1}, computePrefixFunction("ababaca"))
}

func AssertEquals(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %v, actual %v", expected, actual)
	}
}

func AssertArrayIntEquals(t *testing.T, expected []int, actual []int) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, actual %v", expected, actual)
	}
}
