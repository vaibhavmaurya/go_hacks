package main

// Type declaration comes inside the square bracket

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

func (l LinkList[T]) isEmpty() bool {
	return l.head == nil
}

func (l LinkList[T]) getLastNode() *node[T] {

	if l.isEmpty() {
		return nil
	}
	currentNode := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	return currentNode
}

func (l LinkList[T]) getSecondLastNode() *node[T] {

	if l.isEmpty() || l.head.next == nil {
		return nil
	}
	currentNode := l.head
	for currentNode.next.next != nil {
		currentNode = currentNode.next
	}
	return currentNode
}

func (l *LinkList[T]) addNode(n T) {
	newNode := &node[T]{next: nil, prev: nil, value: n}

	if l.isEmpty() {
		l.head = newNode
	} else {
		currentNode := l.getLastNode()
		currentNode.next = newNode
	}
}

func (l *LinkList[int]) removeFront() bool {
	if l.isEmpty() {
		// TODO: Implement exception handling here
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

func (l *LinkList[T]) removeTail() bool {
	if l.isEmpty() {
		// TODO: Implement exception handling here
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

func (l LinkList[T]) traverse() []T {
	var items []T

	if l.isEmpty() {
		return items
	}

	for currentNode := l.head; currentNode != nil; currentNode = currentNode.next {
		items = append(items, currentNode.value)
	}
	return items
}
