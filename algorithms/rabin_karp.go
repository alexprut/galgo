package algorithms

import "math"

func RabinKarp(text string, pattern string, d int, q int) int {
	if len(pattern) == 0 {
		return 0
	}
	n := len(text)
	m := len(pattern)
	// Value of the high-order digit position of an m-digit window
	h := int(math.Pow(float64(d), float64(m-1))) % q
	p := 0
	t := 0
	for i := 0; i < m; i++ {
		// Using Horner's rule
		p = (d*p + int(pattern[i])) % q
		t = (d*t + int(text[i])) % q
	}
	for s := 0; s <= n-m; s++ {
		if p == t {
			// spurious hit
			isMatch := true
			for i := 0; i < m; i++ {
				if pattern[i] != text[s+i] {
					isMatch = false
				}
				if isMatch {
					return s
				}
			}
		}
		if s < n-m {
			prev := int(text[s])
			succ := int(text[s+m])
			t = (d*(t-prev*h) + succ) % q
			if t < 0 {
				t = q + t
			}
		}
	}
	return -1
}
