package events


var RequestMainView requestMainView

type RequestMainViewHandler interface {
	HandleRequestMainView()
}

type requestMainView struct {
	handlers []RequestMainViewHandler
}

func (e *requestMainView) Register(handler RequestMainViewHandler) {
	e.handlers = append(e.handlers, handler)
}

func (e *requestMainView) Trigger() {
	for _, handler := range e.handlers {
		go handler.HandleRequestMainView()
	}
}

