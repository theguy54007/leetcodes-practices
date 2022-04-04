=begin
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
=end

def single_number1(nums)
    map_count = {}
    nums.each do |n|
        map_count[n] ||= 0
        map_count[n] += 1
    end
    map_count.each do |k,v|
        return k if v == 1
    end

end


# method2
#   1. a ^ 0 = 0 ^ a = a
#   2. a ^ a = 0
#   3. a ^ b ^ a = a ^ a ^ b = 0 ^ b = b
def single_number2(nums)
  nums.reduce(&:^)
end
