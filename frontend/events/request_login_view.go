package events

var RequestLoginView requestLoginView

type RequestLoginViewHandler interface {
	HandleRequestLoginView()
}

type requestLoginView struct {
	handlers []RequestLoginViewHandler
}

func (e *requestLoginView) Register(handler RequestLoginViewHandler) {
	e.handlers = append(e.handlers, handler)
}

func (e *requestLoginView) Trigger() {
	for _, handler := range e.handlers {
		go handler.HandleRequestLoginView()
	}
}
