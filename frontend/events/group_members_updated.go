package events

var GroupMembersUpdated groupMembersUpdated

type GroupMembersUpdatedHandler interface {
	HandleGroupMembersUpdated()
}

type groupMembersUpdated struct {
	handlers []GroupMembersUpdatedHandler
}


func (e *groupMembersUpdated) Register(handler GroupMembersUpdatedHandler) {
	e.handlers = append(e.handlers, handler)
}

func (e *groupMembersUpdated) Trigger() {
	for _, handler := range e.handlers {
		go handler.HandleGroupMembersUpdated()
	}
}

