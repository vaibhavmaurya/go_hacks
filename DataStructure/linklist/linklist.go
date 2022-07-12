package linklist

import (
	"fmt"
	"go_hacks/DataStructure"
)

// Type declaration comes inside the square bracket

var errorList DataStructure.ErrorList = make(DataStructure.ErrorList, 0)

type node[T comparable] struct {
	next  *node[T]
	prev  *node[T]
	value T
}

type LinkList[T comparable] struct {
	head *node[T]
}

// type doubleLinkList[Q comparable] struct {
// 	head *node[Q]
// 	tail *node[Q]
// }

// Single LinkList

func (l LinkList[T]) String() string {
	return "I am Link List"
}

func (l LinkList[T]) IsEmpty() bool {
	return l.head == nil
}

func (l LinkList[T]) getLastNode() *node[T] {

	if l.IsEmpty() {
		return nil
	}
	currentNode := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	return currentNode
}

func (l LinkList[T]) getSecondLastNode() *node[T] {

	if l.IsEmpty() || l.head.next == nil {
		return nil
	}
	currentNode := l.head
	for currentNode.next.next != nil {
		currentNode = currentNode.next
	}
	return currentNode
}

func (l *LinkList[T]) AddNode(n T) {
	newNode := &node[T]{next: nil, prev: nil, value: n}

	if l.IsEmpty() {
		l.head = newNode
	} else {
		currentNode := l.getLastNode()
		currentNode.next = newNode
	}
}

func (l *LinkList[T]) AddNodeMultiple(n ...T) {
	start := 0
	if l.IsEmpty() {
		l.head = &node[T]{next: nil, prev: nil, value: n[start]}
		start++
	} else {
		currentNode := l.getLastNode()
		for ; start < len(n); start++ {
			currentNode.next = &node[T]{next: nil, prev: nil, value: n[start]}
			currentNode = currentNode.next
		}
	}
}

func (l *LinkList[int]) RemoveFront() bool {
	if l.IsEmpty() {
		// TODO: Implement exception handling here
		errorList = errorList.New("Remove Failed", "LinkList", "Already Empty")
		return false
	} else {
		if l.head.next == nil {
			l.head = nil
		} else {
			l.head = l.head.next
		}
		return true
	}
}

func (l *LinkList[T]) RemoveTail() bool {
	if l.IsEmpty() {
		// TODO: Implement exception handling here
		errorList = errorList.New("Remove Failed", "LinkList", "Already Empty")
		return false
	} else {
		if l.head.next == nil {
			l.head = nil
		} else {
			currentNode := l.getSecondLastNode()
			currentNode.next = nil
		}
		return true
	}
}

func (l LinkList[T]) Traverse() []T {
	var items []T

	if l.IsEmpty() {
		return items
	}

	for currentNode := l.head; currentNode != nil; currentNode = currentNode.next {
		items = append(items, currentNode.value)
	}
	return items
}

func ShowErrors() {
	fmt.Println("Show errors here: ", len(errorList))
	for _, e := range errorList.GetErrors() {
		fmt.Println(e)
	}
}
