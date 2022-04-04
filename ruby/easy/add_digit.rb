=begin
Given an integer num, repeatedly add all its digits until the result has only one digit, and return it.

Example 1:

Input: num = 38
Output: 2
Explanation: The process is
38 --> 3 + 8 --> 11
11 --> 1 + 1 --> 2
Since 2 has only one digit, return it.
Example 2:

Input: num = 0
Output: 0
=end


def add_digits(num)
    res = sum_digits(num)
    res = sum_digits(res) while res > 9
    res
end


def sum_digits(num)
    num_arr = num.to_s.split('').map(&:to_i)
    num_arr.sum
end
