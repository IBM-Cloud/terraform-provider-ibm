package compute

import (
	"github.com/go-openapi/strfmt"
	"github.ibm.com/Bluemix/riaas-go-client/riaas/client/compute"
	"github.ibm.com/Bluemix/riaas-go-client/riaas/models"

	"github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/Bluemix/riaas-go-client/session"
	"github.ibm.com/Bluemix/riaas-go-client/utils"
)

// KeyClient ...
type KeyClient struct {
	session *session.Session
}

// NewKeyClient ...
func NewKeyClient(sess *session.Session) *KeyClient {
	return &KeyClient{
		sess,
	}
}

// List ...
func (f *KeyClient) List(start string) ([]*models.Key, string, error) {
	return f.ListWithFilter("", start)
}

// ListWithFilter ...
func (f *KeyClient) ListWithFilter(resourceGroupID, start string) ([]*models.Key, string, error) {
	params := compute.NewGetKeysParamsWithTimeout(f.session.Timeout)
	if resourceGroupID != "" {
		params.WithResourceGroupID(&resourceGroupID)
	}
	if start != "" {
		params.WithStart(&start)
	}
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation

	resp, err := f.session.Riaas.Compute.GetKeys(params, session.Auth(f.session))

	if err != nil {
		return nil, "", errors.ToError(err)
	}

	return resp.Payload.Keys, utils.GetNext(resp.Payload.Next), nil
}

// Get ...
func (f *KeyClient) Get(id string) (*models.Key, error) {
	params := compute.NewGetKeysIDParamsWithTimeout(f.session.Timeout).WithID(id)
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Compute.GetKeysID(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Create ...
func (f *KeyClient) Create(name, keystring, resourcegroupID string) (*models.Key, error) {
	keytype := compute.PostKeysBodyTypeRsa
	var body = compute.PostKeysBody{
		Name:      name,
		PublicKey: &keystring,
		Type:      &keytype,
	}
	if resourcegroupID != "" {
		resourcegroupuuid := strfmt.UUID(resourcegroupID)
		var resourcegroup = compute.PostKeysParamsBodyResourceGroup{
			ID: resourcegroupuuid,
		}
		body.ResourceGroup = &resourcegroup
	}
	params := compute.NewPostKeysParamsWithTimeout(f.session.Timeout).WithBody(body)
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Compute.PostKeys(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Delete ...
func (f *KeyClient) Delete(id string) error {
	params := compute.NewDeleteKeysIDParamsWithTimeout(f.session.Timeout).WithID(id)
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation
	_, err := f.session.Riaas.Compute.DeleteKeysID(params, session.Auth(f.session))
	return err
}

// Update ...
func (f *KeyClient) Update(id, name string) (*models.Key, error) {
	var body = compute.PatchKeysIDBody{
		Name: name,
	}
	params := compute.NewPatchKeysIDParamsWithTimeout(f.session.Timeout).WithID(id).WithBody(body)
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Compute.PatchKeysID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}
