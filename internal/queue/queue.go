package queue

import (
	"github.com/rohan3011/taskqueue/internal/task"
)

type TaskQueue interface {
	Enqueue(task *task.Task) error
	Dequeue() (*task.Task, error)
	Acknowledge(task *task.Task) error
}
