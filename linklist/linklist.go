package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}
type linklist struct {
	head *Node
}

func (list *linklist) insertAtEnd(data int) {
	newNode := &Node{data: data}
	//check list is empty
	if list.head == nil {
		list.head = newNode
		return
	}
	//go to last

	temp := list.head

	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode

}

func (list *linklist) print() {
	current := list.head
	for current.next != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Printf("nil")

}

func (list *linklist) insertAtStart(data int) {
	newnode := &Node{data: data}

	//check link list is empty or not
	if list.head == nil {
		list.head = newnode
		return
	}

	current := list.head
	list.head = newnode
	list.head.next = current

}

func (list *linklist) reverse() {
	var prev *Node = nil
	curr := list.head
	future := curr.next
	for curr != nil {
		curr.next = prev
		prev = curr
		curr = future
		future = curr.next
	}
	list.head = prev

}

func main() {

	fmt.Println("create the linklist ")
	list := linklist{}
	// list.insertAtEnd(10)
	// list.insertAtEnd(20)
	// list.insertAtEnd(30)
	// list.insertAtEnd(40)
	// list.print()
	list.insertAtStart(90)
	list.insertAtStart(40)
	list.insertAtStart(50)
	list.insertAtStart(10)

	list.print()
	list.reverse()
}
