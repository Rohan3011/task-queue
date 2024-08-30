package worker

import (
	"fmt"
	"math"
	"time"

	"github.com/rohan3011/taskqueue/internal/task"
)

const (
	MaxRetries = 5
)

type Worker struct {
	ID       string
	TaskChan chan *task.Task
}

func (w *Worker) Start() {
	for t := range w.TaskChan {
		w.processTask(t)
	}
}

func (w *Worker) processTask(t *task.Task) {
	handler, err := task.GetHandler(t.Type)

	if err != nil {
		fmt.Printf("Worker %s failed to get handler for task %s: %v\n", w.ID, t.ID, err)
		return
	}

	for {
		err := handler.Handle(t)
		if err == nil {
			// Acknowledge task completion
			t.Status = task.StatusCompleted
			break
		}

		t.Retries++
		if t.Retries > MaxRetries {
			t.Status = task.StatusFailed
			break
		}

		backoff := time.Duration(math.Pow(2, float64(t.Retries))) * time.Second
		time.Sleep(backoff)
	}
}
