package core

import (
	"github.com/goexl/log"
	"github.com/goexl/task"
	"github.com/pangum/task/internal/config"
)

func NewAgent(config *config.Task, logger log.Logger) *task.Agent {
	return task.New().Logger(logger).Retries(config.Retries).Build()
}
