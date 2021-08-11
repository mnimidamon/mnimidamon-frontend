package viewmodels

import (
	"fyne.io/fyne/v2/dialog"
	"mnimidamonbackend/client/computer"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	"mnimidamonbackend/frontend/views/server"
	"mnimidamonbackend/models"
)

var Computers *computersViewModel

func init() {
	Computers = &computersViewModel {
		Models: []*models.Computer{},
	}

	// Register on selected user.
	events.CurrentUserUpdated.Register(Computers)
}

type computersViewModel struct {
	Models []*models.Computer
}

func (vm *computersViewModel) HandleCurrentUserUpdated() {
	vm.GetAll()
}

func (vm *computersViewModel) GetAll() {
	go func() {
		resp, err := server.Mnimidamon.Computer.GetCurrentUserComputers(&computer.GetCurrentUserComputersParams{
			Context:    server.ApiContext,
		}, CurrentUser.Auth)

		if err != nil {
			dialog.ShowError(err, global.MainWindow)
			return
		}

		vm.Models = resp.Payload

		global.Log("computers %v", vm.Models)
		vm.TriggerUpdateEvent()
	}()
}

func (vm *computersViewModel) TriggerUpdateEvent() {
	events.ComputersUpdated.Trigger()
}
