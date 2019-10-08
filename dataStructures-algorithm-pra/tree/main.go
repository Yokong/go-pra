package main

import (
	"fmt"
	"go-pra/dataStructures-algorithm-pra/tree/tree"
)

func main() {
	root := &tree.TreeNode{
		Val: 10,
	}
	root.Left = &tree.TreeNode{
		Val: 5,
	}
	// root.Right = &tree.TreeNode{
	// 	Val: 5,
	// }
	a := tree.IsSymmetric(root)
	fmt.Println(a)
}
