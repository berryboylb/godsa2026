// DSA Daily — 2026-05-06
// Problem: https://leetcode.com/problems/valid-palindrome

package dsa06052026

import (
	"fmt"
	"unicode"
)

func isAlphaNumeric(r rune) bool {
    return unicode.IsLetter(r) || unicode.IsDigit(r)
}

// I hate working with strings best time so far is 0(N) USED TO O(N*2)
func solution(s string) bool {

	start, end := 0, len(s)-1

	for start < end {

		for start < end && !isAlphaNumeric(rune(s[start])) {
			start++
		}

		for start < end && !isAlphaNumeric(rune(s[end])) {
			end--
		}

		if unicode.ToLower(rune(s[start])) != unicode.ToLower(rune(s[end])) {
			fmt.Println("didn't match")
			return false
		}

		start++
		end--
	}

	fmt.Println("match")
	return true
}

func Main() {

	solution("0P")
}
