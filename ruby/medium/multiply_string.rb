=begin
Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.

Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.



Example 1:

Input: num1 = "2", num2 = "3"
Output: "6"
Example 2:

Input: num1 = "123", num2 = "456"
Output: "56088"
=end

def multiply(num1, num2)
    (int_to_str(num1) * int_to_str(num2)).to_s
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
