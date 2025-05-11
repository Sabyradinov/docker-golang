package http

import (
	"go-server/internal/handler"
)

type Handlers struct {
	help *handler.Help
}

func newHandler() Handlers {
	return Handlers{
		help: handler.NewHelpHandler(),
	}
}
