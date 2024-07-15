package linked_list

type LinkedList struct {
	head *Node[any]
}

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
