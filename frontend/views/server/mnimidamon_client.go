package server

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"golang.org/x/net/context"
	apiclient "mnimidamonbackend/client"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	"mnimidamonbackend/frontend/services"
	"strconv"
)

var Mnimidamon mnimidamon
var ApiContext = context.Background()

var CompAuth = new(compAuthWriter)
var UserAuth = new(userAuthWriter)

type mnimidamon struct {
	*apiclient.Mnimidamon
}

func (m mnimidamon) HandleConfirmConfig(config global.Config) {
	global.Log("base http client created")
	// Make the Mnimidamon http client.
	transport := httptransport.New(config.Server.Host + ":" + strconv.Itoa(config.Server.Port), apiclient.DefaultBasePath, nil)
	Mnimidamon.Mnimidamon = apiclient.New(transport, strfmt.Default)
}

func (m mnimidamon) HandleServerConfirmConfig(payload global.ServerConfig) {
	global.Log("base http client created")
	transport := httptransport.New(payload.Host + ":" + strconv.Itoa(payload.Port), apiclient.DefaultBasePath, nil)
	Mnimidamon.Mnimidamon = apiclient.New(transport, strfmt.Default)
}

type compAuthWriter struct{ runtime.ClientAuthInfoWriter }

func (c *compAuthWriter) HandleConfirmConfig(config global.Config) {
	c.ClientAuthInfoWriter = runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		err := r.SetHeaderParam("X-COMP-KEY", config.Computer.Key)
		if err != nil {
			return err
		}
		return r.SetHeaderParam("X-AUTH-KEY", config.User.Key)
	})
}

type userAuthWriter struct{ runtime.ClientAuthInfoWriter }

func (u *userAuthWriter) HandleConfirmConfig(config global.Config) {
	u.ClientAuthInfoWriter = httptransport.APIKeyAuth("X-AUTH-KEY", "header", config.User.Key)
}

func (u *userAuthWriter) HandleUserConfirmConfig(config global.UserConfig) {
	u.ClientAuthInfoWriter = httptransport.APIKeyAuth("X-AUTH-KEY", "header", config.Key)
}

func init() {
	events.ConfirmUserConfig.Register(UserAuth)

	events.ConfirmConfig.Register(CompAuth)
	events.ConfirmConfig.Register(UserAuth)

	events.ConfirmServerConfig.Register(Mnimidamon)
	events.ConfirmConfig.Register(Mnimidamon)

	// If it's stored, retrieve it.
	if services.ConfigurationStore.IsStored() {
		config := services.ConfigurationStore.GetConfig()
		UserAuth.HandleConfirmConfig(*config)
		CompAuth.HandleConfirmConfig(*config)
		Mnimidamon.HandleConfirmConfig(*config)
	}
}
