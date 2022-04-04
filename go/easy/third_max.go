/*
Given an integer array nums, return the third distinct maximum number in this array. If the third maximum does not exist, return the maximum number.



Example 1:

Input: nums = [3,2,1]
Output: 1
Explanation:
The first distinct maximum is 3.
The second distinct maximum is 2.
The third distinct maximum is 1.
Example 2:

Input: nums = [1,2]
Output: 2
Explanation:
The first distinct maximum is 2.
The second distinct maximum is 1.
The third distinct maximum does not exist, so the maximum (2) is returned instead.
Example 3:

Input: nums = [2,2,3,1]
Output: 1
Explanation:
The first distinct maximum is 3.
The second distinct maximum is 2 (both 2's are counted together since they have the same value).
The third distinct maximum is 1.
*/

package thirdMax

import "sort"

func thirdMax(nums []int) int {
	num_map := make(map[int]bool)
	for _, n := range nums {
		num_map[n] = true
	}

	uniq_nums := []int{}
	for n, _ := range num_map {
		uniq_nums = append(uniq_nums, n)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(uniq_nums)))

	if len(uniq_nums) > 2 {
		return uniq_nums[2]
	} else {
		return uniq_nums[0]
	}
}
