package events

import (
	"github.com/go-openapi/strfmt"
)

var ConfirmConfig confirmConfig

type ConfirmConfigPayload struct {
	FolderPath string
	Host       strfmt.Hostname
	Port       int
}

type ConfirmConfigHandler interface {
	HandleConfirmConfig(ccp ConfirmConfigPayload)
}

type confirmConfig struct {
	handlers []ConfirmConfigHandler
}

func (e *confirmConfig) Register(handler ConfirmConfigHandler) {
	e.handlers = append(e.handlers, handler)
}

func (e *confirmConfig) Trigger(ccp ConfirmConfigPayload) {
	for _, handler := range e.handlers {
		go handler.HandleConfirmConfig(ccp)
	}
}
