package main

// Time complexity: O(n*m)
func groupAnagrams(strs []string) [][]string {
	grouped := make(map[[26]byte][]string)

	for _, str := range strs {
		var key [26]byte
		for _, char := range str {
			key[char-'a']++
		}
		grouped[key] = append(grouped[key], str)
	}

	res := make([][]string, 0, len(grouped))
	for _, group := range grouped {
		res = append(res, group)
	}
	return res
}
