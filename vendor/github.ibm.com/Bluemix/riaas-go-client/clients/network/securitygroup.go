package network

import (
	"github.com/go-openapi/strfmt"
	"github.ibm.com/Bluemix/riaas-go-client/riaas/client/network"
	"github.ibm.com/Bluemix/riaas-go-client/riaas/models"

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
	return f.ListWithFilter("", "", start)
}

// ListWithFilter ...
func (f *SecurityGroupClient) ListWithFilter(vpcID, resourcegroupID, start string) ([]*models.SecurityGroup, string, error) {
	params := network.NewGetSecurityGroupsParamsWithTimeout(f.session.Timeout)

	if vpcID != "" {
		params = params.WithVpcID(&vpcID)
	}
	if resourcegroupID != "" {
		params = params.WithResourceGroupID(&resourcegroupID)
	}
	if start != "" {
		params = params.WithStart(&start)
	}
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation

	resp, err := f.session.Riaas.Network.GetSecurityGroups(params, session.Auth(f.session))

	if err != nil {
		return nil, "", riaaserrors.ToError(err)
	}

	return resp.Payload.SecurityGroups, utils.GetNext(resp.Payload.Next), nil
}

// Get ...
func (f *SecurityGroupClient) Get(id string) (*models.SecurityGroup, error) {
	params := network.NewGetSecurityGroupsIDParamsWithTimeout(f.session.Timeout).WithID(id)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Network.GetSecurityGroupsID(params, session.Auth(f.session))

	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// Create ...
func (f *SecurityGroupClient) Create(sgdef *models.PostSecurityGroupsParamsBody) (*models.SecurityGroup, error) {
	params := network.NewPostSecurityGroupsParamsWithTimeout(f.session.Timeout).WithBody(sgdef)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Network.PostSecurityGroups(params, session.Auth(f.session))
	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// Delete ...
func (f *SecurityGroupClient) Delete(id string) error {
	params := network.NewDeleteSecurityGroupsIDParamsWithTimeout(f.session.Timeout).WithID(id)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	_, err := f.session.Riaas.Network.DeleteSecurityGroupsID(params, session.Auth(f.session))
	return riaaserrors.ToError(err)
}

// Update ...
func (f *SecurityGroupClient) Update(id, name string) (*models.SecurityGroup, error) {
	var body = models.PatchSecurityGroupsIDParamsBody{
		Name: name,
	}

	params := network.NewPatchSecurityGroupsIDParamsWithTimeout(f.session.Timeout).WithID(id).WithRequestBody(&body)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Network.PatchSecurityGroupsID(params, session.Auth(f.session))
	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// ListNetworkInterfaces ...
func (f *SecurityGroupClient) ListNetworkInterfaces(secgrpID string) ([]*models.ServerNetworkInterface, error) {
	params := network.NewGetSecurityGroupsSecurityGroupIDNetworkInterfacesParamsWithTimeout(f.session.Timeout)
	params = params.WithSecurityGroupID(secgrpID)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation

	resp, err := f.session.Riaas.Network.GetSecurityGroupsSecurityGroupIDNetworkInterfaces(params, session.Auth(f.session))

	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload.NetworkInterfaces, nil
}

// DeleteNetworkInterface ...
func (f *SecurityGroupClient) DeleteNetworkInterface(secgrpID, networkIntfID string) error {
	params := network.NewDeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParamsWithTimeout(f.session.Timeout).WithSecurityGroupID(secgrpID).WithID(networkIntfID)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	_, err := f.session.Riaas.Network.DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesID(params, session.Auth(f.session))
	return riaaserrors.ToError(err)
}

// GetNetworkInterface ...
func (f *SecurityGroupClient) GetNetworkInterface(secgrpID, networkIntfID string) (*models.ServerNetworkInterface, error) {
	params := network.NewGetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParamsWithTimeout(f.session.Timeout).WithSecurityGroupID(secgrpID).WithID(networkIntfID)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Network.GetSecurityGroupsSecurityGroupIDNetworkInterfacesID(params, session.Auth(f.session))
	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// AddNetworkInterface ...
func (f *SecurityGroupClient) AddNetworkInterface(secgrpID, networkIntfID string) (*models.ServerNetworkInterface, error) {
	params := network.NewPutSecurityGroupsSecurityGroupIDNetworkInterfacesIDParamsWithTimeout(f.session.Timeout).WithSecurityGroupID(secgrpID).WithID(networkIntfID)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Network.PutSecurityGroupsSecurityGroupIDNetworkInterfacesID(params, session.Auth(f.session))
	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// ListRules ...
func (f *SecurityGroupClient) ListRules(secgrpID string) ([]*models.SecurityGroupRule, error) {
	params := network.NewGetSecurityGroupsSecurityGroupIDRulesParamsWithTimeout(f.session.Timeout)
	params = params.WithSecurityGroupID(secgrpID)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation

	resp, err := f.session.Riaas.Network.GetSecurityGroupsSecurityGroupIDRules(params, session.Auth(f.session))

	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload.Rules, nil
}

// AddRule ...
func (f *SecurityGroupClient) AddRule(secgrpID, direction, ipversion, protocol, remoteAddress, remoteCIDR, remoteSecGrpID string,
	icmpType, icmpCode, portMin, portMax int64) (*models.SecurityGroupRule, error) {

	remote := models.PostSecurityGroupsSecurityGroupIDRulesParamsBodyRemote{}

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

	rule := models.PostSecurityGroupsSecurityGroupIDRulesParamsBody{
		Direction: &direction,
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

	params := network.NewPostSecurityGroupsSecurityGroupIDRulesParamsWithTimeout(f.session.Timeout)
	params = params.WithSecurityGroupID(secgrpID).WithBody(&rule)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation

	resp, err := f.session.Riaas.Network.PostSecurityGroupsSecurityGroupIDRules(params, session.Auth(f.session))

	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// DeleteRule ...
func (f *SecurityGroupClient) DeleteRule(secgrpID, ruleID string) error {
	params := network.NewDeleteSecurityGroupsSecurityGroupIDRulesIDParamsWithTimeout(f.session.Timeout).WithSecurityGroupID(secgrpID).WithID(ruleID)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	_, err := f.session.Riaas.Network.DeleteSecurityGroupsSecurityGroupIDRulesID(params, session.Auth(f.session))
	return riaaserrors.ToError(err)
}

// GetRule ...
func (f *SecurityGroupClient) GetRule(secgrpID, ruleID string) (*models.SecurityGroupRule, error) {
	params := network.NewGetSecurityGroupsSecurityGroupIDRulesIDParamsWithTimeout(f.session.Timeout).WithSecurityGroupID(secgrpID).WithID(ruleID)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Network.GetSecurityGroupsSecurityGroupIDRulesID(params, session.Auth(f.session))
	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// UpdateRule ...
func (f *SecurityGroupClient) UpdateRule(secgrpID, ruleID, direction, ipversion, protocol, remoteAddress, remoteCIDR, remoteSecGrpID string,
	icmpType, icmpCode, portMin, portMax int64) (*models.SecurityGroupRule, error) {

	remote := models.PatchSecurityGroupsSecurityGroupIDRulesIDParamsBodyRemote{}

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

	rule := models.PatchSecurityGroupsSecurityGroupIDRulesIDParamsBody{}

	if direction != "" {
		rule.Direction = direction
	}

	if ipversion != "" {
		rule.IPVersion = ipversion
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

	params := network.NewPatchSecurityGroupsSecurityGroupIDRulesIDParamsWithTimeout(f.session.Timeout)
	params = params.WithSecurityGroupID(secgrpID).WithBody(&rule).WithID(ruleID)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation

	resp, err := f.session.Riaas.Network.PatchSecurityGroupsSecurityGroupIDRulesID(params, session.Auth(f.session))

	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}
