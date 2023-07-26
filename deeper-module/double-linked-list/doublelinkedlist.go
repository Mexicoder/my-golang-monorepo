package double_linked_list

import (
	"errors"
)

type Node struct {
	data     string
	next     *Node
	previous *Node
}

func NewNode(val string) *Node {
	return &Node{
		data:     val,
		next:     nil,
		previous: nil,
	}
}

type DoublyLinkedList struct {
	length int
	head   *Node
	tail   *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}

// Append a new node onto the end of the linked list.
// assigns the new node to tail.next and updates tail with new node
// O(1) - Constant time
func (l *DoublyLinkedList) Append(val string) *DoublyLinkedList {
	newNode := NewNode(val)
	if l.length == 0 {
		l.head = newNode
		l.head.next = nil
		l.head.previous = nil
		l.tail = newNode
		l.tail.next = nil
		l.tail.previous = nil
	} else {
		newNode.previous = l.tail
		l.tail.next = newNode
		l.tail = newNode
	}
	l.length++

	return l
}

// Prepend n new node onto the beginning of the list.
// create a new node, assign current head to node and update head to new node
func (l *DoublyLinkedList) Prepend(val string) *DoublyLinkedList {
	newNode := NewNode(val)
	if l.length == 0 {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head.previous = newNode
		l.head = newNode
	}
	l.length++

	return l
}

// Insert the new value at the index specified
// O(n) - Linear time b/c of loop for inserting in between head and tail
// O(1) - Constant space
func (l *DoublyLinkedList) Insert(val string, index int) error {
	newNode := NewNode(val)
	// Check if index is out of range
	if index < 0 || index > l.length {
		return errors.New("out of range")
	}

	// Check if list is empty
	if l.length == 0 && index == 0 {
		l.head = newNode
		l.tail = newNode
		l.length++
		return nil
	}

	// Insert newNode at head
	if index == 0 && l.length > 0 {
		l.Prepend(val)
		return nil
	}

	// Insert newNode at tail
	if index == l.length {
		l.Append(val)
		return nil
	}

	// // Insert newNode in between
	//nodeCur := l.head
	//for i := 0; i < index-1; i++ {
	//	nodeCur = nodeCur.next
	//}

	leader, err := l.traverseToIndex(index - 1)
	if err != nil {
		return err
	}
	follower := leader.next
	leader.next = newNode
	newNode.next = follower
	follower.previous = newNode
	newNode.previous = leader
	l.length++
	return nil
}

// Delete node at index
// O(n) - Linear time b/c we use a loop to travers the list
func (l *DoublyLinkedList) Delete(index int) error {

	if l.length == 0 {
		return errors.New("error, empty list")
	}

	// remove head node
	if index == 0 {
		if l.length == 1 {
			l.head = nil
			l.tail = nil
		} else {
			l.head = l.head.next
			l.head.previous = nil
		}
		l.length--
		return nil
	}

	// remove a node after head
	node, err := l.traverseToIndex(index - 1)
	if err != nil {
		return err
	}
	if node.next == l.tail {
		node.next = nil
		l.length--
	} else {
		node.next.next.previous = node
		node.next = node.next.next
		l.length--
	}

	if l.length == 1 {
		l.tail = l.head
		l.tail.next = nil
		l.tail.previous = nil
		l.head.next = nil
		l.head.previous = nil
	}
	return nil
}

// PrintList returns all values in an Array
// O(n) - Linear time b/c we use a loop to travers and collect all values of the list
// O(n) - Linear space b/c we are populating an array
func (l *DoublyLinkedList) PrintList() []string {
	vals := make([]string, 0, l.length)
	nodeCur := l.head
	for nodeCur != nil {
		//var formattedVal string
		//if nodeCur.next == nil {
		//	formattedVal = fmt.Sprint(nodeCur.data, "{", nodeCur.previous.data, ",", nodeCur.next, "}")
		//} else if nodeCur.previous == nil {
		//	formattedVal = fmt.Sprint(nodeCur.data, "{", nodeCur.previous, ",", nodeCur.next.data, "}")
		//} else {
		//	formattedVal = fmt.Sprint(nodeCur.data, "{", nodeCur.previous.data, ",", nodeCur.next.data, "}")
		//}

		vals = append(vals, nodeCur.data)
		//vals = append(vals, formattedVal)
		nodeCur = nodeCur.next
	}
	return vals
}

// Search for the value by traversing the nodes in the linked list. returns the index of the value.
// if not found -1 is returned
// O(n) - Linear time b/c we use a loop to travers and find the value in the list
// O(1) - Constant space
func (l *DoublyLinkedList) Search(val string) int {
	nodeCur := l.head
	for i := 0; i < l.length; i++ {
		if nodeCur.data == val {
			return i
		}
		nodeCur = nodeCur.next
	}
	return -1
}

// traverseToIndex
func (l *DoublyLinkedList) traverseToIndex(index int) (*Node, error) {
	if index > l.length {
		return nil, errors.New("out of bounds")
	}
	// TODO: add: if index > length/2 traverse backwards... making this O(n/2)
	currNode := l.head
	for i := 0; i < l.length; i++ {
		if i == index {
			return currNode, nil
		}
		currNode = currNode.next
	}
	return nil, errors.New("something went wrong")
}

// Length returns the up-to-date length property
// O(1) - Constant time
// O(1) - Constant space
func (l *DoublyLinkedList) Length() int {
	return l.length
}
