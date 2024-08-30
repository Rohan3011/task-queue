package emails

import (
	"fmt"
	"testing"
	"time"

	"github.com/rohan3011/taskqueue/internal/queue"
	"github.com/rohan3011/taskqueue/internal/task"
	"github.com/rohan3011/taskqueue/internal/worker"
)

func TestEmailTaskQueueProcessing(t *testing.T) {
	// Initialize the in-memory task queue
	taskQueue := queue.NewInMemoryQueue()

	// Register a mock handler for task processing
	task.RegisterHandler("email", func() task.TaskHandler {
		return &EmailTaskHandler{}
	})

	// Create a worker pool with 2 workers
	workerPool := worker.NewWorkerPool(2)
	workerPool.Start()

	// Create and enqueue tasks
	for i := 0; i < 5; i++ {
		task := &task.Task{
			ID:        fmt.Sprintf("task_%d", i),
			Type:      "email",
			Payload:   []byte(`{"recipient":"test@example.com","subject":"Hello","body":"This is a test email."}`),
			Status:    task.StatusPending,
			CreatedAt: time.Now(),
		}
		if err := taskQueue.Enqueue(task); err != nil {
			t.Fatalf("Failed to enqueue task: %v", err)
		}

		// Submit the task to the worker pool for processing
		workerPool.SubmitTask(task)
	}

	// Wait for all tasks to be processed
	time.Sleep(1 * time.Second)

	// Dequeue tasks and verify their status
	for i := 0; i < 5; i++ {
		tq, err := taskQueue.Dequeue()
		if err != nil {
			t.Fatalf("Failed to dequeue task: %v", err)
		}

		if tq == nil {
			t.Fatalf("Expected a task but got nil")
		}

		if tq.Status != task.StatusCompleted {
			t.Errorf("Expected task %s to be completed, but got status: %s", tq.ID, tq.Status)
		}
	}
}
