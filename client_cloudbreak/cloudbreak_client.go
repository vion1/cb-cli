package client_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	httptransport "github.com/go-swagger/go-swagger/httpkit/client"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/client_cloudbreak/accountpreferences"
	"github.com/hortonworks/hdc-cli/client_cloudbreak/blueprints"
	"github.com/hortonworks/hdc-cli/client_cloudbreak/cluster"
	"github.com/hortonworks/hdc-cli/client_cloudbreak/clustertemplates"
	"github.com/hortonworks/hdc-cli/client_cloudbreak/connectors"
	"github.com/hortonworks/hdc-cli/client_cloudbreak/constraints"
	"github.com/hortonworks/hdc-cli/client_cloudbreak/credentials"
	"github.com/hortonworks/hdc-cli/client_cloudbreak/events"
	"github.com/hortonworks/hdc-cli/client_cloudbreak/flexsubscriptions"
	"github.com/hortonworks/hdc-cli/client_cloudbreak/ldap"
	"github.com/hortonworks/hdc-cli/client_cloudbreak/networks"
	"github.com/hortonworks/hdc-cli/client_cloudbreak/rdsconfigs"
	"github.com/hortonworks/hdc-cli/client_cloudbreak/recipes"
	"github.com/hortonworks/hdc-cli/client_cloudbreak/securitygroups"
	"github.com/hortonworks/hdc-cli/client_cloudbreak/settings"
	"github.com/hortonworks/hdc-cli/client_cloudbreak/smartsensesubscriptions"
	"github.com/hortonworks/hdc-cli/client_cloudbreak/sssd"
	"github.com/hortonworks/hdc-cli/client_cloudbreak/stacks"
	"github.com/hortonworks/hdc-cli/client_cloudbreak/subscriptions"
	"github.com/hortonworks/hdc-cli/client_cloudbreak/templates"
	"github.com/hortonworks/hdc-cli/client_cloudbreak/topologies"
	"github.com/hortonworks/hdc-cli/client_cloudbreak/usages"
	"github.com/hortonworks/hdc-cli/client_cloudbreak/users"
	"github.com/hortonworks/hdc-cli/client_cloudbreak/util"
)

// Default cloudbreak HTTP client.
var Default = NewHTTPClient(nil)

// NewHTTPClient creates a new cloudbreak HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Cloudbreak {
	if formats == nil {
		formats = strfmt.Default
	}
	transport := httptransport.New("localhost", "/", []string{"https", "http"})
	return New(transport, formats)
}

// New creates a new cloudbreak client
func New(transport client.Transport, formats strfmt.Registry) *Cloudbreak {
	cli := new(Cloudbreak)
	cli.Transport = transport

	cli.Accountpreferences = accountpreferences.New(transport, formats)

	cli.Blueprints = blueprints.New(transport, formats)

	cli.Cluster = cluster.New(transport, formats)

	cli.Clustertemplates = clustertemplates.New(transport, formats)

	cli.Connectors = connectors.New(transport, formats)

	cli.Constraints = constraints.New(transport, formats)

	cli.Credentials = credentials.New(transport, formats)

	cli.Events = events.New(transport, formats)

	cli.Flexsubscriptions = flexsubscriptions.New(transport, formats)

	cli.Ldap = ldap.New(transport, formats)

	cli.Networks = networks.New(transport, formats)

	cli.Rdsconfigs = rdsconfigs.New(transport, formats)

	cli.Recipes = recipes.New(transport, formats)

	cli.Securitygroups = securitygroups.New(transport, formats)

	cli.Settings = settings.New(transport, formats)

	cli.Smartsensesubscriptions = smartsensesubscriptions.New(transport, formats)

	cli.Sssd = sssd.New(transport, formats)

	cli.Stacks = stacks.New(transport, formats)

	cli.Subscriptions = subscriptions.New(transport, formats)

	cli.Templates = templates.New(transport, formats)

	cli.Topologies = topologies.New(transport, formats)

	cli.Usages = usages.New(transport, formats)

	cli.Users = users.New(transport, formats)

	cli.Util = util.New(transport, formats)

	return cli
}

// Cloudbreak is a client for cloudbreak
type Cloudbreak struct {
	Accountpreferences *accountpreferences.Client

	Blueprints *blueprints.Client

	Cluster *cluster.Client

	Clustertemplates *clustertemplates.Client

	Connectors *connectors.Client

	Constraints *constraints.Client

	Credentials *credentials.Client

	Events *events.Client

	Flexsubscriptions *flexsubscriptions.Client

	Ldap *ldap.Client

	Networks *networks.Client

	Rdsconfigs *rdsconfigs.Client

	Recipes *recipes.Client

	Securitygroups *securitygroups.Client

	Settings *settings.Client

	Smartsensesubscriptions *smartsensesubscriptions.Client

	Sssd *sssd.Client

	Stacks *stacks.Client

	Subscriptions *subscriptions.Client

	Templates *templates.Client

	Topologies *topologies.Client

	Usages *usages.Client

	Users *users.Client

	Util *util.Client

	Transport client.Transport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Cloudbreak) SetTransport(transport client.Transport) {
	c.Transport = transport

	c.Accountpreferences.SetTransport(transport)

	c.Blueprints.SetTransport(transport)

	c.Cluster.SetTransport(transport)

	c.Clustertemplates.SetTransport(transport)

	c.Connectors.SetTransport(transport)

	c.Constraints.SetTransport(transport)

	c.Credentials.SetTransport(transport)

	c.Events.SetTransport(transport)

	c.Flexsubscriptions.SetTransport(transport)

	c.Ldap.SetTransport(transport)

	c.Networks.SetTransport(transport)

	c.Rdsconfigs.SetTransport(transport)

	c.Recipes.SetTransport(transport)

	c.Securitygroups.SetTransport(transport)

	c.Settings.SetTransport(transport)

	c.Smartsensesubscriptions.SetTransport(transport)

	c.Sssd.SetTransport(transport)

	c.Stacks.SetTransport(transport)

	c.Subscriptions.SetTransport(transport)

	c.Templates.SetTransport(transport)

	c.Topologies.SetTransport(transport)

	c.Usages.SetTransport(transport)

	c.Users.SetTransport(transport)

	c.Util.SetTransport(transport)

}
