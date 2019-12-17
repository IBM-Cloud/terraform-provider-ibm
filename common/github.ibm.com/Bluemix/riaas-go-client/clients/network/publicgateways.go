package network

import (
	"github.com/go-openapi/strfmt"
	"github.ibm.com/Bluemix/riaas-go-client/errors"
	riaaserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/Bluemix/riaas-go-client/riaas/client/network"
	"github.ibm.com/Bluemix/riaas-go-client/riaas/models"
	"github.ibm.com/Bluemix/riaas-go-client/session"
	"github.ibm.com/Bluemix/riaas-go-client/utils"
)

// PublicGatewayClient ...
type PublicGatewayClient struct {
	session *session.Session
}

// PublicGatewayClient ...
func NewPublicGatewayClient(sess *session.Session) *PublicGatewayClient {
	return &PublicGatewayClient{
		sess,
	}
}

// List ...
func (f *PublicGatewayClient) List(start string) ([]*models.PublicGateway, string, error) {
	return f.ListWithFilter(start)
}

// ListWithFilter ...
func (f *PublicGatewayClient) ListWithFilter(start string) ([]*models.PublicGateway, string, error) {
	params := network.NewGetPublicGatewaysParamsWithTimeout(f.session.Timeout)
	if start != "" {
		params = params.WithStart(&start)
	}
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation

	resp, err := f.session.Riaas.Network.GetPublicGateways(params, session.Auth(f.session))

	if err != nil {
		return nil, "", errors.ToError(err)
	}

	return resp.Payload.PublicGateways, utils.GetNext(resp.Payload.Next), nil
}

// Get ...
func (f *PublicGatewayClient) Get(id string) (*models.PublicGateway, error) {
	params := network.NewGetPublicGatewaysIDParamsWithTimeout(f.session.Timeout).WithID(id)
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Network.GetPublicGatewaysID(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

/// Create ...
func (f *PublicGatewayClient) Create(name, zoneName, vpcID, FloatingIPID, FloatingIPaddr string) (*models.PublicGateway, error) {

	var body = network.PostPublicGatewaysBody{
		Name: name,
	}

	var zone = network.PostPublicGatewaysParamsBodyZone{
		Name: zoneName,
	}
	body.Zone = &zone

	targetvpc := strfmt.UUID(vpcID)
	var vpc = network.PostPublicGatewaysParamsBodyVpc{
		ID: targetvpc,
	}
	body.Vpc = &vpc
	if FloatingIPaddr != "" {
		var floatingip = network.PostPublicGatewaysParamsBodyFloatingIP{
			Address: FloatingIPaddr,
		}
		body.FloatingIP = &floatingip
	}
	if FloatingIPID != "" {
		var floatingip = network.PostPublicGatewaysParamsBodyFloatingIP{
			ID: strfmt.UUID(FloatingIPID),
		}
		body.FloatingIP = &floatingip
	}
	params := network.NewPostPublicGatewaysParamsWithTimeout(f.session.Timeout).WithBody(body)
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Network.PostPublicGateways(params, session.Auth(f.session))
	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// Delete ...
func (f *PublicGatewayClient) Delete(id string) error {
	params := network.NewDeletePublicGatewaysIDParamsWithTimeout(f.session.Timeout).WithID(id)
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation
	_, err := f.session.Riaas.Network.DeletePublicGatewaysID(params, session.Auth(f.session))
	return riaaserrors.ToError(err)
}

// Update ...
func (f *PublicGatewayClient) Update(id, name string) (*models.PublicGateway, error) {
	var body = network.PatchPublicGatewaysIDBody{
		Name: name,
	}
	params := network.NewPatchPublicGatewaysIDParamsWithTimeout(f.session.Timeout).WithID(id).WithBody(body)
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Network.PatchPublicGatewaysID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}
