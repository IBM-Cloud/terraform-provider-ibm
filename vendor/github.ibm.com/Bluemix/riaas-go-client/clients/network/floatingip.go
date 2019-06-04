package network

import (
	"github.ibm.com/riaas/rias-api/riaas/client/network"
	"github.ibm.com/riaas/rias-api/riaas/models"

	"github.com/go-openapi/strfmt"
	riaaserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/Bluemix/riaas-go-client/session"
	"github.ibm.com/Bluemix/riaas-go-client/utils"
)

// FloatingIPClient ...
type FloatingIPClient struct {
	session *session.Session
}

// NewFloatingIPClient ...
func NewFloatingIPClient(sess *session.Session) *FloatingIPClient {
	return &FloatingIPClient{
		sess,
	}
}

// List ...
func (f *FloatingIPClient) List(start string) ([]*models.FloatingIP, string, error) {
	return f.ListWithFilter("", "", "", start)
}

// ListWithFilter ...
func (f *FloatingIPClient) ListWithFilter(tag, zoneName, resourcegroupID, start string) ([]*models.FloatingIP, string, error) {
	params := network.NewGetFloatingIpsParams()
	if tag != "" {
		params = params.WithTag(&tag)
	}
	if zoneName != "" {
		params = params.WithZoneName(&zoneName)
	}
	if resourcegroupID != "" {
		params = params.WithResourceGroupID(&resourcegroupID)
	}
	if start != "" {
		params = params.WithStart(&start)
	}
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Network.GetFloatingIps(params, session.Auth(f.session))

	if err != nil {
		return nil, "", riaaserrors.ToError(err)
	}

	return resp.Payload.FloatingIps, utils.GetNext(resp.Payload.Next), nil
}

// Get ...
func (f *FloatingIPClient) Get(id string) (*models.FloatingIP, error) {
	params := network.NewGetFloatingIpsIDParams().WithID(id)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Network.GetFloatingIpsID(params, session.Auth(f.session))

	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// Create ...
func (f *FloatingIPClient) Create(name, zoneName, resourcegroupID, targetID string) (*models.FloatingIP, error) {

	var body = models.PostFloatingIpsParamsBody{
		Name: name,
	}

	if zoneName != "" {
		var zone = models.NameReference{
			Name: zoneName,
		}
		body.Zone = &zone
	}

	if targetID != "" {
		targetUUID := strfmt.UUID(targetID)
		var target = models.PostFloatingIpsParamsBodyTarget{
			ID: targetUUID,
		}
		body.Target = &target
	}

	if resourcegroupID != "" {
		resourcegroupuuid := strfmt.UUID(resourcegroupID)
		var resourcegroup = models.PostFloatingIpsParamsBodyResourceGroup{
			ID: resourcegroupuuid,
		}
		body.ResourceGroup = &resourcegroup
	}

	params := network.NewPostFloatingIpsParams().WithBody(&body)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Network.PostFloatingIps(params, session.Auth(f.session))
	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// Delete ...
func (f *FloatingIPClient) Delete(id string) error {
	params := network.NewDeleteFloatingIpsIDParams().WithID(id)
	params.Version = "2019-03-26"
	_, err := f.session.Riaas.Network.DeleteFloatingIpsID(params, session.Auth(f.session))
	return riaaserrors.ToError(err)
}

// Update ...
func (f *FloatingIPClient) Update(id, name, targetID string) (*models.FloatingIP, error) {
	var body = models.PatchFloatingIpsIDParamsBody{}

	if name != "" {
		body.Name = name
	}

	if targetID != "" {
		targetUUID := strfmt.UUID(targetID)
		var target = models.PatchFloatingIpsIDParamsBodyTarget{
			ID: targetUUID,
		}
		body.Target = &target
	}

	params := network.NewPatchFloatingIpsIDParams().WithID(id).WithBody(&body)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Network.PatchFloatingIpsID(params, session.Auth(f.session))
	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}
