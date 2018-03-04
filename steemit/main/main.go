package main

import (
	"steemit/data_structures/lists"
	//"fmt"
	"fmt"
)

func main() {
	list := lists.NewDoublyLinked()

	list.Insert(1)
	list.Insert(2)
	list.Insert(1)

	fmt.Print(list.IsPalindrome())
}
