package main

import "fmt"

// Palindrome Number
// https://leetcode.com/problems/palindrome-number/description/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}

	oldNum := x
	newNum := 0

	// 1. find digit length of x
	digits := 10
	for x%digits != x {
		digits *= 10
	}

	// dump down by one digit
	digits /= 10

	for x != 0 {
		newNum += (x % 10) * digits

		digits /= 10
		x /= 10
	}

	fmt.Printf("OldNum: %d\nNewNum: %d\n", oldNum, newNum)

	return oldNum == newNum
}
