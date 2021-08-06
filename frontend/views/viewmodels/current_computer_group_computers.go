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

var CurrentComputerGroupComputers *currentComputerGroupComputersViewModel

func init() {
	CurrentComputerGroupComputers = &currentComputerGroupComputersViewModel{
		Models: []*models.GroupComputer{},
	}
	
	// Register to fetch data upon computer change.
	events.CurrentComputerUpdated.Register(CurrentComputerGroupComputers)
}

type currentComputerGroupComputersViewModel struct {
	Models []*models.GroupComputer
}

func (vm *currentComputerGroupComputersViewModel) IsMemberOf(group *models.Group) bool {
	for _, gc := range vm.Models {
		if gc.GroupID == group.GroupID {
			return true
		}
	}
	return false
}

func (vm *currentComputerGroupComputersViewModel) Find(groupID int64, computerID int64) *models.GroupComputer {
	for _, gc := range vm.Models {
		if gc.GroupID == groupID && gc.ComputerID == computerID {
			return gc
		}
	}
	return nil
}

func (vm *currentComputerGroupComputersViewModel) Add(groupComputer *models.GroupComputer) {
	vm.Models = append(vm.Models, groupComputer)
	vm.TriggerUpdateEvent()
}

func (vm *currentComputerGroupComputersViewModel) HandleCurrentComputerUpdated() {
	vm.GetAllGroupComputers()
}

func (vm *currentComputerGroupComputersViewModel) GetAllGroupComputers()()  {
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

func (vm *currentComputerGroupComputersViewModel) TriggerUpdateEvent()()  {
	events.CurrentComputerGroupComputersUpdated.Trigger()
}
