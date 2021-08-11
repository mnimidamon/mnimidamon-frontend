package events

var ComputersUpdated computersUpdatedUpdated

type ComputersUpdatedHandler interface {
	HandleComputersUpdated()
}

type computersUpdatedUpdated struct {
	handlers []ComputersUpdatedHandler
}


func (e *computersUpdatedUpdated) Register(handler ComputersUpdatedHandler) {
	e.handlers = append(e.handlers, handler)
}

func (e *computersUpdatedUpdated) Trigger() {
	for _, handler := range e.handlers {
		go handler.HandleComputersUpdated()
	}
}

