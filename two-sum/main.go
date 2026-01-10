package main

// If a + b = T
// then b = T - a

// Time complexity
// O(n)
func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int, len(nums))

	for i, num := range nums {
		complement := target - num

		if j, ok := indexMap[complement]; ok {
			return []int{j, i}
		}

		indexMap[num] = i
	}

	return nil
}
