package events

import (
	"mnimidamonbackend/models"
)

var LoginUser loginUser

type LoginUserHandler interface {
	HandleUserLogin(models.User)
}

type loginUser struct {
	handlers []LoginUserHandler
}

func (e *loginUser) Register(handler LoginUserHandler) {
	e.handlers = append(e.handlers, handler)
}

func (e *loginUser) Trigger(user models.User) {
	for _, handler := range e.handlers {
		go handler.HandleUserLogin(user)
	}
}
