package main

import "slices"

// Time complexity: O(n)
// Total nya: O(n . k log k)
func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, str := range strs {
		strSorted := sortString(str)
		groups[strSorted] = append(groups[strSorted], str)
	}

	var res [][]string

	for _, strs := range groups {
		res = append(res, strs)
	}
	return res
}

// Time complexity: O(k log k)
func sortString(str string) string {
	chars := []rune(str)
	slices.Sort(chars)
	return string(chars)
}
