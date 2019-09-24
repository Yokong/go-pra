package main

import (
	"fmt"
)

type Elem int

type Tree struct {
	E     Elem
	Left  *Tree
	Right *Tree
}

func Find(item Elem, t *Tree) *Tree {
	if t == nil {
		return nil
	}
	if item < t.E {
		return Find(item, t.Left)
	} else if item > t.E {
		return Find(item, t.Right)
	} else {
		return t
	}
}

func Insert(item Elem, t *Tree) *Tree {
	if t == nil {
		return &Tree{E: item}
	} else if t.E == 0 {
		t.E = item
		return t
	} else if item < t.E {
		t.Left = Insert(item, t.Left)
	} else if item > t.E {
		t.Right = Insert(item, t.Right)
	}
	return t
}

func CreateTree() *Tree {
	return &Tree{}
}

func main() {
	t := CreateTree()
	t = Insert(10, t)
	t = Insert(20, t)
	t = Insert(5, t)
	fmt.Println(t)
}
