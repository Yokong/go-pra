package main

import "fmt"

type Item interface{}

type Stack struct {
	Elem Item
	Next *Stack
}

func (s *Stack) IsEmpty() bool {
	return s.Next == nil
}

func (s *Stack) Push(item Item) {
	if s.IsEmpty() {
		s.Next = &Stack{
			Elem: item,
			Next: nil,
		}
	} else {
		temp := s.Next
		s.Next = &Stack{
			Elem: item,
			Next: temp,
		}
	}
}

func (s *Stack) Pop() (Item, *Stack) {
	return s.Next.Elem, s.Next
}

func main() {
	s := new(Stack)
	s.Push(12)
	s.Push(30)
	s.Push(40)
	i, s := s.Pop()
	fmt.Println(i)
	i, s = s.Pop()
	fmt.Println(i)
}
