package network

import (
	"github.com/go-openapi/strfmt"
	"github.ibm.com/Bluemix/riaas-go-client/errors"
	riaaserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/Bluemix/riaas-go-client/session"
	"github.ibm.com/Bluemix/riaas-go-client/utils"
	"github.ibm.com/riaas/rias-api/riaas/client/network"
	"github.ibm.com/riaas/rias-api/riaas/models"
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
	return f.ListWithFilter(nil, start)
}

// ListWithFilter ...
func (f *PublicGatewayClient) ListWithFilter(tag *string, start string) ([]*models.PublicGateway, string, error) {
	params := network.NewGetPublicGatewaysParams()
	if tag != nil {
		params = params.WithTag(tag)
	}
	if start != "" {
		params = params.WithStart(&start)
	}
	params.Version = "2019-03-26"

	resp, err := f.session.Riaas.Network.GetPublicGateways(params, session.Auth(f.session))

	if err != nil {
		return nil, "", errors.ToError(err)
	}

	return resp.Payload.PublicGateways, utils.GetNext(resp.Payload.Next), nil
}

// Get ...
func (f *PublicGatewayClient) Get(id string) (*models.PublicGateway, error) {
	params := network.NewGetPublicGatewaysIDParams().WithID(id)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Network.GetPublicGatewaysID(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

/// Create ...
func (f *PublicGatewayClient) Create(name, zoneName, vpcID, FloatingIPaddr string) (*models.PublicGateway, error) {

	var body = models.PostPublicGatewaysParamsBody{
		Name: name,
	}

	var zone = models.NameReference{
		Name: zoneName,
	}
	body.Zone = &zone

	targetvpc := strfmt.UUID(vpcID)
	var vpc = models.PostPublicGatewaysParamsBodyVpc{
		ID: targetvpc,
	}
	body.Vpc = &vpc
	if FloatingIPaddr != "" {
		var floatingip = models.PostPublicGatewaysParamsBodyFloatingIP{
			Address: FloatingIPaddr,
		}
		body.FloatingIP = &floatingip
	}

	params := network.NewPostPublicGatewaysParams().WithBody(&body)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Network.PostPublicGateways(params, session.Auth(f.session))
	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// Delete ...
func (f *PublicGatewayClient) Delete(id string) error {
	params := network.NewDeletePublicGatewaysIDParams().WithID(id)
	params.Version = "2019-03-26"
	_, err := f.session.Riaas.Network.DeletePublicGatewaysID(params, session.Auth(f.session))
	return riaaserrors.ToError(err)
}

// Update ...
func (f *PublicGatewayClient) Update(id, name string) (*models.PublicGateway, error) {
	var body = models.PatchPublicGatewaysIDParamsBody{
		Name: name,
	}
	params := network.NewPatchPublicGatewaysIDParams().WithID(id).WithBody(&body)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Network.PatchPublicGatewaysID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}
