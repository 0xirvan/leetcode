package main

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	bucket := make([][]int, len(nums)+1)
	for num, f := range freq {
		bucket[f] = append(bucket[f], num)
	}

	res := make([]int, 0, k)

	for i := len(bucket) - 1; i >= 1 && len(res) < k; i-- {
		for _, num := range bucket[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}
