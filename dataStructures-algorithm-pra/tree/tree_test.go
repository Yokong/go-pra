package tree

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	tree := &Tree{Item: 10}
	tree.Insert(11)
	fmt.Println(tree.Right.Item)
	t.Log(tree)
}

func BenchmarkTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := &Tree{Item: 1}
		tree.Insert(i)
	}
}
