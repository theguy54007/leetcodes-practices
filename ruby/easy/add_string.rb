=begin
Given two non-negative integers, num1 and num2 represented as string, return the sum of num1 and num2 as a string.

You must solve the problem without using any built-in library for handling large integers (such as BigInteger). You must also not convert the inputs to integers directly.



Example 1:

Input: num1 = "11", num2 = "123"
Output: "134"
Example 2:

Input: num1 = "456", num2 = "77"
Output: "533"
Example 3:

Input: num1 = "0", num2 = "0"
Output: "0"
=end

def add_strings(num1, num2)
    (int_to_str(num1) + int_to_str(num2)).to_s
end

def int_to_str(str)
    sum = 0
    multi_n = 1
    chars = str.chars

    i = chars.size - 1
    while i >= 0 do
        sum += (chars[i].ord - 48) * multi_n
        multi_n *= 10
        i -= 1
    end
    sum
end
