package events

import (
	"mnimidamonbackend/frontend/global"
)

func init() {
	constructedNotifier := &configConstructedNotifier{}

	ConfirmServerConfig.Register(constructedNotifier)
	RestartConfiguration.Register(constructedNotifier)
	ConfirmUserConfig.Register(constructedNotifier)
	ConfirmComputerConfig.Register(constructedNotifier)
}

// Listen for when the config is fully constructed and notify the listeners.
type configConstructedNotifier struct {
	Config global.Config
}

func (i *configConstructedNotifier) HandleComputerConfirmConfig(payload global.ComputerConfig) {
	global.Log("computer configuration")
	i.Config.Computer = &payload
	i.DistributeConfigDoneEventIfCompleted()
}

func (i *configConstructedNotifier) HandleRestartConfigurationHandler() {
	global.Log("config reset")
	i.Config.User = nil
	i.Config.Computer = nil
	i.Config.Server = nil
}

func (i *configConstructedNotifier) HandleServerConfirmConfig(payload global.ServerConfig) {
	global.Log("server configuration")
	i.Config.Server = &payload
	i.DistributeConfigDoneEventIfCompleted()
}

func (i *configConstructedNotifier) HandleUserConfirmConfig(payload global.UserConfig) {
	global.Log("user configuration")
	i.Config.User = &payload
	i.DistributeConfigDoneEventIfCompleted()
}

func (i *configConstructedNotifier) DistributeConfigDoneEventIfCompleted() {
	if i.Config.Server != nil && i.Config.Computer != nil && i.Config.User != nil {
		global.Log("Configuration constructed")
		ConfirmConfig.Trigger(i.Config)
	}
}
