package events

var Authenticated authenticated

type AuthenticatedHandler interface {
	HandleAuthenticated()
}

type authenticated struct {
	handlers []AuthenticatedHandler
}

func (e *authenticated) Register(handler AuthenticatedHandler) {
	e.handlers = append(e.handlers, handler)
}

func (e *authenticated) Trigger() {
	for _, handler := range e.handlers {
		go handler.HandleAuthenticated()
	}
}

