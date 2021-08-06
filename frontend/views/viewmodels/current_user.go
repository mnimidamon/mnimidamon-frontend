package viewmodels

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
)

var CurrentUser *currentUserViewModel

func init() {
	CurrentUser = &currentUserViewModel{
		Model: nil,
		Auth:  new(userAuthWriter),
	}

	events.ConfirmConfig.Register(CurrentUser)
	events.ConfirmUserConfig.Register(CurrentUser)
}

type userAuthWriter struct {
	runtime.ClientAuthInfoWriter
}

type currentUserViewModel struct {
	Model *global.UserConfig
	Auth  *userAuthWriter
}

func (vm *currentUserViewModel) Select(userConfig *global.UserConfig) {
	vm.Model = userConfig
	global.Log("selected user %v", userConfig)
	vm.UpdateAuth(*userConfig)
	vm.TriggerUpdateEvent()
}

func (vm *currentUserViewModel) TriggerUpdateEvent() {
	events.CurrentUserUpdated.Trigger()
}

func (vm *currentUserViewModel) HandleConfirmConfig(config global.Config) {
	vm.UpdateAuth(*config.User)
}

func (vm *currentUserViewModel) HandleUserConfirmConfig(config global.UserConfig) {
	vm.UpdateAuth(config)
}

func (vm *currentUserViewModel) UpdateAuth(config global.UserConfig) {
	vm.Auth.ClientAuthInfoWriter = httptransport.APIKeyAuth("X-AUTH-KEY", "header", config.Key)
}