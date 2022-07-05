package main

import (
	"fmt"
	"go_hacks/data_structure/linklist"
)

func main() {
	lst := linklist.LinkList[int]{}
	lst.AddNode(10)
	lst.AddNode(11)
	lst.AddNode(12)
	lst.AddNode(13)
	lst.AddNode(14)

	fmt.Println(lst.Traverse())

	lst.RemoveFront()

	fmt.Println(lst.Traverse())

	lst.RemoveTail()

	fmt.Println(lst.Traverse())

	lst.RemoveFront()
	lst.RemoveFront()
	fmt.Println(lst.Traverse())
	lst.RemoveFront()
	lst.RemoveTail()
	fmt.Println(lst.Traverse())
}
