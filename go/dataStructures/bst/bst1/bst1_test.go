package bst1

import (
	"fmt"
	"testing"
)

func TestBST(t *testing.T) {
	values := []string{"d", "b", "c", "e", "a"}
	data := []string{"data", "beta", "count", "echo", "alpha"}

	tree := &Tree{}
	for i, v := range values {
		err := tree.Insert(v, data[i])
		if err != nil {
			t.Error(err)
		}
	}

	fmt.Print("Sorted Values: | ")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ":", n.Data, " | ") })
	fmt.Println()

	s := "d"
	d, found := tree.Find(s)
	if !found {
		t.Fatal("Not found")
	}
	fmt.Println("Found " + s + ": '" + d + "'")

	err := tree.Delete(s)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("After deleted '" + s + ": ")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()
}
