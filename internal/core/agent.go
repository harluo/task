package core

import (
	"github.com/goexl/log"
	"github.com/goexl/task"
)

type Agent = task.Agent

func newAgent(logger log.Logger) *task.Agent {
	return task.New().Logger(logger).Build()
}
