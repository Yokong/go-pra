package tree

import (
	"testing"
)

func TestTree(t *testing.T) {
	tree := New()
	tree.Insert(11)
	t.Log(tree.Left)
}
