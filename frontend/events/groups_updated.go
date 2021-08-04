package events

var GroupsUpdated groupsUpdated

type GroupsUpdatedHandler interface {
	HandleGroupsUpdate()
}

type groupsUpdated struct {
	handlers []GroupsUpdatedHandler
}


func (e *groupsUpdated) Register(handler GroupsUpdatedHandler) {
	e.handlers = append(e.handlers, handler)
}

func (e *groupsUpdated) Trigger() {
	for _, handler := range e.handlers {
		go handler.HandleGroupsUpdate()
	}
}
