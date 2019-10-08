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
	root1 := &tree.TreeNode{
		Val: 10,
	}
	root1.Left = &tree.TreeNode{
		Val: 2,
	}
	a := tree.IsSameTree(root, root1)
	fmt.Println(a)
}
