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
	
	// Register to fetch data upon computer change.
	events.CurrentComputerUpdated.Register(GroupComputers)
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

func (vm *groupComputersViewModel) Find(groupID int64, computerID int64) *models.GroupComputer {
	for _, gc := range vm.Models {
		if gc.GroupID == groupID && gc.ComputerID == computerID {
			return gc
		}
	}
	return nil
}

func (vm *groupComputersViewModel) Add(groupComputer *models.GroupComputer) {
	vm.Models = append(vm.Models, groupComputer)
	vm.TriggerUpdateEvent()
}

func (vm *groupComputersViewModel) HandleCurrentComputerUpdated() {
	vm.GetAllGroupComputers()
}

func (vm *groupComputersViewModel) GetAllGroupComputers()()  {
	go func() {
		resp, err := server.Mnimidamon.GroupComputer.GetGroupComputersOfComputer(&group_computer.GetGroupComputersOfComputerParams{
			ComputerID: services.ConfigurationStore.GetConfig().Computer.ComputerID,
			Context:    server.ApiContext,
		}, CurrentComputer.Auth)

		if err != nil {
			dialog.ShowError(err, global.MainWindow)
			return
		}

		vm.Models = resp.Payload
		global.Log("Computer Group Computers %v", vm.Models)
		vm.TriggerUpdateEvent()
	}()
}

func (vm *groupComputersViewModel) TriggerUpdateEvent()()  {
	events.GroupComputersUpdated.Trigger()
}
