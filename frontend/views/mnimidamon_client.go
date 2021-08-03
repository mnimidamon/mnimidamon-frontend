package views

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"golang.org/x/net/context"
	apiclient "mnimidamonbackend/client"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	"mnimidamonbackend/frontend/services"
)

var mnimidamon *apiclient.Mnimidamon
var apiContext = context.Background()

var compAuth = new(compAuthWriter)
var userAuth = new(userAuthWriter)

type compAuthWriter struct{ runtime.ClientAuthInfoWriter }

func (c *compAuthWriter) HandleConfirmConfig(config global.Config) {
	c.ClientAuthInfoWriter = httptransport.APIKeyAuth("X-COMP-KEY", "header", config.Computer.Key)
}

func (c *compAuthWriter) HandleComputerConfirmConfig(config global.ComputerConfig) {
	c.ClientAuthInfoWriter = httptransport.APIKeyAuth("X-COMP-KEY", "header", config.Key)
}

type userAuthWriter struct{ runtime.ClientAuthInfoWriter }

func (u *userAuthWriter) HandleConfirmConfig(config global.Config) {
	u.ClientAuthInfoWriter = httptransport.APIKeyAuth("X-AUTH-KEY", "header", config.User.Key)
}

func (u *userAuthWriter) HandleUserConfirmConfig(config global.UserConfig) {
	u.ClientAuthInfoWriter = httptransport.APIKeyAuth("X-AUTH-KEY", "header", config.Key)
}

func init() {
	events.ConfirmUserConfig.Register(userAuth)
	events.ConfirmComputerConfig.Register(compAuth)
	events.ConfirmConfig.Register(compAuth)
	events.ConfirmConfig.Register(userAuth)

	// If it's stored, retrieve it.
	if services.ConfigurationStore.IsStored() {
		config := services.ConfigurationStore.GetConfig()
		userAuth.HandleConfirmConfig(*config)
		compAuth.HandleConfirmConfig(*config)
	}
}
