package main

import (
	"fmt"
)

func main() {
	list1child2 := ListNode{9, nil}
	list1child1 := ListNode{9, &list1child2}
	list1 := ListNode{9, &list1child1}

	list2child2 := ListNode{9, nil}
	list2child1 := ListNode{9, &list2child2}
	list2 := ListNode{9, &list2child1}

	finalLn := addTwoNumbers(&list1, &list2)
	pointer := finalLn
	for pointer != nil {
		fmt.Print(pointer.Val)
		pointer = pointer.Next
	}
}
