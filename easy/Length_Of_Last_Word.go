package easy

import (
	"strconv"
	"strings"
)

/*
58. Length of Last Word
Easy
Topics
Companies
Given a string s consisting of words and spaces, return the length of the last word in the string.

A word is a maximal
substring

	consisting of non-space characters only.

Example 1:

Input: s = "Hello World"
Output: 5
Explanation: The last word is "World" with length 5.
Example 2:

Input: s = "   fly me   to   the moon  "
Output: 4
Explanation: The last word is "moon" with length 4.
Example 3:

Input: s = "luffy is still joyboy"
Output: 6
Explanation: The last word is "joyboy" with length 6.

Constraints:

1 <= s.length <= 104
s consists of only English letters and spaces ' '.
There will be at least one word in s.
*/
func Length_Of_Last_Word(s string) int {
	n := len(s)
	result := 0
	for n > 0 {
		n--
		if s[n] != ' ' {
			result++
		} else if result > 0 {
			return result
		}
	}
	return result
}

func reverseWord(s string) string {
	revStrings := strings.Split(s, " ")

	for i, j := 0, len(revStrings)-1; i < j; i, j = i+1, j-1 {
		revStrings[i], revStrings[j] = revStrings[j], revStrings[i]
	}
	return strings.Join(revStrings, " ")
}

/*
Given a string s, return the longest
palindromic

substring

	in s
*/
func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	maxLen := 0
	res := ""
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			subStr := s[i:j]
			if isPalindrome(subStr) {
				if len(subStr) > maxLen {
					res = subStr
					maxLen = len(subStr)
				}
			}
		}
	}
	return res
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

/*
567. Permutation in String
Solved
Medium
Topics
Companies
Hint
Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.

In other words, return true if one of s1's permutations is the substring of s2.



Example 1:

Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Explanation: s2 contains one permutation of s1 ("ba").
Example 2:

Input: s1 = "ab", s2 = "eidboaoo"
Output: false
*/

func checkInclusion(s1 string, s2 string) bool {
	set1 := [26]int{}
	for _, char := range s1 {
		set1[char-97]++
	}
	set2 := [26]int{}
	l := 0
	for index, char := range s2 {
		set2[char-97]++

		if index >= len(s1) {
			remChar := s2[l]
			set2[remChar-97]--
			l++
		}
		if set1 == set2 {
			return true
		}
	}
	return false
}

// 1945. Sum of Digits of String After Convert
func getLucky(s string, k int) int {
	str := convertToStringNum(s)
	ans := repeat(str, k)
	return ans
}

func convertToStringNum(s string) string {
	strNum := ""
	for _, v := range s {
		num := strconv.Itoa(int(v - 'a' + 1))
		strNum += num
	}
	return strNum
}

func repeat(s string, k int) int {
	ans := 0
	for i := 0; i < k; i++ {
		sum := 0
		for _, v := range s {
			num, _ := strconv.Atoi(string(v))
			sum += num
		}
		s = strconv.Itoa(sum)
		ans = sum
	}
	return ans
}
