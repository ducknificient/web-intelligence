package main

type Queue struct {
	Queue []string
}

// NewQueue initializes a new queue
func NewQueue() *Queue {
	return &Queue{Queue: []string{}}
}

// Enqueue adds an item to the queue
func (q *Queue) Enqueue(item string) {
	q.Queue = append(q.Queue, item)
}

// Dequeue removes and returns an item from the queue
func (q *Queue) Dequeue() string {
	if q.IsEmpty() {
		return ""
	}
	item := q.Queue[0]
	q.Queue = q.Queue[1:]
	return item
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.Queue) == 0
}

// Contains checks if the queue contains a specific item
func (q *Queue) Contains(item string) bool {
	for _, value := range q.Queue {
		if value == item {
			return true
		}
	}
	return false
}
