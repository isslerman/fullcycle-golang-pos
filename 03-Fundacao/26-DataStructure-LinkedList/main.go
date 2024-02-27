package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
	size int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.size++
}

func (l *linkedList) print() {
	current := l.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
}

// DeleteNode deletes a node with the given value from the linked list
func (l *linkedList) DeleteNode(value int) {
	if l.head == nil {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		return
	}

	current := l.head
	for current.next != nil && current.next.data != value {
		current = current.next
	}

	if current.next != nil {
		current.next = current.next.next
	}
}

func main() {
	myList := linkedList{}
	node1 := &node{data: 10}
	node2 := &node{data: 20}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(&node{data: 30})

	myList.print()

}
