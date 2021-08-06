package events

var SelectedGroupUpdated selectedGroupUpdated

type SelectedGroupUpdatedHandler interface {
	HandleSelectedGroupUpdated()
}
type selectedGroupUpdated struct {
	handlers []SelectedGroupUpdatedHandler
}

func (e *selectedGroupUpdated) Register(handler SelectedGroupUpdatedHandler) {
	e.handlers = append(e.handlers, handler)
}

func (e *selectedGroupUpdated) Trigger() {
	for _, handler := range e.handlers {
		go handler.HandleSelectedGroupUpdated()
	}
}