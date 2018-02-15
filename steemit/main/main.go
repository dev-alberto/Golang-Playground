package main

import (
	"fmt"
	"steemit/data_structures"
)

func main() {
	q := data_structures.NewQueue(17)

	for i := 10; i < 15; i++ {
		q.Enqueue(i)
	}

	q.Print()

	fmt.Println("********")

	q.Dequeue()
	q.Dequeue()

	q.Print()
	fmt.Println("********")

	q.Enqueue(15)

	q.Print()
	fmt.Println("********")
}
