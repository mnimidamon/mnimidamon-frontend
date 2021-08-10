package viewmodels

import (
	"fyne.io/fyne/v2/dialog"
	"mnimidamonbackend/client/backup"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	"mnimidamonbackend/frontend/views/server"
	"mnimidamonbackend/models"
)

var Backups *backupsViewModel

func init() {
	Backups = &backupsViewModel {
		Models: []*models.Backup{},
	}

	events.SelectedGroupUpdated.Register(Backups)
}

type backupsViewModel struct {
	Models []*models.Backup
}

func (vm *backupsViewModel) HandleSelectedGroupUpdated() {
	vm.GetAllBackups()
}

func (vm *backupsViewModel) GetAllBackups() {
	go func() {
		resp, err := server.Mnimidamon.Backup.GetGroupBackups(&backup.GetGroupBackupsParams{
			GroupID:    SelectedGroup.Model.GroupID,
			Context:    server.ApiContext,
		}, CurrentComputer.Auth)

		if err != nil {
			dialog.ShowError(err, global.MainWindow)
			return
		}

		vm.Models = resp.Payload

		global.Log("backups %v", vm.Models)
		vm.TriggerUpdateEvent()
	}()
}

func (vm *backupsViewModel) Add(backup *models.Backup) {
	// Check if the group was not switched while the backup was loading itself into existence xD.
	if backup.GroupID != SelectedGroup.Model.GroupID {
		return
	}
	vm.Models = append(vm.Models, backup)
	vm.TriggerUpdateEvent()
}

func (vm *backupsViewModel) TriggerUpdateEvent() {
	events.BackupsUpdated.Trigger()
}


