package events

var RequestRegisterView requestRegisterView

type RequestRegisterViewHandler interface {
	HandleRequestRegisterView()
}

type requestRegisterView struct {
	handlers []RequestRegisterViewHandler
}

func (e *requestRegisterView) Register(handler RequestRegisterViewHandler) {
	e.handlers = append(e.handlers, handler)
}

func (e *requestRegisterView) Trigger() {
	for _, handler := range e.handlers {
		go handler.HandleRequestRegisterView()
	}
}
