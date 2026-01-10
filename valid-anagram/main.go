package main

// Time complexity
// Best case : O(1) jika return awal (s = "abc" t = "ab")
// Average case: O(n)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	freq := make(map[rune]int, len(s))

	for _, char := range s {
		freq[char]++
	}
	for _, char := range t {
		freq[char]--
	}

	for _, count := range freq {
		if count != 0 {
			return false
		}
	}
	return true
}
