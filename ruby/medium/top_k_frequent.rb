=begin
Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.



Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
Example 2:

Input: nums = [1], k = 1
Output: [1]
=end

def top_k_frequent(nums, k)
    count_hash = {}
    nums.each do |n|
        count_hash[n] ||= 0
        count_hash[n] += 1
    end
    count_hash.sort_by{|d, v| -v}.map{|d| d[0]}[0..(k - 1)]
end
