package agent

import (
	"github.com/google/uuid"
)

type Job struct {
	UUID    string
	Command string
}

func NewJob(command string) Job {
	return Job{
		UUID:    uuid.New().String(),
		Command: command,
	}
}
