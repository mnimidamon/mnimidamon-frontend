package events

import (
	"mnimidamonbackend/frontend/global"
)

var ConfirmUserConfig confirmUserConfig

type ConfirmUserConfigHandler interface {
	HandleUserConfirmConfig(config global.UserConfig)
}

type confirmUserConfig struct {
	handlers []ConfirmUserConfigHandler
}

func (e *confirmUserConfig) Register(handler ConfirmUserConfigHandler) {
	e.handlers = append(e.handlers, handler)
}

func (e *confirmUserConfig) Trigger(config global.UserConfig) {
	for _, handler := range e.handlers {
		go handler.HandleUserConfirmConfig(config)
	}
}
