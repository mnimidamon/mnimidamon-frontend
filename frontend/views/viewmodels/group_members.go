package viewmodels

import (
	"fyne.io/fyne/v2/dialog"
	"mnimidamonbackend/client/computer"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	"mnimidamonbackend/frontend/views/server"
	"mnimidamonbackend/models"
)

var GroupComputers *groupComputersViewModel

func init() {
	GroupComputers = &groupComputersViewModel {
		Models: []*models.GroupComputer{},
	}

	events.SelectedGroupUpdated.Register(GroupComputers)
}

type groupComputersViewModel struct {
	Models []*models.GroupComputer
}

func (vm *groupComputersViewModel) HandleSelectedGroupUpdated() {
	vm.GetAllGroupComputers()
}

func (vm *groupComputersViewModel) GetAllGroupComputers()  {
	go func() {
		resp, err := server.Mnimidamon.Computer.GetCurrentUserGroupComputers(&computer.GetCurrentUserGroupComputersParams{
			GroupID:    SelectedGroup.Model.GroupID,
			Context:    server.ApiContext,
		}, CurrentComputer.Auth)

		if err != nil {
			dialog.ShowError(err, global.MainWindow)
			return
		}

		vm.Models = resp.Payload

		global.Log("selected group group computers %v", vm.Models)
		vm.TriggerUpdateEvent()
	}()
}

func (vm *groupComputersViewModel) TriggerUpdateEvent() {
	events.GroupComputersUpdated.Trigger()
}