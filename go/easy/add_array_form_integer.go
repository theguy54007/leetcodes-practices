/*
The array-form of an integer num is an array representing its digits in left to right order.

For example, for num = 1321, the array form is [1,3,2,1].
Given num, the array-form of an integer, and an integer k, return the array-form of the integer num + k.



Example 1:

Input: num = [1,2,0,0], k = 34
Output: [1,2,3,4]
Explanation: 1200 + 34 = 1234
Example 2:

Input: num = [2,7,4], k = 181
Output: [4,5,5]
Explanation: 274 + 181 = 455
Example 3:

Input: num = [2,1,5], k = 806
Output: [1,0,2,1]
Explanation: 215 + 806 = 1021
*/
package addToArrayForm

import (
	"strconv"
	"strings"
)

func addToArrayForm(num []int, k int) []int {
	k_arr := []int{}
	k_s_arr := strings.Split(strconv.Itoa(k), "")
	for _, ks := range k_s_arr {
		ki, _ := strconv.Atoi(ks)
		k_arr = append(k_arr, ki)
	}

	i := len(num) - 1
	j := len(k_arr) - 1
	carry := 0
	res := []int{}

	for i >= 0 || j >= 0 {
		sum := 0
		if i >= 0 {
			sum += num[i]
		}

		if j >= 0 {
			sum += k_arr[j]
		}

		sum += carry
		carry = sum / 10
		n := sum % 10
		res = append([]int{n}, res...)
		i--
		j--
	}

	if carry > 0 {
		res = append([]int{1}, res...)
	}
	return res

}
