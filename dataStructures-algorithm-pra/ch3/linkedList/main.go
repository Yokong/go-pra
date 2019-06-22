package main

import (
	"fmt"
)

type Item interface{}

type List struct {
	Elem Item
	Next *List
}

func New(item Item) *List {
	return &List{Elem: item, Next: nil}
}

func (l *List) IsEmpty() bool {
	return l.Next == nil
}

func (l *List) Insert(item Item) {
	for l.Next != nil {
		l = l.Next
	}
	l.Next = &List{Elem: item}
}

func (l *List) Remove(item Item) {
	for l.Next != nil && l.Next.Elem != item {
		l = l.Next
	}
	if l.Next != nil {
		l.Next = l.Next.Next
	}
}

func (l *List) Str() string {
	var str string
	for l.Next != nil {
		str += fmt.Sprint(l.Elem)
		l = l.Next
		str += fmt.Sprint("->")
	}
	str += fmt.Sprint(l.Elem)
	return str
}

func main() {
	l := &List{Elem: 0}
	for i := 1; i < 10; i++ {
		l.Insert(i)
	}
	l.Remove(9)
	fmt.Println(l.Str())
}
