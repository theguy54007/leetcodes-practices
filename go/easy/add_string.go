/*
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
*/

package addStrings

func addStrings(num1 string, num2 string) string {

	res := ""
	i := len(num1) - 1
	j := len(num2) - 1
	carry := byte(0)

	for i >= 0 || j >= 0 {
		sum := byte(0)

		if i >= 0 {
			sum += num1[i] - '0'

		}

		if j >= 0 {
			sum += num2[j] - '0'

		}

		sum += carry

		carry = sum / 10

		n := sum % 10

		res = string(n+'0') + res
		i--
		j--
	}

	if carry > 0 {
		res = string(carry+'0') + res
	}

	return res

}
