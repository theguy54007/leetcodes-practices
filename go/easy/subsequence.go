/*
Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).



Example 1:

Input: s = "abc", t = "ahbgdc"
Output: true
Example 2:

Input: s = "axc", t = "ahbgdc"
Output: false

*/

package isSubsequence

// solution 1
// func isSubsequence(s string, t string) bool {
// 	t_index_mappings := make(map[string][]int)

// 	t_arr := strings.Split(t, "")

// 	for i, c := range t_arr {
// 		if _, ok := t_index_mappings[c]; !ok {
// 			t_index_mappings[c] = []int{}
// 		}

// 		t_index_mappings[c] = append(t_index_mappings[c], i)
// 	}

// 	s_arr := strings.Split(s, "")
// 	prev := 0

// 	for _, c := range s_arr {
// 		idxs, ok := t_index_mappings[c]

// 		if !ok {
// 			return false
// 		}

// 		filtered_idxs := filter_slice(idxs, prev)

// 		if len(filtered_idxs) == 0 {
// 			return false
// 		}

// 		prev = filtered_idxs[0] + 1
// 	}
// 	return true
// }

// func filter_slice(nums []int, fv int) []int {
// 	filtered := []int{}
// 	for _, v := range nums {
// 		if v >= fv {
// 			filtered = append(filtered, v)
// 		}
// 	}
// 	return filtered
// }

// solution2
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
