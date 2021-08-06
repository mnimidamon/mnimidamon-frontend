package viewmodels

import (
	"fyne.io/fyne/v2/dialog"
	"mnimidamonbackend/client/group"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	"mnimidamonbackend/frontend/views/server"
	"mnimidamonbackend/models"
)

var GroupMembers *groupMembersViewModel

func init() {
	GroupMembers = &groupMembersViewModel{
		Models: []*models.User{},
	}

	// Register selected group updates.
	events.SelectedGroupUpdated.Register(GroupMembers)
}

type groupMembersViewModel struct {
	Models []*models.User
}

func (vm *groupMembersViewModel) HandleSelectedGroupUpdated() {
	vm.GetAllMembers()
}

func (vm *groupMembersViewModel) GetAllMembers() {
	go func() {
		resp, err := server.Mnimidamon.Group.GetGroupMembers(&group.GetGroupMembersParams{
			GroupID:    SelectedGroup.Model.GroupID,
			Context:    server.ApiContext,
		}, CurrentComputer.Auth)

		if err != nil {
			dialog.ShowError(err, global.MainWindow)
			return
		}

		vm.Models = resp.Payload

		global.Log("members %v", vm.Models)
		vm.TriggerUpdateEvent()
	}()
}

func (vm *groupMembersViewModel) TriggerUpdateEvent() {
	events.GroupMembersUpdated.Trigger()
}
