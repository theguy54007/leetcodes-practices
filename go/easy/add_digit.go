/*
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
*/

package addDigits

import (
	"strconv"
	"strings"
)

func addDigits(num int) int {
	res := sumDigits(num)

	for res > 9 {
		res = sumDigits(res)
	}
	return res
}

func sumDigits(num int) int {
	sum := 0
	var num_arr = strings.Split(strconv.Itoa(num), "")

	for _, v := range num_arr {
		n, _ := strconv.Atoi(v)
		sum += n
	}
	return sum
}
