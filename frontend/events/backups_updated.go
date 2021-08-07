package events

var BackupsUpdated backupsUpdated

type BackupsUpdatedHandler interface {
	HandleBackupsUpdate()
}

type backupsUpdated struct {
	handlers []BackupsUpdatedHandler
}

func (e *backupsUpdated) Register(handler BackupsUpdatedHandler) {
	e.handlers = append(e.handlers, handler)
}

func (e *backupsUpdated) Trigger() {
	for _, handler := range e.handlers {
		go handler.HandleBackupsUpdate()
	}
}
