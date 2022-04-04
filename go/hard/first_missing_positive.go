/*
Given an unsorted integer array nums, return the smallest missing positive integer.

You must implement an algorithm that runs in O(n) time and uses constant extra space.



Example 1:

Input: nums = [1,2,0]
Output: 3
Example 2:

Input: nums = [3,4,-1,1]
Output: 2
Example 3:

Input: nums = [7,8,9,11,12]
Output: 1
*/

package firstMissingPositive

func firstMissingPositive(nums []int) int {
	map_nums := make(map[int]bool)

	for _, n := range nums {
		map_nums[n] = true
	}

	var res int
	i := 1
	for i <= len(nums) {
		if _, ok := map_nums[i]; !ok {
			res = i
			break
		}
		i++
	}

	if res == 0 {
		res = i
	}
	return res
}
