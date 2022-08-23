package linklist

import (
	"fmt"
	"testing"
)

func TestLinkListSingleElement(t *testing.T) {
	lst := LinkList[string]{}

	lst.AddNode("Hello")

	check := lst.Traverse()

	switch {
	case len(check) != 1:
		t.Errorf("There must be one element returned by Traverse")
		fallthrough
	case check[0] != "Hello":
		t.Errorf("Traverese must return Hello as a first element")
	default:
		fmt.Println("SUCCESS")
	}

}

func TestLinkListRemoval(t *testing.T) {
	lst := LinkList[string]{}

	lst.AddNode("Hello")
	lst.AddNode("Check")

	check := lst.Traverse()

	switch {
	case len(check) != 1:
		t.Errorf("There must be one element returned by Traverse")
		fallthrough
	case check[0] != "Hello" || check[1] != "Check":
		t.Errorf("Traverese must return Hello as a first element and Check as second element")
	default:
		fmt.Println("SUCCESS")
	}

	c := lst.RemoveFront()
	check = lst.Traverse()

	switch {
	case !c || len(check) != 1:
		t.Errorf("There must be one element returned by Traverse")
		fallthrough
	case check[0] != "Check":
		t.Errorf("Hello is not removed")
	default:
		fmt.Println("SUCCESS")
	}

}
