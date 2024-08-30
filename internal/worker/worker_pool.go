package worker

import (
	"fmt"

	"github.com/rohan3011/taskqueue/internal/task"
)

type WorkerPool struct {
	Workers   []*Worker
	TaskQueue chan *task.Task
}

func NewWorkerPool(workerCount int) *WorkerPool {
	wp := &WorkerPool{
		Workers:   make([]*Worker, workerCount),
		TaskQueue: make(chan *task.Task, 100),
	}

	for i := 0; i < workerCount; i++ {
		wp.Workers[i] = &Worker{
			ID:       fmt.Sprintf("worker-%d", i),
			TaskChan: wp.TaskQueue,
		}

	}

	return wp
}

func (wp *WorkerPool) Start() {
	for _, worker := range wp.Workers {
		go worker.Start()
	}
}

func (wp *WorkerPool) SubmitTask(t *task.Task) {
	wp.TaskQueue <- t
}
