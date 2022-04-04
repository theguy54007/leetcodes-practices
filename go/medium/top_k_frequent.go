/*
Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.



Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
Example 2:

Input: nums = [1], k = 1
Output: [1]
*/

package topKFrequent

import "sort"

func topKFrequent(nums []int, k int) []int {
	map_freq := make(map[int]int)
	map_freq_sort := [][]int{}

	res := []int{}
	for _, n := range nums {
		if _, ok := map_freq[n]; !ok {
			map_freq[n] = 0
		}

		map_freq[n] += 1
	}

	for n, f := range map_freq {
		map_freq_sort = append(map_freq_sort, []int{n, f})
	}

	sort.Slice(map_freq_sort, func(i, j int) bool {
		return map_freq_sort[i][1] > map_freq_sort[j][1]
	})

	for _, f := range map_freq_sort {
		res = append(res, f[0])
		if len(res) >= k {
			break
		}
	}

	return res
}
