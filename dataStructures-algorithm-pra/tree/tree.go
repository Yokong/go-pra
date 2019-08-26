package tree

type Tree struct {
	Item  interface{}
	Left  *Tree
	Right *Tree
}

func New() *Tree { return &Tree{} }

func (t *Tree) Insert(Item interface{}) {
	if Item.(int) < t.Item.(int) {
		if t.Left == nil {
			t.Left = &Tree{Item: Item}
		} else {
			t.Left.Insert(Item)
		}
	} else {
		if t.Right == nil {
			t.Right = &Tree{Item: Item}
		} else {
			t.Right.Insert(Item)
		}
	}
}
