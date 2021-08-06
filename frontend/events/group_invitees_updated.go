package events

var GroupInviteesUpdated groupInviteesUpdated

type GroupInviteesUpdatedHandler interface {
	HandleGroupInviteesUpdated()
}

type groupInviteesUpdated struct {
	handlers []GroupInviteesUpdatedHandler
}


func (e *groupInviteesUpdated) Register(handler GroupInviteesUpdatedHandler) {
	e.handlers = append(e.handlers, handler)
}

func (e *groupInviteesUpdated) Trigger() {
	for _, handler := range e.handlers {
		go handler.HandleGroupInviteesUpdated()
	}
}

