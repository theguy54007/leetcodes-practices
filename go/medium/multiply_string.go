/*
Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.

Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.



Example 1:

Input: num1 = "2", num2 = "3"
Output: "6"
Example 2:

Input: num1 = "123", num2 = "456"
Output: "56088"
*/

package multiply

import "strconv"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	num3 := make([]int, len(num1)+len(num2))
	res := ""

	for i := 0; i < len(num1); i++ {
		for j := 0; j < len(num2); j++ {
			num3[i+j+1] += int((num1[i] - 48) * (num2[j] - 48))
		}
	}

	for i := len(num3) - 1; i > 0; i-- {
		num3[i-1] += num3[i] / 10
		num3[i] %= 10
	}

	for i := 0; i < len(num3); i++ {
		v := num3[i]
		if !(i == 0 && v == 0) {
			res += strconv.Itoa(v)
		}

	}

	return res
}
