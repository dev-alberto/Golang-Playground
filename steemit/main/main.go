package main

import (
	"steemit/data_structures/lists"
	//"fmt"
)

func main() {
	list := lists.NewWithValues(11, 22, 33)
	/*
	list.Insert(1, 1, 2, 3)
	list.Print()

	fmt.Println()
	list.Remove(3)

	list.Print()
	*/
	list.Append(1, 2, 3)
	list.Print()

}
