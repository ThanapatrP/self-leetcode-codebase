package main

// https://leetcode.com/problems/add-two-numbers/
// Status: Accepted

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// OLD VER
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	pointer1 := l1
// 	pointer2 := l2

// 	num1 := 0
// 	num2 := 0

// 	multer := 1

// 	for pointer1 != nil {
// 		num1 += pointer1.Val * multer
// 		pointer1 = pointer1.Next

// 		multer *= 10
// 	}

// 	multer = 1
// 	for pointer2 != nil {
// 		num2 += pointer2.Val * multer
// 		pointer2 = pointer2.Next

// 		multer *= 10
// 	}

// 	sum := num1 + num2
// 	var finalStartPointer *ListNode = nil
// 	var finalLastPointer *ListNode = nil

// 	if sum == 0 {
// 		return &ListNode{0, nil}
// 	}

// 	for sum > 0 {
// 		var digit int = sum % 10
// 		newL := ListNode{digit, nil}
// 		if finalStartPointer == nil {
// 			finalStartPointer = &newL
// 			finalLastPointer = &newL
// 		} else {
// 			finalLastPointer.Next = &newL
// 			finalLastPointer = finalLastPointer.Next
// 		}
// 		sum /= 10
// 	}

// 	return finalStartPointer
// }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pointer1 := l1
	pointer2 := l2

	var pointerToFinalStart *ListNode = nil
	var pointerToFinalLast *ListNode = nil

	hold := 0

	for pointer1 != nil || pointer2 != nil {
		ls1Val := 0
		ls2Val := 0

		if pointer1 != nil {
			ls1Val = pointer1.Val
		}else{
			ls1Val = 0
		}

		if pointer2 != nil {
			ls2Val = pointer2.Val
		}else{
			ls2Val = 0
		}

		res := ls1Val + ls2Val + hold
		hold = res / 10
		res %= 10

		if pointerToFinalStart == nil {
			pointerToFinalStart = &ListNode{res, nil}
			pointerToFinalLast = pointerToFinalStart
		}else{
			newDigit := ListNode{res, nil}
			pointerToFinalLast.Next = &newDigit
			pointerToFinalLast = pointerToFinalLast.Next
		}
		
		if pointer1 != nil {
			pointer1 = pointer1.Next
		}

		if pointer2 != nil {
			pointer2 = pointer2.Next
		}
	}

	if hold != 0 {
		newDigit := ListNode{hold, nil}
		pointerToFinalLast.Next = &newDigit
		pointerToFinalLast = pointerToFinalLast.Next

		hold = 0
	}

	return pointerToFinalStart
}
