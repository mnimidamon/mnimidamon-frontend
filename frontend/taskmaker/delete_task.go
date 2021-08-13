package taskmaker

import (
	"context"
	"mnimidamonbackend/frontend/services"
	"mnimidamonbackend/frontend/views/viewmodels"
	"mnimidamonbackend/models"
	"strconv"
)

type DeleteTask struct {
	backupID int64
	progress int
}

func (task *DeleteTask) Execute(ctx context.Context) error {
	task.progress = 0

	var err error
	if services.BackupStorage.IsStored(int(task.backupID)) {
		err = services.BackupStorage.DeleteBackup(int(task.backupID))
	}

	for _, b := range viewmodels.Backups.Models {
		if b.BackupID == task.backupID {
			viewmodels.Backups.Remove(b)
		}
	}

	task.progress = 100
	return err
}

func (task *DeleteTask) GetProgress() int {
	return task.progress
}

func NewDeleteTask(backupID int64) QueuedTask {
	return &DeleteTask{
		backupID: backupID,
	}
}

func (task *DeleteTask) Label() string {
	return "Deleting backup with id " + strconv.FormatInt(task.backupID, 10)
}

func GetAllDeletionTasks(backups []*models.Backup) []QueuedTask {
	storedBackupIDS := services.BackupStorage.GetAllStoredBackupsIDS()

	var deletionTasks []QueuedTask

	for _, sID := range storedBackupIDS {
		shouldBeDeleted := true
		for _, b := range backups {
			if sID == b.BackupID {
				shouldBeDeleted = false
				break
			}
		}

		if shouldBeDeleted {
			deletionTasks = append(deletionTasks, NewDeleteTask(sID))
		}
	}

	return deletionTasks
}
