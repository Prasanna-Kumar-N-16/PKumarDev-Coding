package easy

import "strings"

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
	return ""
}
