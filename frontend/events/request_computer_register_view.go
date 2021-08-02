package events

var RequestComputerRegisterView requestComputerRegisterView

type RequestComputerRegisterViewHandler interface {
	HandleRequestComputerRegisterView()
}

type requestComputerRegisterView struct {
	handlers []RequestComputerRegisterViewHandler
}

func (e *requestComputerRegisterView) Register(handler RequestComputerRegisterViewHandler) {
	e.handlers = append(e.handlers, handler)
}

func (e *requestComputerRegisterView) Trigger() {
	for _, handler := range e.handlers {
		go handler.HandleRequestComputerRegisterView()
	}
}
