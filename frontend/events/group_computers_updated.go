package events

var GroupComputersUpdated groupComputersUpdated

type GroupComputersUpdatedHandler interface {
	HandleGroupComputersUpdated()
}

type groupComputersUpdated struct {
	handlers []GroupComputersUpdatedHandler
}


func (e *groupComputersUpdated) Register(handler GroupComputersUpdatedHandler) {
	e.handlers = append(e.handlers, handler)
}

func (e *groupComputersUpdated) Trigger() {
	for _, handler := range e.handlers {
		go handler.HandleGroupComputersUpdated()
	}
}

