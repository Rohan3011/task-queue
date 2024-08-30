package queue

import (
	"sync"

	"github.com/rohan3011/taskqueue/internal/task"
)

type InMemoryQueue struct {
	tasks []*task.Task
	mutex sync.Mutex
}

func NewInMemoryQueue() *InMemoryQueue {
	return &InMemoryQueue{
		tasks: make([]*task.Task, 0),
	}
}

func (q *InMemoryQueue) Enqueue(t *task.Task) error {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.tasks = append(q.tasks, t)
	return nil
}

func (q *InMemoryQueue) Dequeue() (*task.Task, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if len(q.tasks) == 0 {
		return nil, nil
	}

	task := q.tasks[0]
	q.tasks = q.tasks[1:]
	return task, nil
}

func (q *InMemoryQueue) Acknowledge(t *task.Task) error {
	// Mark task as acknowledged, or remove it from the system if desired.
	return nil
}
