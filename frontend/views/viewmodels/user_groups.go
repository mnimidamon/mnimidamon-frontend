package viewmodels

import (
	"fyne.io/fyne/v2/dialog"
	"mnimidamonbackend/client/current_user"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	"mnimidamonbackend/frontend/views/server"
	"mnimidamonbackend/models"
)

var Groups *groupsViewModel

func init() {
	Groups = &groupsViewModel{
		Models: []*models.Group{},
	}

	// Register current computer updates.
	events.CurrentComputerUpdated.Register(Groups)
}

type groupsViewModel struct {
	Models []*models.Group
}

func (vm *groupsViewModel) HandleCurrentComputerUpdated() {
	vm.GetAllGroups()
}

func (vm *groupsViewModel) Add(group *models.Group) {
	vm.Models = append(vm.Models, group)
	vm.TriggerUpdateEvent()
}

func (vm *groupsViewModel) GetAllGroups() {
	go func() {
		resp, err := server.Mnimidamon.CurrentUser.GetCurrentUserGroups(&current_user.GetCurrentUserGroupsParams{
			Context: server.ApiContext,
		}, CurrentComputer.Auth)

		if err != nil {
			dialog.ShowError(err, global.MainWindow)
			return
		}

		vm.Models = resp.Payload

		global.Log("Groups %v", vm.Models)
		vm.TriggerUpdateEvent()
	}()
}

func (vm *groupsViewModel) TriggerUpdateEvent() {
	events.GroupsUpdated.Trigger()
}
