package events

// Event describes the selection of server configurations.
var ConfirmServerConfig confirmServerConfig

type ConfirmServerConfigPayload struct {
	FolderPath string
	Host       string
	Port       int
}

type ConfirmServerConfigHandler interface {
	HandleServerConfirmConfig(payload ConfirmServerConfigPayload)
}

type confirmServerConfig struct {
	handlers []ConfirmServerConfigHandler
}

func (e *confirmServerConfig) Register(handler ConfirmServerConfigHandler) {
	e.handlers = append(e.handlers, handler)
}

func (e *confirmServerConfig) Trigger(payload ConfirmServerConfigPayload) {
	for _, handler := range e.handlers {
		go handler.HandleServerConfirmConfig(payload)
	}
}
