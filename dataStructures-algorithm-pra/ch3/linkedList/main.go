package main

import "fmt"

type List struct {
	Item interface{}
	Next *List
}

func (l *List)IsEmpty() bool {
	return l.Next == nil
}

func main() {
	l := List{Item: 5}
	l.Next = &List{Item: "lala"}
	fmt.Println(l.IsEmpty())
}
