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

var Invites *userInvitesViewModel


func init() {
	Invites = &userInvitesViewModel{
		Models: []*models.Invite{},
	}

	// If the user is logged in.
	if services.ConfigurationStore.IsStored() {
		Invites.GetAllInvites()
	}

	// Register on confirm config.
	events.Authenticated.Register(Invites)
}

type userInvitesViewModel struct {
	Models []*models.Invite
}

func (vm *userInvitesViewModel) HandleAuthenticated() {
	vm.GetAllInvites()
}

func (vm *userInvitesViewModel) Remove(i *models.Invite) {
	for j, x := range vm.Models {
		if x.Group.GroupID == i.Group.GroupID {
			vm.Models = append(vm.Models[:j], vm.Models[j+1:]...)
		}
	}
	
	vm.TriggerUpdateEvent()
}

func (vm *userInvitesViewModel) GetAllInvites()  {
	go func() {
		resp, err := server.Mnimidamon.CurrentUser.GetCurrentUserInvites(&current_user.GetCurrentUserInvitesParams{
			Context:    server.ApiContext,
		}, server.CompAuth)

		if err != nil {
			dialog.ShowError(err, global.MainWindow)
			return
		}

		vm.Models = resp.Payload

		global.Log("Invites %v", vm.Models)
		vm.TriggerUpdateEvent()
	}()
}

func (vm *userInvitesViewModel) TriggerUpdateEvent() {
	events.InvitesUpdated.Trigger()
}