/*
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.



Example 1:

Input: nums = [2,2,1]
Output: 1
Example 2:

Input: nums = [4,1,2,1,2]
Output: 4
Example 3:

Input: nums = [1]
Output: 1
*/

package singleNumber

func singleNumber(nums []int) int {
	map_counts := make(map[int]int)
	res := 0
	for _, n := range nums {

		if _, ok := map_counts[n]; !ok {
			map_counts[n] = 0
		}

		map_counts[n] += 1
	}

	for n, c := range map_counts {
		if c == 1 {
			res = n
		}
	}

	return res
}
