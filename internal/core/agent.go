package core

import (
	"github.com/goexl/log"
	"github.com/goexl/task"
)

func NewAgent(logger log.Logger) *task.Agent {
	return task.New().Logger(logger).Build()
}
