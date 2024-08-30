package task

import "errors"

type TaskHandlerFactory func() TaskHandler

var handlers = make(map[string]TaskHandlerFactory)

func RegisterHandler(taskType string, factory TaskHandlerFactory) {
	handlers[taskType] = factory
}

func GetHandler(taskType string) (TaskHandler, error) {
	factory, ok := handlers[taskType]
	if !ok {
		return nil, errors.New("handler not found for task type")
	}
	return factory(), nil
}
