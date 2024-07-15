package linked_list

import "fmt"

type LinkedList struct {
	head *Node[any]
}

// Add Node
func (L *LinkedList) Add(value any) {
	var newNode *Node[any] = new(Node[any])
	newNode.data = value
	newNode.next = nil

	if L.head == nil {
		L.head = newNode
	} else {
		var temp *Node[any] = L.head

		for temp.next != nil {
			temp = temp.next
		}
		temp.next = newNode
	}
}

// Linked List Traversal
func (L *LinkedList) Traversal() {
	if L.head == nil {
		fmt.Println("Your linked list is empty")
	} else {
		var temp *Node[any] = L.head

		for temp != nil {
			fmt.Printf("[%v] -> ", temp.data)
			temp = temp.next
		}
		fmt.Printf("nil")
	}
}

// Includes
func (L *LinkedList) Includes(value any) bool {
	var temp *Node[any] = L.head

	for temp != nil {
		if temp.data == value {
			return true
		}
		temp = temp.next
	}
	return false
}

// Element Count
func (L *LinkedList) Count() int {
	var temp *Node[any] = L.head
	var count int = 0

	for temp != nil {
		count++
		temp = temp.next
	}

	return count
}
