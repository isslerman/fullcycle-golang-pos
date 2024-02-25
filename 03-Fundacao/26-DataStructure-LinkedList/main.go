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

func main() {
	myList := linkedList{}
	node1 := &node{data: 10}
	node2 := &node{data: 20}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(&node{data: 30})

	myList.print()

}
