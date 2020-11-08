package datastructures

import (
	"testing"
)

func TestInsertFindAndRemove(t *testing.T) {
	var trie = NewTrie()

	AssertEquals(t, 0, trie.Size())

	trie.Insert("data")
	trie.Insert("date")
	trie.Insert("structure")
	trie.Insert("struct")
	trie.Insert("algorithm")
	trie.Insert("algo")

	AssertEquals(t, 6, trie.Size())
	AssertFalse(t, trie.Search(""))
	AssertFalse(t, trie.SearchPrefix(""))
	AssertTrue(t, trie.Search("algo"))
	AssertTrue(t, trie.Search("algorithm"))
	AssertTrue(t, trie.SearchPrefix("algorithm"))
	AssertTrue(t, trie.SearchPrefix("algorith"))
	AssertFalse(t, trie.SearchPrefix("algorithms"))
	AssertTrue(t, trie.SearchPrefix("structure"))
	AssertTrue(t, trie.SearchPrefix("data"))
	AssertEquals(t, 2, trie.SizePrefix("alg"))
	AssertEquals(t, 2, trie.SizePrefix("algo"))

	AssertFalse(t, trie.Remove(""))
	AssertFalse(t, trie.Remove("alg"))
	AssertFalse(t, trie.Remove("algos"))
	AssertEquals(t, 6, trie.Size())
	AssertTrue(t, trie.Remove("algo"))
	AssertEquals(t, 5, trie.Size())

	AssertFalse(t, trie.RemovePrefix(""))
	AssertTrue(t, trie.RemovePrefix("dat"))
	AssertEquals(t, 3, trie.Size())

	AssertTrue(t, trie.RemovePrefix("s"))
	AssertTrue(t, trie.RemovePrefix("a"))
	AssertEquals(t, 0, trie.Size())
}

func AssertEquals(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %v, actual %v", expected, actual)
	}
}

func AssertTrue(t *testing.T, actual bool)  {
	if !actual {
		t.Errorf("Expected %v, actual %v", true, actual)
	}
}

func AssertFalse(t *testing.T, actual bool)  {
	if actual {
		t.Errorf("Expected %v, actual %v", false, actual)
	}
}