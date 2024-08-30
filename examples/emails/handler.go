package emails

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/rohan3011/taskqueue/internal/task"
)

type EmailTaskHandler struct{}

func (h *EmailTaskHandler) Handle(t *task.Task) error {
	var emailPayload struct {
		Recipient string `json:"recipient"`
		Subject   string `json:"subject"`
		Body      string `json:"body"`
	}
	if err := json.Unmarshal(t.Payload, &emailPayload); err != nil {
		return err
	}

	// Simulate email sending
	time.Sleep(time.Second * 100)
	fmt.Printf("Successfully sent email to %s", emailPayload.Recipient)

	return nil
}
