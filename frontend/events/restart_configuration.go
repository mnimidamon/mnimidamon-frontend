package events

var RestartConfiguration restartConfiguration

type RestartConfigurationHandler interface {
	HandleRestartConfigurationHandler()
}

type restartConfiguration struct {
	handlers []RestartConfigurationHandler
}

func (e *restartConfiguration) Register(handler RestartConfigurationHandler) {
	e.handlers = append(e.handlers, handler)
}

func (e *restartConfiguration) Trigger() {
	for _, handler := range e.handlers {
		go handler.HandleRestartConfigurationHandler()
	}
}


