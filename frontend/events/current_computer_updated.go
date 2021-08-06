package events

var CurrentComputerUpdated currentComputerUpdated

type CurrentComputerUpdatedHandler interface {
	HandleCurrentComputerUpdated()
}
type currentComputerUpdated struct {
	handlers []CurrentComputerUpdatedHandler
}

func (e *currentComputerUpdated) Register(handler CurrentComputerUpdatedHandler) {
	e.handlers = append(e.handlers, handler)
}

func (e *currentComputerUpdated) Trigger() {
	for _, handler := range e.handlers {
		go handler.HandleCurrentComputerUpdated()
	}
}
