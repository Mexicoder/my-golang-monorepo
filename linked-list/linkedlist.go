package linked_list

import (
	"errors"
)

type Node struct {
	data string
	next *Node
}

func NewNode(val string) *Node {
	return &Node{
		data: val,
		next: nil,
	}
}

type SingleLinkedList struct {
	length int
	head   *Node
	tail   *Node
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}

// Append a new node onto the end of the linked list.
// assigns the new node to tail.next and updates tail with new node
// O(1) - Constant time
func (l *SingleLinkedList) Append(val string) *SingleLinkedList {
	newNode := NewNode(val)
	if l.length == 0 {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
	l.length++

	return l
}

// Prepend n new node onto the beginning of the list.
// create a new node, assign current head to node and update head to new node
func (l *SingleLinkedList) Prepend(val string) *SingleLinkedList {
	newNode := NewNode(val)
	if l.length == 0 {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	}
	l.length++

	return l
}

// Insert the new value at the index specified
// O(n) - Linear time b/c of loop for inserting in between head and tail
// O(1) - Constant space
func (l *SingleLinkedList) Insert(val string, index int) error {
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

	// Insert newNode in between
	nodeCur := l.head
	for i := 0; i < index-1; i++ {
		nodeCur = nodeCur.next
	}
	nodeHolder := nodeCur.next
	nodeCur.next = newNode
	newNode.next = nodeHolder
	l.length++
	return nil
}

// Delete node at index
// O(n) - Linear time b/c we use a loop to travers the list
func (l *SingleLinkedList) Delete(index int) error {

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
		node.next = node.next.next
		l.length--
	}

	if l.length == 1 {
		l.tail = l.head
	}
	return nil
}

// PrintList returns all values in an Array
// O(n) - Linear time b/c we use a loop to travers and collect all values of the list
// O(n) - Linear space b/c we are populating an array
func (l *SingleLinkedList) PrintList() []string {
	vals := make([]string, 0, l.length)
	nodeCur := l.head
	for nodeCur != nil {
		//for i := 0; i < l.length; i++ {
		vals = append(vals, nodeCur.data)
		nodeCur = nodeCur.next
	}
	return vals
}

// Search for the value by traversing the nodes in the linked list. returns the index of the value.
// if not found -1 is returned
// O(n) - Linear time b/c we use a loop to travers and find the value in the list
// O(1) - Constant space
func (l *SingleLinkedList) Search(val string) int {
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
func (l *SingleLinkedList) traverseToIndex(index int) (*Node, error) {
	if index > l.length {
		return nil, errors.New("out of bounds")
	}

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
func (l *SingleLinkedList) Length() int {
	return l.length
}

func (l *SingleLinkedList) Reverse() *SingleLinkedList {
	reversedList := NewSingleLinkedList()

	if l.length == 0 {
		return l
	}

	if l.head.next == nil {
		return l
	}

	currNode := l.head
	for currNode != nil {
		reversedList.Prepend(currNode.data)
		currNode = currNode.next
	}
	return reversedList
}
func (l *SingleLinkedList) Reverse2() *SingleLinkedList {

	if l.length == 0 {
		return l
	}

	if l.head.next == nil {
		return l
	}

	newHead := l.tail
	var prev *Node
	curr := l.head

	for curr != nil {
		// re-point
		next := curr.next
		curr.next = prev
		// shift
		prev = curr
		curr = next
	}
	l.head = newHead
	l.tail = prev

	/////////////////////////
	//
	//first := l.head
	//l.tail = l.head
	//
	////second := l.head.next
	//second := first.next
	//for second != nil {
	//	// re-pointing
	//	temp := second.next
	//	second.next = first
	//	//shifting
	//	first = second // memory leak!!!
	//	second = temp
	//}
	//l.head.next = nil
	//l.head = first

	return l
}

//func main() {
//	list := NewSingleLinkedList()
//
//	list.Append("hi")
//	list.Append("there")
//	list.Append("!!!")
//
//	fmt.Println(list.PrintList())
//	fmt.Println("length: ", list.Length())
//
//	fmt.Println(list.Search("there"))
//
//	list.Delete(1)
//
//	fmt.Println(list.PrintList())
//	fmt.Println("length: ", list.Length())
//
//	list.Insert("world", 1)
//
//	fmt.Println(list.PrintList())
//	fmt.Println("length: ", list.Length())
//
//	fmt.Println(list.Reverse().PrintList())
//	fmt.Println(list.Reverse2().PrintList())
//}
