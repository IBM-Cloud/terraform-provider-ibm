package network

import (
	"github.com/go-openapi/strfmt"
	"github.ibm.com/riaas/rias-api/riaas/client/network"
	"github.ibm.com/riaas/rias-api/riaas/models"

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
	return f.ListWithFilter("", start, "")
}

// ListWithFilter ...
func (f *VPCClient) ListWithFilter(tag, start, resourcegroupID string) ([]*models.Vpc, string, error) {
	params := network.NewGetVpcsParams()
	if tag != "" {
		params = params.WithTag(&tag)
	}
	if start != "" {
		params = params.WithStart(&start)
	}
	if resourcegroupID != "" {
		params = params.WithResourceGroupID(&resourcegroupID)
	}
	params.Version = "2019-03-26"

	resp, err := f.session.Riaas.Network.GetVpcs(params, session.Auth(f.session))

	if err != nil {
		return nil, "", errors.ToError(err)
	}

	return resp.Payload.Vpcs, utils.GetNext(resp.Payload.Next), nil
}

// Get ...
func (f *VPCClient) Get(id string) (*models.Vpc, error) {
	params := network.NewGetVpcsIDParams().WithID(id)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Network.GetVpcsID(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Create ...
func (f *VPCClient) Create(name string, classicAccess bool, defaultacl, rg string) (*models.Vpc, error) {

	var body = models.PostVpcsParamsBody{
		Name:          name,
		ClassicAccess: classicAccess,
	}
	if defaultacl != "" {
		body.DefaultNetworkACL = &models.PostVpcsParamsBodyDefaultNetworkACL{
			ID: strfmt.UUID(defaultacl),
		}
	}
	if rg != "" {
		body.ResourceGroup = &models.PostVpcsParamsBodyResourceGroup{
			ID: strfmt.UUID(rg),
		}
	}

	params := network.NewPostVpcsParams().WithBody(&body)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Network.PostVpcs(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Delete ...
func (f *VPCClient) Delete(id string) error {
	params := network.NewDeleteVpcsIDParams().WithID(id)
	params.Version = "2019-03-26"
	_, err := f.session.Riaas.Network.DeleteVpcsID(params, session.Auth(f.session))
	return errors.ToError(err)
}

// Update ...
func (f *VPCClient) Update(id, name string) (*models.Vpc, error) {
	var body = models.PatchVpcsIDParamsBody{
		Name: name,
	}
	params := network.NewPatchVpcsIDParams().WithID(id).WithBody(&body)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Network.PatchVpcsID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// UpdateDefaultNWACL ...
func (f *VPCClient) UpdateDefaultNWACL(id, aclid string) (*models.NetworkACL, error) {
	var body = models.ResourceReference{
		ID: strfmt.UUID(aclid),
	}
	params := network.NewPutVpcsVpcIDDefaultNetworkACLParams().WithVpcID(id).WithBody(&body)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Network.PutVpcsVpcIDDefaultNetworkACL(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// CreateAddressPrefix ...
func (f *VPCClient) CreateAddressPrefix(addressPrefixes *models.PostVpcsVpcIDAddressPrefixesParamsBody, vpcId string) (*models.AddressPoolPrefix, error) {
	params := network.NewPostVpcsVpcIDAddressPrefixesParams().WithBody(addressPrefixes).WithVpcID(vpcId)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Network.PostVpcsVpcIDAddressPrefixes(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// UpdateAddressPrefix ...
func (f *VPCClient) UpdateAddressPrefix(addressPrefixes *models.PatchVpcsVpcIDAddressPrefixesIDParamsBody, vpcID, addressPrefixID string) (*models.AddressPoolPrefix, error) {
	params := network.NewPatchVpcsVpcIDAddressPrefixesIDParams().WithBody(addressPrefixes).WithID(addressPrefixID).WithVpcID(vpcID)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Network.PatchVpcsVpcIDAddressPrefixesID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// GetAddressPrefix ...
func (f *VPCClient) GetAddressPrefix(vpcID, addressPrefixesID string) (*models.AddressPoolPrefix, error) {
	params := network.NewGetVpcsVpcIDAddressPrefixesIDParams().WithVpcID(vpcID).WithID(addressPrefixesID)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Network.GetVpcsVpcIDAddressPrefixesID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// DeleteAddressPrefix ...
func (f *VPCClient) DeleteAddressPrefix(vpcID, addressPrefixesID string) error {
	params := network.NewDeleteVpcsVpcIDAddressPrefixesIDParams().WithVpcID(vpcID).WithID(addressPrefixesID)
	params.Version = "2019-03-26"
	_, err := f.session.Riaas.Network.DeleteVpcsVpcIDAddressPrefixesID(params, session.Auth(f.session))
	return errors.ToError(err)
}

// ListPrefixes ...
func (f *VPCClient) ListPrefixes(id string) ([]*models.AddressPoolPrefix, error) {
	params := network.NewGetVpcsVpcIDAddressPrefixesParams().WithVpcID(id)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Network.GetVpcsVpcIDAddressPrefixes(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload.AddressPrefixes, nil
}

// GetSecurityGroups ...
func (f *VPCClient) GetSecurityGroup(id string) (*models.DefaultSecurityGroup, error) {
	params := network.NewGetVpcsVpcIDDefaultSecurityGroupParams().WithVpcID(id)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Network.GetVpcsVpcIDDefaultSecurityGroup(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}
