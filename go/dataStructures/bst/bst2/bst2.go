package bst2

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (b *BST) Insert(value int) *BST {
	current := b
	for current != nil {
		if value < current.Value {
			if current.Left == nil {
				current.Left = &BST{Value: value}
				break
			} else {
				current = current.Left
			}
		} else {
			if current.Right == nil {
				current.Right = &BST{Value: value}
				break
			} else {
				current = current.Right
			}
		}
	}

	return b
}

func (b *BST) Contains(value int) bool {
	current := b
	for current != nil {
		if value < current.Value {
			current = current.Left
		} else if value > current.Value {
			current = current.Right
		} else {
			return true
		}
	}
	return false
}

func (b *BST) Remove(value int) *BST {
	b.remove(value, nil)
	return b
}

func (b *BST) remove(value int, parent *BST) {
	current := b
	for current != nil {
		if value < current.Value {
			parent = current
			current = current.Left
		} else if value > current.Value {
			parent = current
			current = current.Right
		} else {
			if current.Left != nil && current.Right != nil {
				current.Value = current.Right.getMinValue()
				current.Right.remove(current.Value, current)
			} else if parent == nil {
				if current.Left != nil {
					current.Value = current.Left.Value
					current.Right = current.Left.Right
					current.Left = current.Left.Left
				} else if current.Right != nil {
					current.Value = current.Right.Value
					current.Left = current.Right.Left
					current.Right = current.Right.Right
				}
			} else if parent.Left == current {
				if current.Left != nil {
					parent.Left = current.Left
				} else {
					parent.Left = current.Right
				}
			} else if parent.Right == current {
				if current.Left != nil {
					parent.Right = current.Left
				} else {
					parent.Right = current.Right
				}
			}
			break
		}
	}
}

func (b *BST) getMinValue() int {
	if b.Left == nil {
		return b.Value
	}
	return b.Left.getMinValue()
}
