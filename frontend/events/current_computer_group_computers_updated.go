package events

var CurrentComputerGroupComputersUpdated currentComputerGroupComputersUpdated

type CurrentComputerGroupComputersUpdatedHandler interface {
	HandleGroupComputersUpdated()
}
type currentComputerGroupComputersUpdated struct {
	handlers []CurrentComputerGroupComputersUpdatedHandler
}

func (e *currentComputerGroupComputersUpdated) Register(handler CurrentComputerGroupComputersUpdatedHandler) {
	e.handlers = append(e.handlers, handler)
}

func (e *currentComputerGroupComputersUpdated) Trigger() {
	for _, handler := range e.handlers {
		go handler.HandleGroupComputersUpdated()
	}
}
