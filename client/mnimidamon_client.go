// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"mnimidamonbackend/client/authorization"
	"mnimidamonbackend/client/backup"
	"mnimidamonbackend/client/computer"
	"mnimidamonbackend/client/current_user"
	"mnimidamonbackend/client/group"
	"mnimidamonbackend/client/group_computer"
	"mnimidamonbackend/client/invite"
	"mnimidamonbackend/client/user"
)

// Default mnimidamon HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "mnimidamon.marmiha.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api/v1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new mnimidamon HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Mnimidamon {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new mnimidamon HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Mnimidamon {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new mnimidamon client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Mnimidamon {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Mnimidamon)
	cli.Transport = transport
	cli.Authorization = authorization.New(transport, formats)
	cli.Backup = backup.New(transport, formats)
	cli.Computer = computer.New(transport, formats)
	cli.CurrentUser = current_user.New(transport, formats)
	cli.Group = group.New(transport, formats)
	cli.GroupComputer = group_computer.New(transport, formats)
	cli.Invite = invite.New(transport, formats)
	cli.User = user.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Mnimidamon is a client for mnimidamon
type Mnimidamon struct {
	Authorization authorization.ClientService

	Backup backup.ClientService

	Computer computer.ClientService

	CurrentUser current_user.ClientService

	Group group.ClientService

	GroupComputer group_computer.ClientService

	Invite invite.ClientService

	User user.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Mnimidamon) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Authorization.SetTransport(transport)
	c.Backup.SetTransport(transport)
	c.Computer.SetTransport(transport)
	c.CurrentUser.SetTransport(transport)
	c.Group.SetTransport(transport)
	c.GroupComputer.SetTransport(transport)
	c.Invite.SetTransport(transport)
	c.User.SetTransport(transport)
}
