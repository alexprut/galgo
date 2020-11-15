package algorithms

import "testing"

func TestShouldFindMatch(t *testing.T) {
	AssertEquals(t, 0, RabinKarp("2359023141526739921", "", 10, 13))
	AssertEquals(t, 12, RabinKarp("2359023141526739921", "67399", 10, 13))
	AssertEquals(t, 12, RabinKarp("2359023141526739921", "67399", 10, 15))
	AssertEquals(t, 0, RabinKarp("2359023141526739921", "235", 10, 15))
	AssertEquals(t, -1, RabinKarp("2359023141526739921", "6327", 10, 15))
	AssertEquals(t, 4, RabinKarp("search here", "ch", 127, 893933))
	AssertEquals(t, -1, RabinKarp("search here", "hr", 127, 893933))
}
