=begin
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

A subarray is a contiguous part of an array.



Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Example 2:

Input: nums = [1]
Output: 1
Example 3:

Input: nums = [5,4,-1,7,8]
Output: 23
=end


def max_sub_array(nums)
  return 0 if nums.size == 0
  return nums[0] if nums.size == 1

  max = nums[0]
  curr = nums[0]

  i = 1
  while i < nums.size
      if curr > 0
          curr += nums[i]
      else
          curr = nums[i]
      end

      if curr > max
          max = curr
      end
      i += 1
  end
  max
end
