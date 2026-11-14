package main

// https://leetcode.com/problems/add-two-numbers/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pointer1 := l1
	pointer2 := l2

	num1 := 0
	num2 := 0

	multer := 1

	for pointer1 != nil {
		num1 += pointer1.Val * multer
		pointer1 = pointer1.Next

		multer *= 10
	}

	multer = 1
	for pointer2 != nil {
		num2 += pointer2.Val * multer
		pointer2 = pointer2.Next

		multer *= 10
	}

	sum := num1 + num2
	var finalStartPointer *ListNode = nil
	var finalLastPointer *ListNode = nil

	if sum == 0 {
		return &ListNode{0, nil}
	}

	for sum > 0 {
		var digit int = sum % 10
		newL := ListNode{digit, nil}
		if finalStartPointer == nil {
			finalStartPointer = &newL
			finalLastPointer = &newL
		} else {
			finalLastPointer.Next = &newL
			finalLastPointer = finalLastPointer.Next
		}
		sum /= 10
	}

	return finalStartPointer
}
