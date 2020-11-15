package algorithms

func KnuthMorrisPratt(text string, pattern string) int {
	if len(pattern) == 0 {
		return 0
	}
	n := len(text)
	m := len(pattern)
	p := computePrefixFunction(pattern)
	q := 0
	for i := 0; i < n; i++ {
		for q > 0 && pattern[q] != text[i] {
			q = p[q]
		}
		if pattern[q] == text[i] {
			q++
		}
		if q == m {
			// Pattern occurs
			return i - m + 1
		}
	}
	return -1
}

func computePrefixFunction(pattern string) []int {
	m := len(pattern)
	p := make([]int, m)
	p[0] = 0
	k := 0
	for q := 1; q < m; q++ {
		for k > 0 && pattern[k] != pattern[q] {
			k = p[k]
		}
		if pattern[k] == pattern[q] {
			k++
		}
		p[q] = k
	}
	return p
}
