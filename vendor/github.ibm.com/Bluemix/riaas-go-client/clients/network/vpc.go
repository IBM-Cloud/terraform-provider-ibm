package network

import (
	"github.com/go-openapi/strfmt"
	"github.ibm.com/Bluemix/riaas-go-client/riaas/client/network"
	"github.ibm.com/Bluemix/riaas-go-client/riaas/client/v_p_cs"
	"github.ibm.com/Bluemix/riaas-go-client/riaas/models"

	"github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/Bluemix/riaas-go-client/session"
	"github.ibm.com/Bluemix/riaas-go-client/utils"
)

// VPCClient ...
type VPCClient struct {
	session *session.Session
}

// NewVPCClient ...
func NewVPCClient(sess *session.Session) *VPCClient {
	return &VPCClient{
		sess,
	}
}

// List ...
func (f *VPCClient) List(start string) ([]*models.Vpc, string, error) {
	return f.ListWithFilter(start, "")
}

// ListWithFilter ...
func (f *VPCClient) ListWithFilter(start, resourcegroupID string) ([]*models.Vpc, string, error) {
	params := network.NewGetVpcsParamsWithTimeout(f.session.Timeout)
	if start != "" {
		params = params.WithStart(&start)
	}
	if resourcegroupID != "" {
		params = params.WithResourceGroupID(&resourcegroupID)
	}
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation

	resp, err := f.session.Riaas.Network.GetVpcs(params, session.Auth(f.session))

	if err != nil {
		return nil, "", errors.ToError(err)
	}

	return resp.Payload.Vpcs, utils.GetNext(resp.Payload.Next), nil
}

// Get ...
func (f *VPCClient) Get(id string) (*models.Vpc, error) {
	params := network.NewGetVpcsIDParamsWithTimeout(f.session.Timeout).WithID(id)
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Network.GetVpcsID(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Create ...
func (f *VPCClient) Create(name string, addressPrefixManagement string, classicAccess bool, defaultacl, rg string) (*models.Vpc, error) {

	var body = network.PostVpcsBody{
		Name:                    name,
		ClassicAccess:           classicAccess,
		AddressPrefixManagement: &addressPrefixManagement,
	}
	if rg != "" {
		body.ResourceGroup = &network.PostVpcsParamsBodyResourceGroup{
			ID: strfmt.UUID(rg),
		}
	}

	params := network.NewPostVpcsParamsWithTimeout(f.session.Timeout).WithBody(body)
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Network.PostVpcs(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Delete ...
func (f *VPCClient) Delete(id string) error {
	params := network.NewDeleteVpcsIDParamsWithTimeout(f.session.Timeout).WithID(id)
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation
	_, err := f.session.Riaas.Network.DeleteVpcsID(params, session.Auth(f.session))
	return errors.ToError(err)
}

// Update ...
func (f *VPCClient) Update(id, name string) (*models.Vpc, error) {
	var body = network.PatchVpcsIDBody{
		Name: name,
	}
	params := network.NewPatchVpcsIDParamsWithTimeout(f.session.Timeout).WithID(id).WithBody(body)
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Network.PatchVpcsID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

/*
// UpdateDefaultNWACL ...
func (f *VPCClient) UpdateDefaultNWACL(id, aclid string) (*models.NetworkACL, error) {
	var body = models.PutVpcsVpcIDDefaultNetworkACLParamsBody{
		Identifier: strfmt.UUID(aclid),
	}
	params := network.NewPutVpcsVpcIDDefaultNetworkACLParamsWithTimeout(f.session.Timeout).WithVpcID(id).WithBody(&body)
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Network.PutVpcsVpcIDDefaultNetworkACL(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}*/

// CreateAddressPrefix ...
func (f *VPCClient) CreateAddressPrefix(addressPrefixes network.PostVpcsVpcIDAddressPrefixesBody, vpcId string) (*models.AddressPrefix, error) {
	params := network.NewPostVpcsVpcIDAddressPrefixesParamsWithTimeout(f.session.Timeout).WithAddressPoolPrefixTemplate(addressPrefixes).WithVpcID(vpcId)
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Network.PostVpcsVpcIDAddressPrefixes(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// UpdateAddressPrefix ...
func (f *VPCClient) UpdateAddressPrefix(addressPrefixes network.PatchVpcsVpcIDAddressPrefixesIDBody, vpcID, addressPrefixID string) (*models.AddressPrefix, error) {
	params := network.NewPatchVpcsVpcIDAddressPrefixesIDParamsWithTimeout(f.session.Timeout).WithAddressPoolPrefixPatch(addressPrefixes).WithID(addressPrefixID).WithVpcID(vpcID)
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Network.PatchVpcsVpcIDAddressPrefixesID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// GetAddressPrefix ...
func (f *VPCClient) GetAddressPrefix(vpcID, addressPrefixesID string) (*models.AddressPrefix, error) {
	params := network.NewGetVpcsVpcIDAddressPrefixesIDParamsWithTimeout(f.session.Timeout).WithVpcID(vpcID).WithID(addressPrefixesID)
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Network.GetVpcsVpcIDAddressPrefixesID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// DeleteAddressPrefix ...
func (f *VPCClient) DeleteAddressPrefix(vpcID, addressPrefixesID string) error {
	params := network.NewDeleteVpcsVpcIDAddressPrefixesIDParamsWithTimeout(f.session.Timeout).WithVpcID(vpcID).WithID(addressPrefixesID)
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation
	_, err := f.session.Riaas.Network.DeleteVpcsVpcIDAddressPrefixesID(params, session.Auth(f.session))
	return errors.ToError(err)
}

// ListPrefixes ...
func (f *VPCClient) ListPrefixes(id string) ([]*models.AddressPrefix, error) {
	params := network.NewGetVpcsVpcIDAddressPrefixesParamsWithTimeout(f.session.Timeout).WithVpcID(id)
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Network.GetVpcsVpcIDAddressPrefixes(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload.AddressPrefixes, nil
}

// GetSecurityGroups ...
func (f *VPCClient) GetSecurityGroup(id string) (*models.DefaultSecurityGroup, error) {
	params := network.NewGetVpcsVpcIDDefaultSecurityGroupParamsWithTimeout(f.session.Timeout).WithVpcID(id)
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Network.GetVpcsVpcIDDefaultSecurityGroup(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// ListRoutes ...
func (f *VPCClient) ListRoutes(vpc_id, zoneName string) ([]*models.Route, error) {
	return f.ListWithFilterRoutes(vpc_id, zoneName)
}

// ListWithFilterRoutes ...
func (f *VPCClient) ListWithFilterRoutes(vpc_id, zoneName string) ([]*models.Route, error) {
	params := v_p_cs.NewListVpcRoutesParamsWithTimeout(f.session.Timeout).WithVpcID(vpc_id)

	if zoneName != "" {
		params = params.WithZoneName(&zoneName)
	}
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation

	resp, err := f.session.Riaas.VPCs.ListVpcRoutes(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload.Routes, nil
}

// CreateRoute ...
func (f *VPCClient) CreateRoute(routes *models.RouteTemplate, vpcID string) (*models.Route, error) {
	params := v_p_cs.NewCreateVpcRouteParamsWithTimeout(f.session.Timeout).WithRouteTemplate(routes).WithVpcID(vpcID)
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.VPCs.CreateVpcRoute(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// DeleteRoute ...
func (f *VPCClient) DeleteRoute(vpcID, routeID string) error {
	params := v_p_cs.NewDeleteVpcRouteParamsWithTimeout(f.session.Timeout).WithVpcID(vpcID).WithID(routeID)
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation
	_, err := f.session.Riaas.VPCs.DeleteVpcRoute(params, session.Auth(f.session))
	return errors.ToError(err)
}

// GetRoute ...
func (f *VPCClient) GetRoute(vpcID, routeID string) (*models.Route, error) {
	params := v_p_cs.NewGetVpcRouteParamsWithTimeout(f.session.Timeout).WithVpcID(vpcID).WithID(routeID)
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.VPCs.GetVpcRoute(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// UpdateRoute ...
func (f *VPCClient) UpdateRoute(routes *models.RoutePatch, vpcID, routeID string) (*models.Route, error) {
	params := v_p_cs.NewUpdateVpcRouteParamsWithTimeout(f.session.Timeout).WithRoutePatch(routes).WithID(routeID).WithVpcID(vpcID)
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.VPCs.UpdateVpcRoute(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}
