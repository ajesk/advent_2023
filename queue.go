package main

func enqueue(queue [][]int, element []int) [][]int {
	queue = append(queue, element) // Simply append to enqueue.
	return queue
}

func dequeue(queue [][]int) [][]int {
	return queue[1:] // Slice off the element once it is dequeued.
}
