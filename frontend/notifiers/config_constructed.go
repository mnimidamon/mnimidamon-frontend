package notifiers

import (
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	"net"
)

func init() {
	constructedNotifier := &configConstructedNotifier{}

	events.ConfirmServerConfig.Register(constructedNotifier)
	events.RestartConfiguration.Register(constructedNotifier)
}

// Listen for when the config is fully constructed and notify the listeners.
type configConstructedNotifier struct {
	Config global.Config
}

func (i *configConstructedNotifier) HandleRestartConfigurationHandler() {
	i.Config.User = nil
	i.Config.Computer = nil
	i.Config.Server = nil
}

func (i *configConstructedNotifier) HandleServerConfirmConfig(payload events.ConfirmServerConfigPayload) {
	i.Config.Server = &global.ServerConfig{
		ServerAddress: net.JoinHostPort(payload.Host, string(rune(payload.Port))),
		FolderPath:    payload.FolderPath,
	}
	i.DistributeConfigDoneEventIfCompleted()
}

func (i *configConstructedNotifier) HandleUserConfirmConfig(payload global.UserConfig) {
	i.Config.User = &payload
	i.DistributeConfigDoneEventIfCompleted()
}

func (i *configConstructedNotifier) DistributeConfigDoneEventIfCompleted() {
	if i.Config.Server != nil && i.Config.Computer != nil && i.Config.User != nil {
		events.ConfirmConfig.Trigger(i.Config)
	}
}
