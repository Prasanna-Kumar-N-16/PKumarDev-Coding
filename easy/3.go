package easy

/*

You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.

Increment the large integer by one and return the resulting array of digits.



Example 1:

Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Incrementing by one gives 123 + 1 = 124.
Thus, the result should be [1,2,4].
Example 2:

Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
Incrementing by one gives 4321 + 1 = 4322.
Thus, the result should be [4,3,2,2].
Example 3:

Input: digits = [9]
Output: [1,0]
Explanation: The array represents the integer 9.
Incrementing by one gives 9 + 1 = 10.
Thus, the result should be [1,0].
*/
func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		// Increment the current digit by 1
		digits[i]++

		// If the digit becomes less than 10, no further carry is needed
		if digits[i] < 10 {
			return digits
		}

		// If the digit is 10, set it to 0 and carry over to the next
		digits[i] = 0
	}

	// If the loop completes, that means there was a carry from the most significant digit,
	// so we need to add an additional digit at the start (i.e., a leading 1).
	return append([]int{1}, digits...)
}
