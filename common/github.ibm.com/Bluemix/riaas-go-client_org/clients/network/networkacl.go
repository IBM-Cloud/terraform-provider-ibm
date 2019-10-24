package network

import (
	"errors"

	riaaserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/Bluemix/riaas-go-client/riaas/client/network"
	"github.ibm.com/Bluemix/riaas-go-client/riaas/models"
	"github.ibm.com/Bluemix/riaas-go-client/session"
	"github.ibm.com/Bluemix/riaas-go-client/utils"
)

// NetworkAclClient ...
type NetworkAclClient struct {
	session *session.Session
}

// NetworkAclClient ...
func NewNetworkAclClient(sess *session.Session) *NetworkAclClient {
	return &NetworkAclClient{
		sess,
	}
}

// List ...
func (f *NetworkAclClient) List(start string) ([]*models.NetworkACL, string, error) {
	return f.ListWithFilter("", start)
}

// ListWithFilter ...
func (f *NetworkAclClient) ListWithFilter(resourcegroupID, start string) ([]*models.NetworkACL, string, error) {
	params := network.NewGetNetworkAclsParamsWithTimeout(f.session.Timeout)
	if resourcegroupID != "" {
		params = params.WithResourceGroupID(&resourcegroupID)
	}

	if start != "" {
		params = params.WithStart(&start)
	}
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation

	resp, err := f.session.Riaas.Network.GetNetworkAcls(params, session.Auth(f.session))
	if err != nil {
		return nil, "", riaaserrors.ToError(err)
	}

	return resp.Payload.NetworkAcls, utils.GetNext(resp.Payload.Next), nil
}

// Get ...
func (f *NetworkAclClient) Get(id string) (*models.NetworkACL, error) {
	params := network.NewGetNetworkAclsIDParamsWithTimeout(f.session.Timeout).WithID(id)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Network.GetNetworkAclsID(params, session.Auth(f.session))

	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// Create ...
func (f *NetworkAclClient) Create(acldef *models.PostNetworkAclsParamsBody) (*models.NetworkACL, error) {
	params := network.NewPostNetworkAclsParamsWithTimeout(f.session.Timeout).WithBody(acldef)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Network.PostNetworkAcls(params, session.Auth(f.session))
	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// Delete ...
func (f *NetworkAclClient) Delete(id string) error {
	params := network.NewDeleteNetworkAclsIDParamsWithTimeout(f.session.Timeout).WithID(id)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	_, err := f.session.Riaas.Network.DeleteNetworkAclsID(params, session.Auth(f.session))
	return riaaserrors.ToError(err)
}

// Update ...
func (f *NetworkAclClient) Update(id, name string) (*models.NetworkACL, error) {
	var body = models.PatchNetworkAclsIDParamsBody{
		Name: name,
	}
	params := network.NewPatchNetworkAclsIDParamsWithTimeout(f.session.Timeout).WithID(id).WithBody(&body)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Network.PatchNetworkAclsID(params, session.Auth(f.session))
	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// ListRules ...
func (f *NetworkAclClient) ListRules(aclID, start string) ([]*models.NetworkACLRule, string, error) {
	params := network.NewGetNetworkAclsNetworkACLIDRulesParamsWithTimeout(f.session.Timeout)
	params = params.WithNetworkACLID(aclID)
	if start != "" {
		params = params.WithStart(&start)
	}
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation

	resp, err := f.session.Riaas.Network.GetNetworkAclsNetworkACLIDRules(params, session.Auth(f.session))

	if err != nil {
		return nil, "", riaaserrors.ToError(err)
	}

	return resp.Payload.Rules, utils.GetNext(resp.Payload.Next), nil
}

// AddRule ...
func (f *NetworkAclClient) AddRule(aclID, name, source, destination, direction, action, protocol string,
	icmpType, icmpCode, portMin, portMax int64,
	before string) (*models.NetworkACLRule, error) {

	rule := models.PostNetworkAclsNetworkACLIDRulesParamsBody{
		Name:      name,
		Direction: direction,
		Protocol:  protocol,
		Action:    action,
	}

	if source != "" {
		rule.Source = source
	}
	if destination != "" {
		rule.Destination = destination
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
			rule.PortMax = portMax
		}
		if portMin > 0 {
			rule.PortMin = portMin
		}
	} else {
		return nil, errors.New("Unknown protocol " + protocol)
	}

	if before != "" {
		rule.Before = &models.PostNetworkAclsNetworkACLIDRulesParamsBodyBefore{
			ID: before,
		}
	}

	params := network.NewPostNetworkAclsNetworkACLIDRulesParamsWithTimeout(f.session.Timeout)
	params = params.WithNetworkACLID(aclID).WithBody(&rule)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation

	resp, err := f.session.Riaas.Network.PostNetworkAclsNetworkACLIDRules(params, session.Auth(f.session))

	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// DeleteRule ...
func (f *NetworkAclClient) DeleteRule(aclID, ruleID string) error {
	params := network.NewDeleteNetworkAclsNetworkACLIDRulesIDParamsWithTimeout(f.session.Timeout).WithNetworkACLID(aclID).WithID(ruleID)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	_, err := f.session.Riaas.Network.DeleteNetworkAclsNetworkACLIDRulesID(params, session.Auth(f.session))
	return riaaserrors.ToError(err)
}

// GetRule ...
func (f *NetworkAclClient) GetRule(aclID, ruleID string) (*models.NetworkACLRule, error) {
	params := network.NewGetNetworkAclsNetworkACLIDRulesIDParamsWithTimeout(f.session.Timeout).WithNetworkACLID(aclID).WithID(ruleID)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Network.GetNetworkAclsNetworkACLIDRulesID(params, session.Auth(f.session))
	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// UpdateRule ...
func (f *NetworkAclClient) UpdateRule(aclID, ruleID, name, source, destination, direction, action, protocol string,
	portMin, portMax, icmpType, icmpCode int64,
	before string) (*models.NetworkACLRule, error) {

	params := network.NewPatchNetworkAclsNetworkACLIDRulesIDParamsWithTimeout(f.session.Timeout).WithNetworkACLID(aclID).WithID(ruleID)
	rule := models.PatchNetworkAclsNetworkACLIDRulesIDParamsBody{}

	if name != "" {
		rule.Name = name
	}

	if source != "" {
		rule.Source = source
	}

	if destination != "" {
		rule.Destination = destination
	}

	if direction != "" {
		rule.Direction = direction
	}

	if action != "" {
		rule.Action = action
	}

	if protocol != "" {
		rule.Protocol = protocol
	}

	if icmpCode >= 0 {
		rule.Type = &icmpCode
	}
	if icmpType >= 0 {
		rule.Code = &icmpType
	}

	if portMax >= 0 {
		rule.PortMax = portMax
	}
	if portMin >= 0 {
		rule.PortMin = portMin
	}

	if before != "" {
		rule.Before = &models.PatchNetworkAclsNetworkACLIDRulesIDParamsBodyBefore{
			ID: before,
		}
	}

	params = params.WithBody(&rule)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation

	resp, err := f.session.Riaas.Network.PatchNetworkAclsNetworkACLIDRulesID(params, session.Auth(f.session))

	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}
