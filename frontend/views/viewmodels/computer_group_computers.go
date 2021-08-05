package viewmodels

import (
	"fyne.io/fyne/v2/dialog"
	"mnimidamonbackend/client/group_computer"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	"mnimidamonbackend/frontend/services"
	"mnimidamonbackend/frontend/views/server"
	"mnimidamonbackend/models"
)

var GroupComputers *groupComputersViewModel

func init() {
	GroupComputers = &groupComputersViewModel{
		Models: []*models.GroupComputer{},
	}
	
	// If the user is logged in.
	if services.ConfigurationStore.IsStored() {
		GroupComputers.GetAllGroupComputers()
	}
	
	// Register to fetch data upon authentication.
	events.Authenticated.Register(GroupComputers)
}

type groupComputersViewModel struct {
	Models []*models.GroupComputer
}

func (vm *groupComputersViewModel) IsMemberOf(group *models.Group) bool {
	for _, gc := range vm.Models {
		if gc.GroupID == group.GroupID {
			return true
		}
	}
	return false
}

func (vm *groupComputersViewModel) HandleAuthenticated() {
	vm.GetAllGroupComputers()
}

func (vm *groupComputersViewModel) GetAllGroupComputers()()  {
	go func() {
		resp, err := server.Mnimidamon.GroupComputer.GetGroupComputersOfComputer(&group_computer.GetGroupComputersOfComputerParams{
			ComputerID: services.ConfigurationStore.GetConfig().Computer.ComputerID,
			Context:    server.ApiContext,
		}, server.CompAuth)

		if err != nil {
			dialog.ShowError(err, global.MainWindow)
			return
		}

		vm.Models = resp.Payload
		global.Log("Computer Group Computers %v", vm.Models)
		events.GroupComputersUpdated.Trigger()
	}()
}
