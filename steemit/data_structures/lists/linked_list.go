package lists

import (
	"fmt"
	"errors"
)

type LinkedList struct {
	head *Node
	size int
}

func New() *LinkedList {
	return &LinkedList{size:0}
}

func NewWithValues(values ...interface{}) *LinkedList {
	newList := New()
	newList.Append(values...)
	return newList
}

// IsEmpty checks if there are any values stored in the list
func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

// Append function appends elements to the end of the list
func (l *LinkedList) Append(values ...interface{}) {
	for _, value := range values {
		node := &Node{data:value}
		if l.IsEmpty() {
			l.head = node
		} else {
			p, _ := l.getNodeAtIndex(l.size - 1)
			p.next = node
		}
		l.size ++
	}
}

// Prepend adds values at the end of the list
func (l *LinkedList) Prepend(values ...interface{}) {
	for _, value := range values {
		newNode := &Node{next:l.head, data:value}
		l.head = newNode
		l.size ++
	}
}

// Insert function inserts values at the specified index and shifts
// elements from that position onwards, if any, to the right
func (l *LinkedList) Insert(index int, values ...interface{}) error {
	if !l.isValidIndex(index) {
		return errors.New("index out of bounds")
	}
	// append at the end
	if index == l.size - 1 {
		l.Append(values...)
		return nil
	}
	listToAdd := NewWithValues(values...)
	lastNodeFromNewList, _ := listToAdd.getNodeAtIndex(listToAdd.size - 1)

	l.size += listToAdd.size
	if index == 0 {
		lastNodeFromNewList.next = l.head
		l.head = listToAdd.head
		return nil
	}
	nodeBeforeInsertIndex := l.getNodeBeforeIndex(index)
	lastNodeFromNewList.next = nodeBeforeInsertIndex.next
	nodeBeforeInsertIndex.next = listToAdd.head

	return nil
}

// Remove function deletes the node at the specified position
func (l *LinkedList) Remove(index int) error {
	if !l.isValidIndex(index) {
		return errors.New("index out of bounds")
	}
	if index == 0 {
		l.head = l.head.next
		return nil
	}
	p := l.getNodeBeforeIndex(index)
	p.next = p.next.next
	l.size --

	return nil
}

// Get function returns the node value at the specified index
func (l *LinkedList) Get(index int) (interface{}, bool) {
	nodeAtIndex, exists := l.getNodeAtIndex(index)
	return nodeAtIndex.data, exists
}


// Contains function checks to see if the value is contained in the list
func (l *LinkedList) Contains(value interface{}) bool {
	if l.IsEmpty() {
		return false
	}
	for p := l.head; p.next != nil; p = p.next {
		if p.data == value {
			return true
		}
	}
	return false
}

// Length function return the size of the list
func (l *LinkedList) Length() int {return l.size}

// Print outputs the array to stdout
func (l *LinkedList) Print() {
	p := l.head
	str := ""
	for p != nil {
		str += fmt.Sprintf("%v --> ", p.data)
		p = p.next
	}
	fmt.Print(str)
}

func (l *LinkedList) isValidIndex(index int) bool {
	return index >= 0 && index < l.size
}

func (l *LinkedList) getNodeBeforeIndex(index int) *Node {
	p := l.head
	for i := 0; i < index-1; i, p = i + 1, p.next {}
	return p
}


func (l *LinkedList) getNodeAtIndex(index int) (*Node, bool) {
	if !l.isValidIndex(index) {
		return nil, false
	}
	p := l.head
	// make the pointer point to the data at the required index
	for i := 0; i < index; i ++ {p = p.next}

	return p, true
}