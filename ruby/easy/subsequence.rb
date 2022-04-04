=begin
Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).



Example 1:

Input: s = "abc", t = "ahbgdc"
Output: true
Example 2:

Input: s = "axc", t = "ahbgdc"
Output: false
=end

def is_subsequence(s, t)
    return true if s.size == 0
    s_chars = s.chars
    t.chars.each do |t_c|
        if t_c == s_chars[0]
            s_chars.shift
            return true if s_chars.empty?
        end
    end
    false
end
