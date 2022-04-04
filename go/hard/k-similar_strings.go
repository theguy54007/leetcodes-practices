/*
Strings s1 and s2 are k-similar (for some non-negative integer k) if we can swap the positions of two letters in s1 exactly k times so that the resulting string equals s2.

Given two anagrams s1 and s2, return the smallest k for which s1 and s2 are k-similar.



Example 1:

Input: s1 = "ab", s2 = "ba"
Output: 1
Example 2:

Input: s1 = "abc", s2 = "bca"
Output: 2
*/

package kSimilarity

func kSimilarity(s1 string, s2 string) int {
	if s1 == s2 {
		return 0
	}

	queue := []string{s1}

	seen_maps := make(map[string]bool)
	seen_maps[s1] = true

	str_size := len(s1)

	step := 0

	for {
		step++

		for total_q := len(queue); total_q > 0; total_q-- {
			t_s1 := queue[0]
			queue = queue[1:]

			i := 0
			for t_s1[i] == s2[i] {
				i++
			}

			for j := i + 1; j < str_size; j++ {
				if t_s1[i] == s2[i] || t_s1[i] != s2[j] {
					continue
				}

				tmp_s := swap(t_s1, i, j)

				if tmp_s == s2 {
					return step
				}

				if !seen_maps[tmp_s] {
					seen_maps[tmp_s] = true
					queue = append(queue, tmp_s)
				}
			}

		}

	}
}

func swap(s string, i, j int) string {
	bs := []byte(s)

	bs[i], bs[j] = bs[j], bs[i]

	return string(bs)
}
