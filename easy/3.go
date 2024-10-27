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

/*
29. Divide Two Integers
Medium
Topics
Companies
Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.

The integer division should truncate toward zero, which means losing its fractional part. For example, 8.345 would be truncated to 8, and -2.7335 would be truncated to -2.

Return the quotient after dividing dividend by divisor.

Note: Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−231, 231 − 1]. For this problem, if the quotient is strictly greater than 231 - 1, then return 231 - 1, and if the quotient is strictly less than -231, then return -231.



Example 1:

Input: dividend = 10, divisor = 3
Output: 3
Explanation: 10/3 = 3.33333.. which is truncated to 3.
Example 2:

Input: dividend = 7, divisor = -3
Output: -2
Explanation: 7/-3 = -2.33333.. which is truncated to -2.
*/

func divide(dividend int, divisor int) int {
    // Handle edge case for overflow
    if dividend == math.MinInt32 && divisor == -1 {
        return math.MaxInt32
    }

    // Determine the sign of the result
    negative := (dividend < 0) != (divisor < 0)

    // Work with positive values for dividend and divisor
    dividendAbs := abs(dividend)
    divisorAbs := abs(divisor)
    
    // Initialize quotient
    quotient := 0

    // Use bit manipulation to find the quotient
    for dividendAbs >= divisorAbs {
        tempDivisor, multiple := divisorAbs, 1
        while dividendAbs >= (tempDivisor << 1) {
            tempDivisor <<= 1
            multiple <<= 1
        }
        dividendAbs -= tempDivisor
        quotient += multiple
    }

    // Adjust the sign of the quotient
    if negative {
        quotient = -quotient
    }

    return quotient
}


