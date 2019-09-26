package main

import "fmt"

const (
	Empty int = -1
)

type Stack struct {
	Top  int
	Elem []interface{}
}

func CreateStack() *Stack {
	return &Stack{Top: -1}
}

func (s *Stack) Push(item interface{}) {
	s.Top++
	s.Elem = append(s.Elem, item)
}

func (s *Stack) Pop() interface{} {
	if s.Top == Empty {
		return Empty
	}
	item := s.Elem[s.Top]
	s.Elem = s.Elem[:s.Top]
	s.Top--
	return item
}

func main() {
	s := CreateStack()
	s.Push(10)
	s.Push(11)
	s.Push(12)
	s.Push(13)
	s.Push(14)
	s.Push(15)
	s.Push(16)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
