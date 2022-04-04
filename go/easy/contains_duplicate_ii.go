/*
Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.

Example 1:

Input: nums = [1,2,3,1], k = 3
Output: true
Example 2:

Input: nums = [1,0,1,1], k = 1
Output: true
Example 3:

Input: nums = [1,2,3,1,2,3], k = 2
Output: false
*/

package contains_duplicate_ii

func contains_duplicate_ii(nums []int, k int) bool {
	map_counts := make(map[int]int)
	for i, n := range nums {
		get_n, ok := map_counts[n]
		if ok && i-get_n <= k {
			return true
		} else {
			map_counts[n] = i
		}
	}
	return false
}
