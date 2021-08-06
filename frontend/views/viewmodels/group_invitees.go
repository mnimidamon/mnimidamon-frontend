package viewmodels

import (
	"fyne.io/fyne/v2/dialog"
	"mnimidamonbackend/client/group"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	"mnimidamonbackend/frontend/views/server"
	"mnimidamonbackend/models"
)

var GroupInvitees *groupInvitees

func init() {
	GroupInvitees = &groupInvitees{
		Models: []*models.Invite{},
	}

	// Register the group selection event, that's when we have to fetch that groups invitees.
	events.SelectedGroupUpdated.Register(GroupInvitees)
}

type groupInvitees struct {
	Models []*models.Invite
}

func (vm *groupInvitees) HandleSelectedGroupUpdated() {
	vm.GetAllInvitees()
}

func (vm *groupInvitees) GetAllInvitees() {
	go func() {
		resp, err := server.Mnimidamon.Group.GetGroupInvites(&group.GetGroupInvitesParams{
			GroupID: SelectedGroup.Model.GroupID,
			Context: server.ApiContext,
		}, CurrentComputer.Auth)

		if err != nil {
			dialog.ShowError(err, global.MainWindow)
			return
		}

		vm.Models = resp.Payload

		global.Log("Invitees %v", vm.Models)
		vm.TriggerUpdateEvent()
	}()
}

func (vm *groupInvitees) TriggerUpdateEvent() {
	events.GroupInviteesUpdated.Trigger()
}

func (vm *groupInvitees) Add(invite *models.Invite) {
	vm.Models = append(vm.Models, invite)
	vm.TriggerUpdateEvent()
}
