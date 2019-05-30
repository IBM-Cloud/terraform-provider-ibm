package compute

import (
	"github.ibm.com/riaas/rias-api/riaas/client/compute"
	"github.ibm.com/riaas/rias-api/riaas/models"

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
	return f.ListWithFilter("", "", start)
}

// ListWithFilter ...
func (f *KeyClient) ListWithFilter(tag, resourceGroupID, start string) ([]*models.Key, string, error) {
	params := compute.NewGetKeysParams()
	if tag != "" {
		params.WithTag(&tag)
	}
	if resourceGroupID != "" {
		params.WithResourceGroupID(&resourceGroupID)
	}
	if start != "" {
		params.WithStart(&start)
	}
	params.Version = "2019-03-26"

	resp, err := f.session.Riaas.Compute.GetKeys(params, session.Auth(f.session))

	if err != nil {
		return nil, "", errors.ToError(err)
	}

	return resp.Payload.Keys, utils.GetPageLink(resp.Payload.Next), nil
}

// Get ...
func (f *KeyClient) Get(id string) (*models.Key, error) {
	params := compute.NewGetKeysIDParams().WithID(id)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Compute.GetKeysID(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Create ...
func (f *KeyClient) Create(name, keystring string) (*models.Key, error) {
	keytype := models.PostKeysParamsBodyTypeRsa
	var body = models.PostKeysParamsBody{
		Name:      name,
		PublicKey: keystring,
		Type:      &keytype,
	}
	params := compute.NewPostKeysParams().WithBody(&body)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Compute.PostKeys(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Delete ...
func (f *KeyClient) Delete(id string) error {
	params := compute.NewDeleteKeysIDParams().WithID(id)
	params.Version = "2019-03-26"
	_, err := f.session.Riaas.Compute.DeleteKeysID(params, session.Auth(f.session))
	return err
}

// Update ...
func (f *KeyClient) Update(id, name string) (*models.Key, error) {
	var body = models.PatchKeysIDParamsBody{
		Name: name,
	}
	params := compute.NewPatchKeysIDParams().WithID(id).WithBody(&body)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Compute.PatchKeysID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}
