package main

import "fmt"

func main() {
	lst := LinkList[int]{}
	lst.addNode(10)
	lst.addNode(11)
	lst.addNode(12)
	lst.addNode(13)
	lst.addNode(14)

	fmt.Println(lst.traverse())

	lst.removeFront()

	fmt.Println(lst.traverse())

	lst.removeTail()

	fmt.Println(lst.traverse())

	lst.removeFront()
	lst.removeFront()
	fmt.Println(lst.traverse())
	lst.removeFront()
	lst.removeTail()
	fmt.Println(lst.traverse())
}
