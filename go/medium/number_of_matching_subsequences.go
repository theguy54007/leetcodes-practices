/*
Given a string s and an array of strings words, return the number of words[i] that is a subsequence of s.

A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

For example, "ace" is a subsequence of "abcde".


Example 1:

Input: s = "abcde", words = ["a","bb","acd","ace"]
Output: 3
Explanation: There are three strings in words that are a subsequence of s: "a", "acd", "ace".
Example 2:

Input: s = "dsahjpjauf", words = ["ahjpjau","ja","ahbwzgqnuk","tnmlanowax"]
Output: 2
*/

package numMatchingSubseq

func numMatchingSubseq(s string, words []string) int {
	res := []bool{}
	seen := make(map[string]bool)
	for _, w := range words {
		is_sub, ok := seen[w]
		if !ok {
			is_sub = isSubsequence(w, s)
			seen[w] = is_sub
		}
		if is_sub {
			res = append(res, true)
		}
	}

	return len(res)
}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	j := 0

	for i := 0; i < len(t); i++ {
		if t[i] == s[j] {
			j++

			if j == len(s) {
				return true
			}
		}
	}
	return false
}
