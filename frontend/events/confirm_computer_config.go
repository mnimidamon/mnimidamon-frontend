package events

import "mnimidamonbackend/frontend/global"

var ConfirmComputerConfig confirmComputerConfig

type ConfirmComputerConfigHandler interface {
	HandleComputerConfirmConfig(config global.ComputerConfig)
}

type confirmComputerConfig struct {
	handlers []ConfirmComputerConfigHandler
}

func (e *confirmComputerConfig) Register(handler ConfirmComputerConfigHandler) {
	e.handlers = append(e.handlers, handler)
}

func (e *confirmComputerConfig) Trigger(config global.ComputerConfig) {
	for _, handler := range e.handlers {
		go handler.HandleComputerConfirmConfig(config)
	}
}


