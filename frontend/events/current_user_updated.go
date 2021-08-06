package events

var CurrentUserUpdated currentUserUpdated

type CurrentUserUpdatedHandler interface {
	HandleCurrentUserUpdated()
}
type currentUserUpdated struct {
	handlers []CurrentUserUpdatedHandler
}

func (e *currentUserUpdated) Register(handler CurrentUserUpdatedHandler) {
	e.handlers = append(e.handlers, handler)
}

func (e *currentUserUpdated) Trigger() {
	for _, handler := range e.handlers {
		go handler.HandleCurrentUserUpdated()
	}
}

