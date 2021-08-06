package server

import (
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"golang.org/x/net/context"
	apiclient "mnimidamonbackend/client"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	"strconv"
)

var Mnimidamon mnimidamon
var ApiContext = context.Background()


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

func init() {
	events.ConfirmServerConfig.Register(Mnimidamon)
	events.ConfirmConfig.Register(Mnimidamon)
}
