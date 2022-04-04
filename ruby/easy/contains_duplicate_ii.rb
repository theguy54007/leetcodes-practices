=begin
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
=end

def contains_nearby_duplicate(nums, k)
    return false if nums.size < 2
    map_counts = {}

    nums.each_with_index do |n,i|
      return true if map_counts[n] && i - map_counts[n] <= k
      map_counts[n] = i
    end

    return false
end
