package main

import "fmt"

// Queue - our queue of ints
type Queue struct {
	Elementes []int
}

func (q *Queue) Enqueue(elem int) {
	q.Elementes = append(q.Elementes, elem)
}

func (q *Queue) Dequeue() int {
	temp := q.Elementes[0]
	q.Elementes = q.Elementes[1:]
	return temp
}

func (q *Queue) Top() int {
	return q.Elementes[0]
}

func (q *Queue) Empty() bool {
	return len(q.Elementes) == 0
}

func main() {
	fmt.Println("Queue")
	queue := Queue{make([]int, 0)}
	queue.Enqueue(1)
	queue.Enqueue(4)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println(queue)
	fmt.Println("Top: ", queue.Top())
	fmt.Println("Dequeue: ", queue.Dequeue())
	fmt.Println("New Top: ", queue.Top())
	fmt.Println("Empty?", queue.Empty())
	fmt.Println("Dequeue: ", queue.Dequeue())
	fmt.Println("Dequeue: ", queue.Dequeue())
	fmt.Println("Dequeue: ", queue.Dequeue())
	fmt.Println("Empty?", queue.Empty())
}
