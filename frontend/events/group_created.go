package events

import "mnimidamonbackend/models"

var GroupCreated groupCreated

type GroupCreatedHandler interface {
	HandleGroupCreated(group models.Group)
}

type groupCreated struct {
	handlers []GroupCreatedHandler
}

func (e *groupCreated) Register(handler GroupCreatedHandler) {
	e.handlers = append(e.handlers, handler)
}

func (e *groupCreated) Trigger(group models.Group) {
	for _, handler := range e.handlers {
		go handler.HandleGroupCreated(group)
	}
}
