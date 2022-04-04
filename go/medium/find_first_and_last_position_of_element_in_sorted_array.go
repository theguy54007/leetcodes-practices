/*
Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.



Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
Example 3:

Input: nums = [], target = 0
Output: [-1,-1]
*/

package searchRange

func searchRange(nums []int, target int) []int {
	res := []int{}

	for i, n := range nums {
		if n == target {
			res = append(res, i)
		}
	}

	count := len(res)
	if count == 0 {
		return []int{-1, -1}
	}

	first := res[0]
	last := first + count - 1

	return []int{first, last}
}
