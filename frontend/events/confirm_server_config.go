package events

import "mnimidamonbackend/frontend/global"

// Event describes the selection of server configurations.
var ConfirmServerConfig confirmServerConfig

type ConfirmServerConfigHandler interface {
	HandleServerConfirmConfig(payload global.ServerConfig)
}

type confirmServerConfig struct {
	handlers []ConfirmServerConfigHandler
}

func (e *confirmServerConfig) Register(handler ConfirmServerConfigHandler) {
	e.handlers = append(e.handlers, handler)
}

func (e *confirmServerConfig) Trigger(payload global.ServerConfig) {
	for _, handler := range e.handlers {
		go handler.HandleServerConfirmConfig(payload)
	}
}
