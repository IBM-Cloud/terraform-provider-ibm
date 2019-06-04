package network

import (
	"github.com/go-openapi/strfmt"
	"github.ibm.com/riaas/rias-api/riaas/client/network"
	"github.ibm.com/riaas/rias-api/riaas/models"

	"errors"

	riaaserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/Bluemix/riaas-go-client/session"
	"github.ibm.com/Bluemix/riaas-go-client/utils"
)

const (
	ProtocolAll  = "all"
	ProtocolICMP = "icmp"
	ProtocolTCP  = "tcp"
	ProtocolUDP  = "udp"

	IPV4 = "ipv4"
	IPV6 = "ipv6"

	DirectionInbound  = "inbound"
	DirectionOutbound = "outbound"
)

// SecurityGroupClient ...
type SecurityGroupClient struct {
	session *session.Session
}

// NewSecurityGroupClient ...
func NewSecurityGroupClient(sess *session.Session) *SecurityGroupClient {
	return &SecurityGroupClient{
		sess,
	}
}

// List ...
func (f *SecurityGroupClient) List(start string) ([]*models.SecurityGroup, string, error) {
	return f.ListWithFilter("", "", "", start)
}

// ListWithFilter ...
func (f *SecurityGroupClient) ListWithFilter(tag, vpcID, resourcegroupID, start string) ([]*models.SecurityGroup, string, error) {
	params := network.NewGetSecurityGroupsParams()
	if tag != "" {
		params = params.WithTag(&tag)
	}
	if vpcID != "" {
		params = params.WithVpcID(&vpcID)
	}
	if resourcegroupID != "" {
		params = params.WithResourceGroupID(&resourcegroupID)
	}
	if start != "" {
		params = params.WithStart(&start)
	}
	params.Version = "2019-03-26"

	resp, err := f.session.Riaas.Network.GetSecurityGroups(params, session.Auth(f.session))

	if err != nil {
		return nil, "", riaaserrors.ToError(err)
	}

	return resp.Payload.SecurityGroups, utils.GetNext(resp.Payload.Next), nil
}

// Get ...
func (f *SecurityGroupClient) Get(id string) (*models.SecurityGroup, error) {
	params := network.NewGetSecurityGroupsIDParams().WithID(id)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Network.GetSecurityGroupsID(params, session.Auth(f.session))

	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// Create ...
func (f *SecurityGroupClient) Create(sgdef *models.PostSecurityGroupsParamsBody) (*models.SecurityGroup, error) {
	params := network.NewPostSecurityGroupsParams().WithBody(sgdef)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Network.PostSecurityGroups(params, session.Auth(f.session))
	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// Delete ...
func (f *SecurityGroupClient) Delete(id string) error {
	params := network.NewDeleteSecurityGroupsIDParams().WithID(id)
	params.Version = "2019-03-26"
	_, err := f.session.Riaas.Network.DeleteSecurityGroupsID(params, session.Auth(f.session))
	return riaaserrors.ToError(err)
}

// Update ...
func (f *SecurityGroupClient) Update(id, name string) (*models.SecurityGroup, error) {
	var body = models.PatchSecurityGroupsIDParamsBody{
		Name: name,
	}

	params := network.NewPatchSecurityGroupsIDParams().WithID(id).WithBody(&body)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Network.PatchSecurityGroupsID(params, session.Auth(f.session))
	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// ListNetworkInterfaces ...
func (f *SecurityGroupClient) ListNetworkInterfaces(secgrpID string) ([]*models.InstanceNetworkInterface, error) {
	params := network.NewGetSecurityGroupsSecurityGroupIDNetworkInterfacesParams()
	params = params.WithSecurityGroupID(secgrpID)
	params.Version = "2019-03-26"

	resp, err := f.session.Riaas.Network.GetSecurityGroupsSecurityGroupIDNetworkInterfaces(params, session.Auth(f.session))

	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload.NetworkInterfaces, nil
}

// DeleteNetworkInterface ...
func (f *SecurityGroupClient) DeleteNetworkInterface(secgrpID, networkIntfID string) error {
	params := network.NewDeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams().WithSecurityGroupID(secgrpID).WithID(networkIntfID)
	params.Version = "2019-03-26"
	_, err := f.session.Riaas.Network.DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesID(params, session.Auth(f.session))
	return riaaserrors.ToError(err)
}

// GetNetworkInterface ...
func (f *SecurityGroupClient) GetNetworkInterface(secgrpID, networkIntfID string) (*models.InstanceNetworkInterface, error) {
	params := network.NewGetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams().WithSecurityGroupID(secgrpID).WithID(networkIntfID)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Network.GetSecurityGroupsSecurityGroupIDNetworkInterfacesID(params, session.Auth(f.session))
	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// AddNetworkInterface ...
func (f *SecurityGroupClient) AddNetworkInterface(secgrpID, networkIntfID string) (*models.InstanceNetworkInterface, error) {
	params := network.NewPutSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams().WithSecurityGroupID(secgrpID).WithID(networkIntfID)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Network.PutSecurityGroupsSecurityGroupIDNetworkInterfacesID(params, session.Auth(f.session))
	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// ListRules ...
func (f *SecurityGroupClient) ListRules(secgrpID string) ([]*models.SecurityGroupRule, error) {
	params := network.NewGetSecurityGroupsSecurityGroupIDRulesParams()
	params = params.WithSecurityGroupID(secgrpID)
	params.Version = "2019-03-26"

	resp, err := f.session.Riaas.Network.GetSecurityGroupsSecurityGroupIDRules(params, session.Auth(f.session))

	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload.Rules, nil
}

// AddRule ...
func (f *SecurityGroupClient) AddRule(secgrpID, direction, ipversion, protocol, remoteAddress, remoteCIDR, remoteSecGrpID string,
	icmpType, icmpCode, portMin, portMax int64) (*models.SecurityGroupRule, error) {

	remote := models.SecurityGroupRuleTemplateRemote{}

	if remoteAddress != "" {
		if remoteCIDR != "" || remoteSecGrpID != "" {
			return nil, errors.New("only one remote field is allowed")
		}
		remote.Address = remoteAddress
	} else if remoteCIDR != "" {
		if remoteSecGrpID != "" {
			return nil, errors.New("only one remote field is allowed")
		}
		remote.CidrBlock = remoteCIDR
	} else if remoteSecGrpID != "" {
		remote.ID = strfmt.UUID(remoteSecGrpID)
	}

	rule := models.SecurityGroupRuleTemplate{
		Direction: direction,
		IPVersion: ipversion,
		Protocol:  &protocol,
	}
	if remoteAddress != "" || remoteCIDR != "" || remoteSecGrpID != "" {
		rule.Remote = &remote
	}

	if protocol == ProtocolAll {

	} else if protocol == ProtocolICMP {
		if icmpCode >= 0 {
			rule.Code = &icmpCode
		}
		if icmpType >= 0 {
			rule.Type = &icmpType
		}
	} else if protocol == ProtocolTCP || protocol == ProtocolUDP {
		if portMax >= 0 {
			rule.PortMax = &portMax
		}
		if portMin > 0 {
			rule.PortMin = &portMin
		}
	} else {
		return nil, errors.New("Unknown protocol " + protocol)
	}

	params := network.NewPostSecurityGroupsSecurityGroupIDRulesParams()
	params = params.WithSecurityGroupID(secgrpID).WithBody(&rule)
	params.Version = "2019-03-26"

	resp, err := f.session.Riaas.Network.PostSecurityGroupsSecurityGroupIDRules(params, session.Auth(f.session))

	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// DeleteRule ...
func (f *SecurityGroupClient) DeleteRule(secgrpID, ruleID string) error {
	params := network.NewDeleteSecurityGroupsSecurityGroupIDRulesIDParams().WithSecurityGroupID(secgrpID).WithID(ruleID)
	params.Version = "2019-03-26"
	_, err := f.session.Riaas.Network.DeleteSecurityGroupsSecurityGroupIDRulesID(params, session.Auth(f.session))
	return riaaserrors.ToError(err)
}

// GetRule ...
func (f *SecurityGroupClient) GetRule(secgrpID, ruleID string) (*models.SecurityGroupRule, error) {
	params := network.NewGetSecurityGroupsSecurityGroupIDRulesIDParams().WithSecurityGroupID(secgrpID).WithID(ruleID)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Network.GetSecurityGroupsSecurityGroupIDRulesID(params, session.Auth(f.session))
	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// UpdateRule ...
func (f *SecurityGroupClient) UpdateRule(secgrpID, ruleID, direction, ipversion, protocol, remoteAddress, remoteCIDR, remoteSecGrpID string,
	icmpType, icmpCode, portMin, portMax int64) (*models.SecurityGroupRule, error) {

	remote := models.SecurityGroupRuleTemplateRemote{}

	if remoteAddress != "" {
		if remoteCIDR != "" || remoteSecGrpID != "" {
			return nil, errors.New("only one remote field is allowed")
		}
		remote.Address = remoteAddress
	} else if remoteCIDR != "" {
		if remoteSecGrpID != "" {
			return nil, errors.New("only one remote field is allowed")
		}
		remote.CidrBlock = remoteCIDR
	} else if remoteSecGrpID != "" {
		remote.ID = strfmt.UUID(remoteSecGrpID)
	}

	rule := models.SecurityGroupRuleTemplate{}

	if direction != "" {
		rule.Direction = direction
	}

	if ipversion != "" {
		rule.IPVersion = ipversion
	}

	if protocol != "" {
		rule.Protocol = &protocol
	}

	if icmpCode >= 0 {
		rule.Type = &icmpCode
	}
	if icmpType >= 0 {
		rule.Code = &icmpType
	}

	if portMax >= 0 {
		rule.PortMax = &portMax
	}
	if portMin >= 0 {
		rule.PortMin = &portMin
	}

	if remoteAddress != "" || remoteCIDR != "" || remoteSecGrpID != "" {
		rule.Remote = &remote
	}

	params := network.NewPatchSecurityGroupsSecurityGroupIDRulesIDParams()
	params = params.WithSecurityGroupID(secgrpID).WithBody(&rule).WithID(ruleID)
	params.Version = "2019-03-26"

	resp, err := f.session.Riaas.Network.PatchSecurityGroupsSecurityGroupIDRulesID(params, session.Auth(f.session))

	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}
