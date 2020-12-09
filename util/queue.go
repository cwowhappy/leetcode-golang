package util

import "errors"

type QueueNode struct {
	Val interface{}
	next *QueueNode
}

type Queue struct {
	Head *QueueNode
	Tail *QueueNode
	Length int //队列允许存放元素的总个数
	Size int   //队列当前元素的个数
}

func (queue *Queue) Enqueue(val interface{}) error {
	if queue.IsFull() {
		return errors.New("队列已满")
	}
	if queue.Tail != nil {
		queue.Tail.next = &QueueNode{Val: val, next: nil}
		queue.Tail = queue.Tail.next
	} else {
		queue.Tail = &QueueNode{Val: val, next: nil}
		queue.Head = queue.Tail
	}
	queue.Size += 1
	return nil
}

func (queue *Queue) Dequeue() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, errors.New("队列为空")
	}
	val := queue.Head.Val
	queue.Head = queue.Head.next
	if queue.Head == nil {
		queue.Tail = nil
	}
	queue.Size -= 1
	return val, nil
}

func (queue *Queue) IsFull() bool {
	return queue.Length == queue.Size
}

func (queue *Queue) IsEmpty() bool {
	return queue.Size == 0
}
