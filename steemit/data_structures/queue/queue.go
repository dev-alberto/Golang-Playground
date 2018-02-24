package queue

import (
	"github.com/pkg/errors"
	"fmt"
)

type Queue struct {
	items []int
	cap int
	length int
	head int
	tail int
}

func NewQueue(cap int) *Queue {
	return &Queue{items:make([]int, cap), cap:cap, length:0, head:0, tail: 0}
}

func (q *Queue) IsEmpty() bool {
	return q.length == 0
}

func (q *Queue) IsFull() bool {
	return q.length == q.cap
}

func (q *Queue) Enqueue(item int) error {
	if q.IsFull() {
		return errors.New("Queue is already at full capacity")
	}
	q.items[q.tail] = item
	q.tail ++
	q.length ++
	return nil
}

// return the front of the queue, mutating the queue
func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Cannot dequeue an empty queue")
	}
	result := q.items[q.head]
	q.head ++
	q.length --
	return result, nil
}

func (q *Queue) Len() int {
	return q.length
}

// Returns the value at the front, without mutation
func (q *Queue) Poll() int {
	return q.items[q.head]
}

// Returns the oldest value in the queue, without mutation
func (q *Queue) Peek() int {
	return q.items[q.tail]
}

func (q *Queue) Print() {
	for i := q.head; i < q.tail; i++ {
		fmt.Println(q.items[i])
	}
}