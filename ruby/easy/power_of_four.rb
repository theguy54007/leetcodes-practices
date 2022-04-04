=begin
Given an integer n, return true if it is a power of four. Otherwise, return false.

An integer n is a power of four, if there exists an integer x such that n == 4x.



Example 1:

Input: n = 16
Output: true
Example 2:

Input: n = 5
Output: false
Example 3:

Input: n = 1
Output: true
=end

def is_power_of_four(n)
    return false if n <= 0

    n /= 4 while n % 4 == 0

    n == 1
end
