package cli

import (
	"time"

	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/hortonworks/hdc-cli/client_cloudbreak/ldap"
	"github.com/hortonworks/hdc-cli/models_cloudbreak"
	"github.com/urfave/cli"
	"strconv"
	"strings"
)

var LdapHeader []string = []string{"Name", "Server", "Domain", "BindDn", "UserSearchBase", "UserSearchFilter", "UserSearchAttribute", "GroupSearchBase"}

type Ldap struct {
	Name                string `json:"Name" yaml:"Name"`
	Server              string `json:"Server" yaml:"Server"`
	Domain              string `json:"Domain,omitempty" yaml:"Domain,omitempty"`
	BindDn              string `json:"BindDn" yaml:"BindDn"`
	UserSearchBase      string `json:"UserSearchBase" yaml:"UserSearchBase"`
	UserSearchFilter    string `json:"UserSearchFilter,omitempty" yaml:"UserSearchFilter,omitempty"`
	UserSearchAttribute string `json:"UserSearchAttribute,omitempty" yaml:"UserSearchAttribute,omitempty"`
	GroupSearchBase     string `json:"GroupSearchBase,omitempty" yaml:"GroupSearchBase,omitempty"`
}

func (l *Ldap) DataAsStringArray() []string {
	return []string{l.Name, l.Server, l.Domain, l.BindDn, l.UserSearchBase, l.UserSearchFilter, l.UserSearchAttribute, l.GroupSearchBase}
}

func CreateLDAP(c *cli.Context) error {
	defer timeTrack(time.Now(), "create LDAP")
	checkRequiredFlags(c)

	name := c.String(FlLdapName.Name)
	domain := c.String(FlLdapDomain.Name)
	bindDn := c.String(FlLdapBindDN.Name)
	bindPassword := c.String(FlLdapBindPassword.Name)
	userSearchBase := c.String(FlLdapUserSearchBase.Name)
	userSearchFilter := c.String(FlLdapUserSearchFilter.Name)
	userSearchAttribute := c.String(FlLdapUserSearchAttribute.Name)
	groupSearchBase := c.String(FlLdapGroupSearchBase.Name)
	server := c.String(FlLdapServer.Name)
	portSeparatorIndex := strings.LastIndex(server, ":")
	if (!strings.Contains(server, "ldap://") && !strings.Contains(server, "ldaps://")) || portSeparatorIndex == -1 {
		logErrorAndExit(errors.New("invalid ldap server address format, e.g: ldaps://10.0.0.1:389"))
	}
	serverPort, _ := strconv.Atoi(server[portSeparatorIndex+1:])
	protocol := server[0:strings.Index(server, ":")]

	cbClient := NewCloudbreakOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))

	return createLDAPImpl(cbClient.Cloudbreak.Ldap.PostPublicLdap, name, server, int32(serverPort), protocol, domain, bindDn, bindPassword,
		userSearchBase, userSearchFilter, userSearchAttribute, groupSearchBase)
}

func createLDAPImpl(createLDAP func(*ldap.PostPublicLdapParams) (*ldap.PostPublicLdapOK, error),
	name string, server string, port int32, protocol string, domain string, bindDn string, bindPassword string,
	userSearchBase string, userSearchFilter string, userSearchAttribute string, groupSearchBase string) error {

	log.Infof("[createLDAPImpl] create LDAP with name: %s", name)
	resp, err := createLDAP(&ldap.PostPublicLdapParams{
		Body: &models_cloudbreak.LdapConfigRequest{
			Name:                name,
			ServerHost:          server[strings.LastIndex(server, "/")+1 : strings.LastIndex(server, ":")],
			ServerPort:          port,
			Protocol:            &protocol,
			Domain:              &domain,
			BindDn:              bindDn,
			BindPassword:        bindPassword,
			UserSearchBase:      userSearchBase,
			UserSearchFilter:    &userSearchFilter,
			UserSearchAttribute: &userSearchAttribute,
			GroupSearchBase:     &groupSearchBase,
		}})

	if err != nil {
		logErrorAndExit(err)
	}

	log.Infof("[createLDAPImpl] LDAP created with name: %s, id: %d", name, *resp.Payload.ID)
	return nil
}

func ListLdaps(c *cli.Context) error {
	defer timeTrack(time.Now(), "list LDAPS")

	cbClient := NewCloudbreakOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))

	output := Output{Format: c.String(FlOutput.Name)}
	return ListLdapsImpl(cbClient.Cloudbreak.Ldap.GetPublicsLdap, output.WriteList)
}

func ListLdapsImpl(getLdaps func(*ldap.GetPublicsLdapParams) (*ldap.GetPublicsLdapOK, error), writer func([]string, []Row)) error {
	resp, err := getLdaps(&ldap.GetPublicsLdapParams{})

	if err != nil {
		logErrorAndExit(err)
	}

	var tableRows []Row
	for _, l := range resp.Payload {
		row := &Ldap{
			Name:                l.Name,
			Server:              fmt.Sprintf("%s://%s:%d", *l.Protocol, l.ServerHost, l.ServerPort),
			Domain:              *l.Domain,
			BindDn:              l.BindDn,
			UserSearchBase:      l.UserSearchBase,
			UserSearchFilter:    *l.UserSearchFilter,
			UserSearchAttribute: *l.UserSearchAttribute,
			GroupSearchBase:     *l.GroupSearchBase,
		}
		tableRows = append(tableRows, row)
	}

	writer(LdapHeader, tableRows)
	return nil
}

func (c *Cloudbreak) GetLdapByName(name string) *models_cloudbreak.LdapConfigResponse {
	defer timeTrack(time.Now(), "get ldap by name")
	log.Infof("[GetLdapByName] get ldap by name: %s", name)

	resp, err := c.Cloudbreak.Ldap.GetPublicLdap(&ldap.GetPublicLdapParams{Name: name})
	if err != nil {
		logErrorAndExit(err)
	}

	id64 := *resp.Payload.ID
	log.Infof("[GetLdapByName] '%s' ldap id: %d", name, id64)
	return resp.Payload
}
