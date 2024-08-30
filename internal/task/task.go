package task

import "time"

type Task struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Payload   []byte    `json:"payload"`
	Status    string    `json:"status"`
	Retries   int       `json:"retires"`
	CreatedAt time.Time `json:"created_at"`
}

type TaskHandler interface {
	Handle(task *Task) error
}

const (
	StatusCompleted = "completed"
	StatusFailed    = "failed"
	StatusPending   = "pending"
)

const (
	TypeEmail = "email"
)
