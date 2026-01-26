package main

import "fmt"

type Node struct {
	val  int
	next *Node
}
type LinkedList struct {
	head *Node
}

func (L *LinkedList) print() {
	current := L.head
	if current == nil {
		fmt.Println("empty LL")
	}
	for current != nil {
		fmt.Println(current.val)
		current = current.next

	}

}
func (l *LinkedList) insertAtBeginning(val int) {
	current := l.head

	newnode := Node{val: val}
	l.head = &newnode
	l.head.next = current

}
func (l *LinkedList) deleteAtBeginning() {
	current := l.head

	if current == nil {
		fmt.Println("empty LL")
		return
	}

	l.head = current.next

}
func main() {
	myList := LinkedList{}

	myList.insertAtBeginning(10)

	myList.insertAtBeginning(30)
	myList.insertAtBeginning(20)
	myList.insertAtBeginning(50)
	myList.insertAtBeginning(40)
	myList.deleteAtBeginning()
	myList.print()

}
