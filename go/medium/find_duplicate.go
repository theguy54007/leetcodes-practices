/*
Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.

There is only one repeated number in nums, return this repeated number.

You must solve the problem without modifying the array nums and uses only constant extra space.



Example 1:

Input: nums = [1,3,4,2,2]
Output: 2
Example 2:

Input: nums = [3,1,3,4,2]
Output: 3
*/

package findDuplicate

func findDuplicate(nums []int) int {
	mapping := make(map[int]int)

	for _, n := range nums {
		if _, ok := mapping[n]; !ok {
			mapping[n] = 0
		}

		mapping[n]++
	}

	var res int
	for n, c := range mapping {
		if c > 1 {
			res = n
			break
		}
	}

	return res
}
