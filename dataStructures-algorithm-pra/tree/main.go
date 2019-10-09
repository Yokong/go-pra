package main

import (
	"fmt"
	"go-pra/dataStructures-algorithm-pra/tree/tree"
)

func main() {
	root := &tree.TreeNode{
		Val: 1,
	}
	root.Left = &tree.TreeNode{
		Val:   2,
		Left:  &tree.TreeNode{Val: 4},
		Right: &tree.TreeNode{Val: 5},
	}
	root.Right = &tree.TreeNode{
		Val:  3,
		Left: &tree.TreeNode{Val: 6},
	}
	a := tree.PreorderTraversal(root)
	fmt.Println(a)
}
