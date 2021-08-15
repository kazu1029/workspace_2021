package bst2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func NewBST(value int) *BST {
	return &BST{Value: value}
}

func TestCase1(t *testing.T) {
	/*
					10
				  /	   \
				 5     15
			   / \     / \
			 2   5    13  22
		   /           \
		  1            14
	*/
	root := NewBST(10)
	root.Left = NewBST(5)
	root.Left.Left = NewBST(2)
	root.Left.Left.Left = NewBST(1)
	root.Left.Right = NewBST(5)
	root.Right = NewBST(15)
	root.Right.Left = NewBST(13)
	root.Right.Left.Right = NewBST(14)
	root.Right.Right = NewBST(22)

	root.Insert(12)
	/*
					10
				  /	   \
				 5     15
			   / \     / \
			 2   5    13  22
		   /         / \
		  1        12  14
	*/
	require.True(t, root.Right.Left.Left.Value == 12)

	root.Remove(10)
	/*
					12
				  /	   \
				 5     15
			   / \     / \
			 2   5    13  22
		   /           \
		  1            14
	*/
	require.True(t, root.Contains(10) == false)
	require.True(t, root.Value == 12)

	require.True(t, root.Contains(15))
}
