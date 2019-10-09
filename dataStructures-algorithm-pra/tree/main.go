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
		Val: 2,
	}
	root.Right = &tree.TreeNode{
		Val: 3,
	}
	a := tree.InorderTraversal(root)
	fmt.Println(a)
	fmt.Println(-1 << 63)
}
