package taskmaker

import (
	"context"
	"mnimidamonbackend/client/backup"
	"mnimidamonbackend/frontend/services"
	"mnimidamonbackend/frontend/views/server"
	"mnimidamonbackend/frontend/views/viewmodels"
	"mnimidamonbackend/models"
)

type UploadTask struct {
	backup *models.Backup
}

func (task *UploadTask) Label() string {
	return "Uploading " + shortenName(task.backup.Filename)
}

func NewUploadTask(backup *models.Backup) *UploadTask {
	return &UploadTask{
		backup: backup,
	}
}

func (task *UploadTask) Execute(ctx context.Context, progress *uint) error {
	rc, err := services.BackupStorage.Get(int(task.backup.BackupID))

	if err != nil {
		return nil
	}

	// Send the backup on the server.
	resp, err := server.Mnimidamon.Backup.UploadBackup(&backup.UploadBackupParams{
		BackupData: rc,
		BackupID:   task.backup.BackupID,
		GroupID:    task.backup.GroupID,
		Context:    ctx,
	}, viewmodels.CurrentComputer.Auth)

	if err != nil {
		return err
	}

	viewmodels.Backups.AddOrUpdate(resp.Payload)
	return nil
}

// Get all upload tasks. Based on which are stored and which have a upload request flag on, make a task for uploading.
func GetAllUploadTasks(backups []*models.Backup) []QueuedTask {
	var tasks []QueuedTask

	for _, b := range backups {
		// Is stored and an upload request is present.
		if b.UploadRequest && services.BackupStorage.IsStored(int(b.BackupID)) {
			tasks = append(tasks, NewUploadTask(b))
		}
	}

	return tasks
}

func shortenName(name string) string {
	if len(name) > 12 {
		name = name[:12] + "..."
	}
	return name
}