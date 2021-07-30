package events

import "mnimidamonbackend/frontend/global"

// This event describes the end of Server, User and Computer configuration.
var ConfirmConfig confirmConfig

type ConfirmConfigHandler interface {
	HandleConfirmConfig(config global.Config)
}

type confirmConfig struct {
	handlers []ConfirmConfigHandler
}

func (e *confirmConfig) Register(handler ConfirmConfigHandler) {
	e.handlers = append(e.handlers, handler)
}

func (e *confirmConfig) Trigger(config global.Config) {
	for _, handler := range e.handlers {
		go handler.HandleConfirmConfig(config)
	}
}

