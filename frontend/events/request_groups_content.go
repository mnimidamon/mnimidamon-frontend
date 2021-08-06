package events

var RequestGroupsContent requestGroupsContent

type RequestGroupsContentHandler interface {
	HandleRequestGroupsContent()
}
type requestGroupsContent struct {
	handlers []RequestGroupsContentHandler
}

func (e *requestGroupsContent) Register(handler RequestGroupsContentHandler) {
	e.handlers = append(e.handlers, handler)
}

func (e *requestGroupsContent) Trigger() {
	for _, handler := range e.handlers {
		go handler.HandleRequestGroupsContent()
	}
}