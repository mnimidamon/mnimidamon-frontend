package services

import (
	"io"
	"mnimidamonbackend/models"
)

// File Backup Store
type BackupStore interface {
	BackupGetter
	BackupSaver
}

// Encryption and decryption.
type BackupEncryption interface {
	BackupEncoder
	BackupDecoder
}

// Encoder
type BackupDecoder interface {
	Decode(backup *models.Backup, readCloser io.ReadCloser) (io.ReadCloser, error)
}
// Decoder
type BackupEncoder interface {
	Encode(backup *models.Backup, readCloser io.ReadCloser) (io.ReadCloser, error)
}

// Reader
type BackupGetter interface {
	Get(backupID int64) (io.ReadCloser, error)
}
// Writer
type BackupSaver interface {
	Save(backup models.Backup, readCloser io.ReadCloser) error
}