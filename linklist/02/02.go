package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type linklist struct {
	head *Node
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

func (list linklist) print() {
	curr := list.head
	for curr != nil {
		fmt.Println(curr)
		curr = curr.next
	}
	fmt.Println("nil")
}
func main() {

	fmt.Println("reverse the linklist")

	//list := linklist{}
	// list.reverse(1)
	// list.reverse(2)
	// list.reverse(3)
	// list.reverse(4)

	// fmt.Println("Original List:")
	// list.Print()

	// list.reverse()

	// fmt.Println("Reversed List:")
	// list.Print()

}
