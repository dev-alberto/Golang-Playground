package lists

import (
	"fmt"
)

type DoublyLinkedList struct {
	head *Node
	size int
}

func NewDoublyLinked() *DoublyLinkedList {
	return &DoublyLinkedList{size:0}
}

// IsEmpty checks if the list is empty by looking at its head
func (l *DoublyLinkedList) IsEmpty() bool {return l.head == nil}

// Insert add a node at the head of the list, a O(1) operation
func (l *DoublyLinkedList) Insert(item interface{}) {
	newNode := &Node{data:item}
	if l.IsEmpty() {
		l.head = newNode
		return
	}

	l.head.prev = newNode
	newNode.next = l.head
	l.head = newNode
}

// ListSearch performs a scan of the list and returns
// the node with the given key. It has O(n) worst case time
func (l *DoublyLinkedList) ListSearch(item interface{}) (*Node, bool) {
	for p := l.head; p != nil ; p = p.next {
		//fmt.Printf("Node, %v", p)
		if p.data == item {
			return p, true
		}
	}
	return nil, false
}

// Delete performs a removal of a node in O(1) time,
// by updating pointers
func (l *DoublyLinkedList) Delete(node *Node) {
	// if the node != head
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		l.head = node.next
	}
	// if the node != the last node
	if node.next != nil {
		node.next.prev = node.prev
	}
}

// RemoveItem remove node with key = item. Runs in O(n) time,
// as it needs to call ListSearch to locate the node
func (l *DoublyLinkedList) RemoveItem(item interface{}) {
	node, ok := l.ListSearch(item)
	if !ok {
		return
	}
	// fmt.Printf("Node found, %v", node)
	l.Delete(node)
}

func (l *DoublyLinkedList) IsPalindrome() bool {
	if l.IsEmpty() {
		return true
	}
	headNode := l.head
	lastNode := l.head
	// make the last node point to the tail of the list
	for lastNode.next != nil {
		lastNode = lastNode.next
	}
	// stoop the loop when the pointers meet at the middle of the
	// list
	for lastNode != headNode {
		if lastNode.data != headNode.data {
			return false
		}
		// one pointer goes in one direction, the other
		// goes in the opposite direction
		lastNode = lastNode.prev
		headNode = headNode.next
	}
	return true
}

func (l *DoublyLinkedList) HasLoop() bool {
	if l.IsEmpty() {
		return false
	}

	slowPointer := l.head
	fastPointer := l.head
	for slowPointer != nil && fastPointer != nil && fastPointer.next != nil {
		slowPointer = slowPointer.next
		fastPointer = fastPointer.next.next
		if slowPointer == fastPointer {
			return true
		}
	}
	return false
}

func (l *DoublyLinkedList) Print() {
	p := l.head
	str := ""
	for p != nil {
		str += fmt.Sprintf("%v <---> ", p.data)
		p = p.next
	}
	fmt.Print(str)
}

// Implement an algorithm for palindrome checking?
// Implement an algorithm to detect a loop in a doubly linked list
