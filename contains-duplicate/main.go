package main

// Time complexity:
// Best Case: O(1) kalau duplicate ditemukan di awal nums = [1, 1]
// Average Case: O(n)
// Worst Case: O(n^2) kalau semua key kena hash collision (hampir tidak mungkin terjadi)
func containsDuplicate(nums []int) bool {
	s := make(map[int]struct{}, len(nums))

	for _, nums := range nums {
		if _, ok := s[nums]; ok {
			return true
		}
		s[nums] = struct{}{}
	}
	return false
}
