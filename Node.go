package linkedlist

type Node[T any] struct {
	data T
	next *Node[T]
}
