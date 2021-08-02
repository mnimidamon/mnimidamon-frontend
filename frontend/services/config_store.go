package services

import (
	"fyne.io/fyne/v2"
	"mnimidamonbackend/frontend/global"
	"mnimidamonbackend/models"
)

// Preference keys.
var (
	isStoredField = "IS_STORED"

	serverHostField = "SERVER_HOST"
	serverPortField = "SERVER_PORT"
	folderPathKey   = "FOLDER_PATH"

	userId        = "USER_ID"
	usernameField = "USERNAME"
	userKeyField  = "USER_KEY"

	computerId        = "COMPUTER_ID"
	computerNameField = "COMPUTER_NAME"
	computerKeyField  = "COMPUTER_KEY"
)

// Global Configuration Store for accessing, deleting and saving our configurations.
var ConfigurationStore configurationStore

func init() {
	// Get the application preferences and save it.
	ConfigurationStore = configurationStore{
		preferences: global.App.Preferences(),
	}
}

type configurationStore struct {
	preferences fyne.Preferences
}

func (i *configurationStore) Delete() {
	i.preferences.SetBool(isStoredField, false)
}

func (i *configurationStore) IsStored() bool {
	return i.preferences.Bool(isStoredField)
}

func (i *configurationStore) SaveConfig(c *global.Config) {
	i.preferences.SetBool(isStoredField, true)

	i.preferences.SetString(userKeyField, c.User.Username)
	i.preferences.SetString(userKeyField, c.User.Key)
	i.preferences.SetString(computerKeyField, c.Computer.Key)
	i.preferences.SetString(computerNameField, c.Computer.Name)
	i.preferences.SetString(serverHostField, c.Server.Host)
	i.preferences.SetInt(serverPortField, c.Server.Port)
	i.preferences.SetString(folderPathKey, c.Server.FolderPath)
	i.preferences.SetInt(userId, int(c.User.UserID))
	i.preferences.SetInt(computerId, int(c.Computer.ComputerID))
}

func (i *configurationStore) GetConfig() *global.Config {
	if !i.IsStored() {
		return nil
	}

	return &global.Config{
		User: &global.UserConfig{
			User: models.User{
				UserID:   int64(i.preferences.Int(userId)),
				Username: i.preferences.String(usernameField),
			},
			Key: i.preferences.String(userKeyField),
		},
		Computer: &global.ComputerConfig{
			Computer: models.Computer{
				ComputerID: int64(i.preferences.Int(computerId)),
				Name:       i.preferences.String(computerNameField),
				OwnerID:    int64(i.preferences.Int(userId)),
			},
			Key: i.preferences.String(computerKeyField),
		},
		Server: &global.ServerConfig{
			Host:       i.preferences.String(serverHostField),
			Port:       i.preferences.Int(serverPortField),
			FolderPath: i.preferences.String(folderPathKey),
		},
	}
}
