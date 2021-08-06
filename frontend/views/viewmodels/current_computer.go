package viewmodels

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
)

var CurrentComputer *currentComputerViewModel

func init() {
	CurrentComputer = &currentComputerViewModel{
		Model: nil,
		Auth:  new(compAuthWriter),
	}

	events.ConfirmConfig.Register(CurrentComputer)
}
func (vm *currentComputerViewModel) HandleConfirmConfig(config global.Config) {
	// Update the user config first.
	CurrentUser.Select(config.User)
	// Update the computer config.
	vm.UpdateAuth(config)
	vm.Select(config.Computer)
}

type compAuthWriter struct{
	runtime.ClientAuthInfoWriter
}

type currentComputerViewModel struct {
	Model *global.ComputerConfig
	Auth  *compAuthWriter
}

func (vm *currentComputerViewModel) Select(computerConfig *global.ComputerConfig) {
	vm.Model = computerConfig
	global.Log("selected computer %v", computerConfig)
	vm.TriggerUpdateEvent()
}

func (vm *currentComputerViewModel) UpdateAuth(config global.Config) {
	vm.Auth.ClientAuthInfoWriter = runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		err := r.SetHeaderParam("X-COMP-KEY", config.Computer.Key)
		if err != nil {
			return err
		}
		return r.SetHeaderParam("X-AUTH-KEY", config.User.Key)
	})
}

func (vm *currentComputerViewModel) TriggerUpdateEvent() {
	events.CurrentComputerUpdated.Trigger()
}
