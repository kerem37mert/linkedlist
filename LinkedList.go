package linkedlist

import (
	"errors"
	"fmt"
)

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
		fmt.Println("nil")
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
func (L *LinkedList) Count() uint {
	var temp *Node[any] = L.head
	var count uint = 0

	for temp != nil {
		count++
		temp = temp.next
	}
	return count
}

// Remove Index
func (L *LinkedList) RemoveIndex(index uint) error {
	if L.Count() > index {

		if index == 0 {
			L.head = L.head.next
		} else {
			var temp *Node[any] = L.head

			for i := 1; i != int(index); i++ {
				temp = temp.next
			}

			temp.next = temp.next.next
		}

		return nil
	} else {
		return errors.New("Index out of range")
	}
}

// Remove First Value
func (L *LinkedList) Remove(value any) {
	if L.head == nil {
		return
	}

	if L.head.data == value {
		L.head = L.head.next
	} else {
		var temp *Node[any] = L.head

		for temp.next != nil {
			if temp.next.data == value {
				temp.next = temp.next.next
			}
			temp = temp.next
		}
	}
}
