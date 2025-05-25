package main

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"sync"
)

type Queue interface {
	Len() int
	Enqueue(val int)
	Dequeue() (val int, err error)
}

type queue struct {
	list  []int
	mutex sync.Mutex
}

func NewQueue(cap int) Queue {
	return &queue{
		list:  make([]int, 0, cap),
		mutex: sync.Mutex{},
	}
}

func (q *queue) Len() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	return len(q.list)
}

func (q *queue) Enqueue(val int) {
	q.mutex.Lock()
	q.list = append(q.list, val)
	q.mutex.Unlock()
}

func (q *queue) Dequeue() (int, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if len(q.list) == 0 {
		return 0, errors.New("queue is empty")
	}

	val := q.list[0]
	q.list = slices.Clone(q.list[0 : len(q.list)-1])

	return val, nil
}

func main() {
	q := NewQueue(10)
	q.Enqueue(2)
	fmt.Println("enqueue:", 2)
	q.Enqueue(4)
	fmt.Println("enqueue:", 4)

	val, err := q.Dequeue()
	if err != nil {
		log.Println("failed to dequeue:", err)
	}
	fmt.Println("dequeue:", val)

	fmt.Println("queue len:", q.Len())
}
