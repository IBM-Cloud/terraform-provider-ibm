package network

import (
	"errors"

	"github.ibm.com/riaas/rias-api/riaas/client/network"
	"github.ibm.com/riaas/rias-api/riaas/models"

	"github.com/go-openapi/strfmt"
	riaaserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/Bluemix/riaas-go-client/session"
	"github.ibm.com/Bluemix/riaas-go-client/utils"
)

// SubnetClient ...
type SubnetClient struct {
	session *session.Session
}

// NewSubnetClient ...
func NewSubnetClient(sess *session.Session) *SubnetClient {
	return &SubnetClient{
		sess,
	}
}

// List ...
func (f *SubnetClient) List(start string) ([]*models.Subnet, string, error) {
	return f.ListWithFilter("", "", "", "", "", start)
}

// ListWithFilter ...
func (f *SubnetClient) ListWithFilter(tag, zoneName, vpcID, networkaclID, resourcegroupID, start string) ([]*models.Subnet, string, error) {
	params := network.NewGetSubnetsParams()
	if tag != "" {
		params = params.WithTag(&tag)
	}
	if zoneName != "" {
		params = params.WithZoneName(&zoneName)
	}
	if vpcID != "" {
		params = params.WithVpcID(&vpcID)
	}
	if networkaclID != "" {
		params = params.WithNetworkACLID(&networkaclID)
	}

	if resourcegroupID != "" {
		params = params.WithResourceGroupID(&resourcegroupID)
	}
	params.Version = "2019-03-26"

	resp, err := f.session.Riaas.Network.GetSubnets(params, session.Auth(f.session))

	if err != nil {
		return nil, "", riaaserrors.ToError(err)
	}

	return resp.Payload.Subnets, utils.GetNext(resp.Payload.Next), nil
}

// Get ...
func (f *SubnetClient) Get(id string) (*models.Subnet, error) {
	params := network.NewGetSubnetsIDParams().WithID(id)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Network.GetSubnetsID(params, session.Auth(f.session))

	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// Create ...
func (f *SubnetClient) Create(name, zoneName, vpcID, networkaclID, publicgwID, generation,
	resourcegroupID, ipv4CIDR string, totalIpv4AddressCount int) (*models.Subnet, error) {

	var body = models.PostSubnetsParamsBody{
		Name:      name,
		IPVersion: models.PostSubnetsParamsBodyIPVersionIPV4,
	}

	var zone = models.NameReference{
		Name: zoneName,
	}
	body.Zone = &zone

	var vpc = models.PostSubnetsParamsBodyVpc{
		ID: strfmt.UUID(vpcID),
	}
	body.Vpc = &vpc

	if networkaclID != "" {
		var networkacl = models.PostSubnetsParamsBodyNetworkACL{
			ID: strfmt.UUID(networkaclID),
		}
		body.NetworkACL = &networkacl
	}

	if publicgwID != "" {
		publicgwUUID := strfmt.UUID(publicgwID)
		var pubgw = models.PostSubnetsParamsBodyPublicGateway{
			ID: publicgwUUID,
		}
		body.PublicGateway = &pubgw
	}

	if generation != "" {
		body.Generation = models.Generation(generation)
	}

	if resourcegroupID != "" {
		resourcegroupuuid := strfmt.UUID(resourcegroupID)
		var resourcegroup = models.PostSubnetsParamsBodyResourceGroup{
			ID: resourcegroupuuid,
		}
		body.ResourceGroup = &resourcegroup
	}

	if ipv4CIDR != "" {
		if totalIpv4AddressCount != 0 {
			return nil, errors.New("only one of ipv4CIDR or totalIpv4AddressCount can be set")
		}
		body.IPV4CidrBlock = ipv4CIDR
	}

	if totalIpv4AddressCount != 0 {
		body.TotalIPV4AddressCount = int64(totalIpv4AddressCount)
	}

	params := network.NewPostSubnetsParams().WithBody(&body)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Network.PostSubnets(params, session.Auth(f.session))
	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// DetachPublicGateway ...
func (f *SubnetClient) DetachPublicGateway(id string) error {
	params := network.NewDeleteSubnetsSubnetIDPublicGatewayParams().WithSubnetID(id)
	params.Version = "2019-03-26"
	_, err := f.session.Riaas.Network.DeleteSubnetsSubnetIDPublicGateway(params, session.Auth(f.session))
	return riaaserrors.ToError(err)
}

// Delete ...
func (f *SubnetClient) Delete(id string) error {
	params := network.NewDeleteSubnetsIDParams().WithID(id)
	params.Version = "2019-03-26"
	_, err := f.session.Riaas.Network.DeleteSubnetsID(params, session.Auth(f.session))
	return riaaserrors.ToError(err)
}

// Update ...
func (f *SubnetClient) Update(id, name, networkaclID, publicgwID string) (*models.Subnet, error) {
	var body = models.PatchSubnetsIDParamsBody{}

	if name != "" {
		body.Name = name
	}

	if networkaclID != "" {
		networkaclUUID := strfmt.UUID(networkaclID)
		var networkacl = models.PatchSubnetsIDParamsBodyNetworkACL{
			ID: networkaclUUID,
		}
		body.NetworkACL = &networkacl
	}

	if publicgwID != "" {
		publicgwUUID := strfmt.UUID(publicgwID)
		var publicgw = models.PatchSubnetsIDParamsBodyPublicGateway{
			ID: publicgwUUID,
		}
		body.PublicGateway = &publicgw
	}

	params := network.NewPatchSubnetsIDParams().WithID(id).WithBody(&body)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Network.PatchSubnetsID(params, session.Auth(f.session))
	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}
