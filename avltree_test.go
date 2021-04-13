package memredis

import (
	"fmt"
	"testing"
)

func TestAvlTree_Insert(t *testing.T) {
	var a = &AvlTree{}
	a.Insert(10, "key")
	a.Insert(10, "v")
	a.Insert(11, "js")
	a.Insert(12, "jx")

	fmt.Println(a.MidOrderTraverse())
}
func TestAvlTree_Search(t *testing.T) {
	var a = &AvlTree{}
	a.Insert(10, "key")
	a.Insert(10, "v")
	a.Insert(11, "js")
	a.Insert(12, "jx")

	fmt.Println(a.Search(10))
	fmt.Println(a.Search(11))
	fmt.Println(a.Search(12))
	fmt.Println(a.Search(13))
}

func TestAvlTree_SearchNode(t *testing.T) {
	var a = &AvlTree{}
	for i := 0; i < (2 << 7); i++ {
		a.Insert(i, i)
	}
	fmt.Println(a.PreOrderTraverse())
	fmt.Println(a.MidOrderTraverse())
	fmt.Println(a.PostOrderTraverse())

	fmt.Println(a.SearchNode(0))
	fmt.Println(a.Search(0))
}
