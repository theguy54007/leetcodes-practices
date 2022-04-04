=begin
A peak element is an element that is strictly greater than its neighbors.

Given an integer array nums, find a peak element, and return its index. If the array contains multiple peaks, return the index to any of the peaks.

You may imagine that nums[-1] = nums[n] = -∞.

You must write an algorithm that runs in O(log n) time.



Example 1:

Input: nums = [1,2,3,1]
Output: 2
Explanation: 3 is a peak element and your function should return the index number 2.
Example 2:

Input: nums = [1,2,1,3,5,6,4]
Output: 5
Explanation: Your function can return either index number 1 where the peak element is 2, or index number 5 where the peak element is 6.

=end
def find_peak_element(nums)
  return 0 if nums.size == 1

  if nums[0] > nums[1]
      return 0
  end

  if nums[-1] > nums[-2]
      return nums.size - 1
  end

  start_i = 1
  end_i = nums.size - 2
  mid = nil
  while start_i <= end_i
      mid = (start_i + end_i)/2
      if nums[mid] > nums[mid + 1] && nums[mid] > nums[mid-1]
          return mid

      elsif nums[mid] < nums[mid - 1]
          end_i = mid - 1
      else
          start_i = mid + 1
      end
  end

  mid
end
