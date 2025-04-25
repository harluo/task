package core

import (
	"github.com/harluo/di"
	"github.com/harluo/pangu"
	"github.com/harluo/task/internal/core"
)

func init() {
	di.New().Get().Dependency().Puts(
		core.NewAgent,
	).Build().Apply()
}
