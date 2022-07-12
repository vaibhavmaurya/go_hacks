package main

//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

import (
	"fmt"

	"go_hacks/DataStructure/linklist"
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

	fmt.Println(lst)

	lst.RemoveTail()
	lst.RemoveTail()

	linklist.ShowErrors()
}
