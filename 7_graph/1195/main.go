package main

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

func (tree *BinaryTree) Insert(value int) {
	if tree.Root == nil {
		tree.Root = &Node{Value: value}
	} else {
		tree.Root.Insert(value)
	}
}

func (node *Node) Insert(value int) {
	if node.Value > value {
		if node.Left == nil {
			node.Left = &Node{Value: value}
		} else {
			node.Left.Insert(value)
		}
	} else if node.Value < value {
		if node.Right == nil {
			node.Right = &Node{Value: value}
		} else {
			node.Right.Insert(value)
		}
	}
}

func main() {

}
