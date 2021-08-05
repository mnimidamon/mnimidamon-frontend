package events


var InvitesUpdated invitesUpdated

type InvitesUpdatedHandler interface {
	HandleInvitesUpdate()
}

type invitesUpdated struct {
	handlers []InvitesUpdatedHandler
}


func (e *invitesUpdated) Register(handler InvitesUpdatedHandler) {
	e.handlers = append(e.handlers, handler)
}

func (e *invitesUpdated) Trigger() {
	for _, handler := range e.handlers {
		go handler.HandleInvitesUpdate()
	}
}

