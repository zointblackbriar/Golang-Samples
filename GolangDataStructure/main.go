package main

import (
	"fmt"
	"container/list"
)

type customQueue struct {
	queue *list.List
}

func(c *customQueue) Enqueue(value string) {
	c.queue.PushBack(value)
}

func (c *customQueue) Dequeue() error {
	if c.queue.Len() > 0 {
		element := c.queue.Front()
		c.queue.Remove(element)
	}

	return fmt.Errorf("Pop error")
}

func (c *customQueue) Front() (string error) {
	if(c.queue.Len() > 0) {
		if value, ok := c.queue.Front()
	}
}


func main() {
	fmt.Println("hello world")

	customQueue := &customQueue{
		queue: list.New(),
	}

	customQueue.Enqueue("hello")
	customQueue.Enqueue("world")
}