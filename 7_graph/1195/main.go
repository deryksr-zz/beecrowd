package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
		tree.Root.insert(value)
	}
}

func (node *Node) insert(value int) {
	if node.Value > value {
		if node.Left == nil {
			node.Left = &Node{Value: value}
		} else {
			rightPlace := node.Left.findFather(value)
			rightPlace.insert(value)
		}
	} else {
		if node.Right == nil {
			node.Right = &Node{Value: value}
		} else {
			rightPlace := node.Right.findFather(value)
			rightPlace.insert(value)
		}
	}
}

func (node *Node) findFather(value int) *Node {
	if node.Value > value {
		if node.Left == nil {
			return node
		}
		return node.Left.findFather(value)
	} else {
		if node.Right == nil {
			return node
		}
		return node.Right.findFather(value)
	}
}

func (node *Node) getPreOrder() (result []int) {
	if node != nil {
		result = append(result, node.Value)
		result = append(result, node.Left.getPreOrder()...)
		result = append(result, node.Right.getPreOrder()...)
	}
	return result
}

func (node *Node) getInOrder() (result []int) {
	if node != nil {
		result = append(result, node.Left.getInOrder()...)
		result = append(result, node.Value)
		result = append(result, node.Right.getInOrder()...)
	}
	return result
}

func (node *Node) getPosOrder() (result []int) {
	if node != nil {
		result = append(result, node.Left.getPosOrder()...)
		result = append(result, node.Right.getPosOrder()...)
		result = append(result, node.Value)
	}
	return result
}

func (tree *BinaryTree) GetPreOrder() string {
	if tree.Root != nil {
		return covnertToString(tree.Root.getPreOrder())
	}
	return ""
}

func (tree *BinaryTree) GetInOrder() string {
	if tree.Root != nil {
		return covnertToString(tree.Root.getInOrder())
	}
	return ""
}

func (tree *BinaryTree) GetPosOrder() string {
	if tree.Root != nil {
		return covnertToString(tree.Root.getPosOrder())
	}
	return ""
}

func covnertToString(value []int) string {
	temp := make([]string, len(value))
	for key, value := range value {
		temp[key] = strconv.Itoa(value)
	}
	return strings.Join(temp, " ")
}

func main() {
	var totalCases, size, current, actual int

	fmt.Scanf("%d", &totalCases)
	actual = 1
	for totalCases > 0 {
		fmt.Scanf("%d", &size)

		bTree := BinaryTree{}
		for size > 0 {
			fmt.Scanf("%d", &current)
			bTree.Insert(current)
			size--
		}
		fmt.Printf("Case %d:\n", actual)
		fmt.Printf("Pre.: %s\n", bTree.GetPreOrder())
		fmt.Printf("In..: %s\n", bTree.GetInOrder())
		fmt.Printf("Post: %s\n\n", bTree.GetPosOrder())
		totalCases--
		actual++
	}

}
