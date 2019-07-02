package main

import (
	"fmt"
)

type Queue struct {
	Elem []interface{}
}

func CreateQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(item interface{}) {
	q.Elem = append(q.Elem, item)
}

func (q *Queue) Dequeue() interface{} {
	if len(q.Elem) == 0 {
		return -1
	}
	item := q.Elem[0]
	q.Elem = q.Elem[1:]
	return item
}

func main() {
	q := CreateQueue()
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	fmt.Println(q)
	i := q.Dequeue()
	fmt.Println(i)
	fmt.Println(q)
}
