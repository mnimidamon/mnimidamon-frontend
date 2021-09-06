package services

import (
	"fmt"
	"io/ioutil"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	"os"
	"path/filepath"
	"strconv"
)

var BackupStorage *backupStoreImpl

func init() {
	BackupStorage = &backupStoreImpl{}

	// Register the config confirm event.
	events.ConfirmConfig.Register(BackupStorage)

	// Register for config reset, that will delete all the local backups and the temp folder.
	events.RestartConfiguration.Register(BackupStorage)
}

type backupStoreImpl struct {
	BasePath string
	TempPath string
}

func (bs *backupStoreImpl) HandleRestartConfigurationHandler() {
	bs.DeleteAllBackupsAndTempFiles()
}

func (bs *backupStoreImpl) DeleteAllBackupsAndTempFiles() {
	global.Log("deleting all backups")
	for _, id := range bs.GetAllStoredBackupsIDS() {
		bs.DeleteBackup(int(id))
	}
}

func (bs *backupStoreImpl) HandleConfirmConfig(config global.Config) {
	global.Log("updating backup storage base path")
	bs.UpdateBasePath(config.Server.FolderPath)
}

func (bs *backupStoreImpl) UpdateBasePath(basePath string) {
	bs.BasePath = filepath.Join(basePath, "backups")
	bs.TempPath = filepath.Join(basePath, "temporary")
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

func (bs *backupStoreImpl) Create(backupID int) (*os.File, error) {
	return os.Create(bs.GetBackupPath(backupID))
}

func (bs *backupStoreImpl) CreateTemp(backupID int) (*os.File, error) {
	return os.Create(bs.GetTempPath(strconv.Itoa(backupID)))
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
	err := os.Remove(bs.GetTempPath(filename))
	if err != nil {
		global.Log("couldn't remove temp file %v", err)
	}
}

func (bs *backupStoreImpl) DeleteBackup(backupID int) error {
	global.Log("deleting backup file %v", backupID)
	err := os.Remove(bs.GetBackupPath(backupID))
	if err != nil {
		global.Log("couldn't remove backup file %v, %v", backupID, err)
	} else {
		global.Log("deleted backup file %v", backupID)
	}
	return err
}

func (bs *backupStoreImpl) GetAllStoredBackupsIDS() []int64 {
	files, err := ioutil.ReadDir(bs.BasePath)
	if err != nil {
		global.Log("error when opening backup storage directory: %v", err)
		return nil
	}

	var ids []int64
	for _, f := range files {
		id, err := strconv.Atoi(f.Name())
		if err == nil {
			ids = append(ids, int64(id))
		}
	}

	return ids
}

func (bs *backupStoreImpl) IsStored(backupID int) bool {
	if _, err := os.Stat(bs.GetBackupPath(backupID)); err == nil {
		return true
	}
	return false
}
func (bs *backupStoreImpl) MoveFromTemp(filename string, backupID int64) error {
	global.Log("moving file from temp %v", filename)

	oldLocation := bs.GetTempPath(filename)
	newLocation := bs.GetBackupPath(int(backupID))

	return os.Rename(oldLocation, newLocation)
}
