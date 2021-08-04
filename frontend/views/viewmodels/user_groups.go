package viewmodels

import (
	"fyne.io/fyne/v2/dialog"
	"mnimidamonbackend/client/current_user"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	"mnimidamonbackend/frontend/services"
	"mnimidamonbackend/frontend/views/server"
	"mnimidamonbackend/models"
)

var Groups *groupsViewModel

func init() {
	Groups = &groupsViewModel{
		Models: []*models.Group{},
	}

	// If the user is logged in.
	if services.ConfigurationStore.IsStored() {
		Groups.GetAllGroups()
	}

	// Register on confirm config.
	events.ConfirmConfig.Register(Groups)
	events.GroupCreated.Register(Groups)
}

type groupsViewModel struct {
	Models []*models.Group
}

func (vm *groupsViewModel) HandleGroupCreated(group models.Group) {
	vm.Models = append(vm.Models, &group)
	events.GroupsUpdated.Trigger()
}

func (vm *groupsViewModel) HandleConfirmConfig(_ global.Config) {
	vm.GetAllGroups()
}

func (vm *groupsViewModel) GetAllGroups() {
	go func() {
		resp, err := server.Mnimidamon.CurrentUser.GetCurrentUserGroups(&current_user.GetCurrentUserGroupsParams{
			Context:    server.ApiContext,
		}, server.CompAuth)

		if err != nil {
			dialog.ShowError(err, global.MainWindow)
			return
		}

		vm.Models = resp.Payload

		global.Log("%v", vm.Models)
		events.GroupsUpdated.Trigger()
	}()
}
