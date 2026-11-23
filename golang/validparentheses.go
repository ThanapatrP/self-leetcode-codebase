package main

import (
	"fmt"
	"slices"
)

// Valid Parentheses
// https://leetcode.com/problems/valid-parentheses/

type StackItem struct {
	val  rune
	prev *StackItem
}

type Stack struct {
	tos *StackItem
}

func pushStack(stack *Stack, val rune) {
	newItem := StackItem{val, stack.tos}
	stack.tos = &newItem
}

func popStack(stack *Stack) rune {
	if stack.tos == nil {
		return '\x00' // null char
	}

	poppedItem := stack.tos

	stack.tos = poppedItem.prev
	poppedItem.prev = nil

	return poppedItem.val
}

func peekTos(stack *Stack) rune {
	if stack.tos != nil {
		return stack.tos.val
	} else {
		return '\x00'
	}
}

func contains(arr []rune, c rune) bool {
	return slices.Contains(arr[:], c)
}

func matchParent(open rune, close rune) bool {
	if open == '(' && close == ')' {
		return true
	} else if open == '[' && close == ']' {
		return true
	} else if open == '{' && close == '}' {
		return true
	}

	return false
}

func printStack(stack *Stack) {
	pointer := stack.tos
	for pointer != nil {
		fmt.Print(pointer.val)
		pointer = pointer.prev
	}
	fmt.Println("]")
}

func isValid(s string) bool {
	openParent := [3]rune{'(', '[', '{'}
	closeParent := [3]rune{')', ']', '}'}

	stack := Stack{nil}

	for _, char := range s {
		// fmt.Printf("# iterate: %c\n", char)
		// fmt.Printf("----Stack---- \n")
		// printStack(&stack)
		// fmt.Printf("------------- \n")

		if contains(openParent[:], char) {
			pushStack(&stack, char)
			// fmt.Printf("pushed: %c\n", peekTos(&stack))
		} else if contains(closeParent[:], char) {
			if matchParent(peekTos(&stack), char) {
				// fmt.Printf("popped: %c\n", peekTos(&stack))
				popStack(&stack)
			} else {
				return false
			}
		} else {
			return false
		}
	}

	// fmt.Printf("%c\n", stack.tos.val)

	if stack.tos != nil {
		return false
	} else {
		return true
	}
}
