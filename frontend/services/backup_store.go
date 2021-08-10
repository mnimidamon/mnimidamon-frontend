package services

import (
	"io"
	"mnimidamonbackend/models"
	"os"
)

// File Backup Store
type BackupStore interface {
	UpdateBasePath(basePath string)
	GetBackupPath(backupID int) string
	GetTempPath(filename string) string
	BackupGetter
	BackupSaver
}

// Reader
type BackupGetter interface {
	Get(backupID int) (*os.File, error)
}

// Writer
type BackupSaver interface {
	Save(backup models.Backup, readCloser io.ReadCloser) error
}