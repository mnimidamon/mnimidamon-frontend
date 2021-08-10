package services

import (
	"fmt"
	"io"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	"mnimidamonbackend/models"
	"os"
	"path/filepath"
	"strconv"
)

var BackupStorage *backupStoreImpl

func init() {
	BackupStorage = &backupStoreImpl{}

	// Register the config confirm event.
	events.ConfirmConfig.Register(BackupStorage)
}

type backupStoreImpl struct {
	BasePath string
	TempPath string
}

func (bs *backupStoreImpl) HandleConfirmConfig(config global.Config) {
	global.Log("updating backup storage base path")
	bs.UpdateBasePath(config.Server.FolderPath)
}

func (bs *backupStoreImpl) UpdateBasePath(basePath string) {
	bs.BasePath = basePath
	bs.TempPath = filepath.Join(basePath, "temp")
	bs.makeRequiredFolders()
}

func (bs *backupStoreImpl) GetBackupPath(backupID int) string {
	return filepath.Join(bs.BasePath, strconv.Itoa(backupID))
}

func (bs *backupStoreImpl) GetTempPath(filename string) string {
	return filepath.Join(bs.TempPath, filename)
}

func (bs *backupStoreImpl) Get(backupID int) (*os.File, error) {
	return os.Open(bs.GetBackupPath(backupID))
}

func (bs *backupStoreImpl) Save(backup models.Backup, readCloser io.ReadCloser) error {
	// TODO poglej kak je na stežniku
	panic("implement me")
}

func (bs *backupStoreImpl) makeRequiredFolders() error {
	// Make file store folder
	if _, err := os.Stat(bs.BasePath); os.IsNotExist(err) {
		err := os.MkdirAll(bs.BasePath, 0700)
		if err != nil {
			return fmt.Errorf("could not create base backup store folder: %w", err)
		}
	}

	if _, err := os.Stat(bs.TempPath); os.IsNotExist(err) {
		err := os.MkdirAll(bs.TempPath, 0700)
		if err != nil {
			return fmt.Errorf("could not create base temp backup store folder: %w", err)
		}
	}

	return nil
}

func (bs *backupStoreImpl) DeleteTempFile(filename string) {
	global.Log("deleting file from temp %v", filename)
	os.Remove(bs.GetTempPath(filename))
}

func (bs *backupStoreImpl) MoveFromTemp(filename string, backupID int64) error {
	global.Log("moving file from temp %v", filename)

	oldLocation := bs.GetTempPath(filename)
	newLocation := bs.GetBackupPath(int(backupID))

	return os.Rename(oldLocation, newLocation)
}
