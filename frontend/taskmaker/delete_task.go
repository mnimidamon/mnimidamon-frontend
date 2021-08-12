package taskmaker

import (
	"context"
	"mnimidamonbackend/frontend/services"
	"mnimidamonbackend/models"
)

type DeleteTask struct {
	backupID int64
}

func NewDeleteTask(backupID int64) QueuedTask {
	return &DeleteTask{
		backupID: backupID,
	}
}

func (task *DeleteTask) Label() string {
	return "Deleting backup with id " + string(task.backupID)
}

func (task *DeleteTask) Execute(ctx context.Context, progress *uint) error {
	err := services.BackupStorage.DeleteBackup(int(task.backupID))
	return err
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
