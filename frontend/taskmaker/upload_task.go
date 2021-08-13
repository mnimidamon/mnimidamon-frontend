package taskmaker

import (
	"context"
	"io"
	"mnimidamonbackend/client/backup"
	"mnimidamonbackend/frontend/services"
	"mnimidamonbackend/frontend/views/server"
	"mnimidamonbackend/frontend/views/viewmodels"
	"mnimidamonbackend/models"
	"os"
)

type UploadTask struct {
	backup *models.Backup
	File   *os.File
	Size   int64
}

func (task *UploadTask) Label() string {
	return "Uploading " + shortenName(task.backup.Filename)
}

func NewUploadTask(backup *models.Backup) *UploadTask {
	return &UploadTask{
		backup:    backup,
		File:      nil,
		Size:      0,
	}
}

func (task *UploadTask) GetProgress() int {
	if task.File == nil {
		return 0
	}

	offset, err := task.File.Seek(0, io.SeekCurrent)
	percentage := int(float64(offset)/float64(task.Size) * 100)
	if err != nil {
		return 100
	}

	return percentage
}

func (task *UploadTask) Execute(ctx context.Context) error {
	file, err := services.BackupStorage.Get(int(task.backup.BackupID))

	if err != nil {
		return nil
	}
	defer file.Close()

	fi, _ := file.Stat()
	task.Size = fi.Size()
	task.File = file


	// Send the backup on the server.
	resp, err := server.Mnimidamon.Backup.UploadBackup(&backup.UploadBackupParams{
		BackupData: task.File,
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
