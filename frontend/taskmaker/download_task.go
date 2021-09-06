package taskmaker

import (
	"context"
	"io"
	"mnimidamonbackend/client/backup"
	"mnimidamonbackend/frontend/global"
	"mnimidamonbackend/frontend/services"
	"mnimidamonbackend/frontend/views/server"
	"mnimidamonbackend/frontend/views/viewmodels"
	"mnimidamonbackend/models"
	"os"
	"strconv"
)

type DownloadTask struct {
	backup *models.Backup
	file *os.File
}

func NewDownloadTask(backup *models.Backup) *DownloadTask {
	return &DownloadTask{
		backup: backup,
	}
}

func (task *DownloadTask) Label() string {
	return "Downloading " + shortenName(task.backup.Filename)
}

func (task *DownloadTask) Execute(ctx context.Context) error {
	file, err := services.BackupStorage.CreateTemp(int(task.backup.BackupID))

	if err != nil {
		return err
	}
	defer services.BackupStorage.DeleteTempFile(strconv.Itoa(int(task.backup.BackupID)))
	defer file.Close()

	task.file = file
	_, err = server.Mnimidamon.Backup.DownloadBackup(&backup.DownloadBackupParams{
		BackupID:   task.backup.BackupID,
		GroupID:    task.backup.GroupID,
		Context:    ctx,
	}, viewmodels.CurrentComputer.Auth, file)

	if err != nil {
		return err
	}


	_, err = file.Seek(0, 0)
	if err != nil {
		return nil
	}


	prefix := task.backup.Filename
	preader := NewPrefixReaderCloser(file, []byte(prefix))
	hash, err := services.CalculateReaderHash(preader)
	if err != nil {
		return err
	}
	file.Close()

	global.Log("hash %v prefix %v", hash, prefix)
	_, err = server.Mnimidamon.Backup.LogComputerBackup(&backup.LogComputerBackupParams{
		BackupID:   task.backup.BackupID,
		Body:       &models.ConfirmDownloadPayload{
			Hash:          &hash,
			PrependString: &prefix,
		},
		GroupID:    task.backup.GroupID,
		Context:    ctx,
	}, viewmodels.CurrentComputer.Auth)

	if err != nil {
		return err
	}

	services.BackupStorage.MoveFromTemp(strconv.Itoa(int(task.backup.BackupID)), task.backup.BackupID)
	viewmodels.Backups.TriggerUpdateEvent()
	return nil
}

func (task *DownloadTask) GetProgress() int {
	offset, err := task.file.Seek(0, io.SeekCurrent)
	percentage := int(float64(offset)/float64(task.backup.Size * 1024) * 100)
	if err != nil {
		return 100
	}

	return percentage
}

func GetAllDownloadTasks(backups []*models.Backup) []QueuedTask {
	var tasks []QueuedTask

	for _, b := range backups {
		// Is stored and an upload request is present.
		if b.OnServer && !services.BackupStorage.IsStored(int(b.BackupID)) {
			tasks = append(tasks, NewDownloadTask(b))
		}
	}

	return tasks
}


func NewPrefixReaderCloser(rc io.ReadCloser, prefix []byte) io.ReadCloser{
	return &prefixedReaderCloser{
		RC:     rc,
		Prefix: prefix,
		i:      0,
	}
}


type prefixedReaderCloser struct {
	RC     io.ReadCloser
	Prefix []byte
	i      int
}

func (prc *prefixedReaderCloser) Close() error {
	return prc.RC.Close()
}

func (prc *prefixedReaderCloser) Read(p []byte) (n int, err error) {
	toRead := len(p)

	// Prefix has already been read.
	if len(prc.Prefix) < prc.i + 1 {
		return prc.RC.Read(p)
	}

	// Copy prefix to byte
	n = copy(p, prc.Prefix[prc.i:])
	prc.i += n
	if n < toRead {
		x, err := prc.RC.Read(p[n:])

		if err != nil {
			return x + n, err
		}

		prc.i += x + n
		return x + n, nil
	}

	prc.i += n
	return n, nil
}
