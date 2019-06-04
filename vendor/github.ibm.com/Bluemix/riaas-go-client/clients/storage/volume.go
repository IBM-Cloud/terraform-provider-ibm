package storage

import (
	"github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/Bluemix/riaas-go-client/session"
	"github.ibm.com/riaas/rias-api/riaas/client/storage"
	"github.ibm.com/riaas/rias-api/riaas/models"
)

// StorageClient ...
type StorageClient struct {
	session *session.Session
}

// NewStorageClient ...
func NewStorageClient(sess *session.Session) *StorageClient {
	return &StorageClient{
		sess,
	}
}

// Create ...StorageClient
func (f *StorageClient) Create(storagedef *storage.PostVolumesParams) (*models.Volume, error) {
	params := storage.NewPostVolumesParams().WithBody(storagedef.Body)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Storage.PostVolumes(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Delete ...
func (f *StorageClient) Delete(id string) error {
	params := storage.NewDeleteVolumesIDParams().WithID(id)
	params.Version = "2019-03-26"
	_, err := f.session.Riaas.Storage.DeleteVolumesID(params, session.Auth(f.session))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// Get ...
func (f *StorageClient) Get(id string) (*models.Volume, error) {
	params := storage.NewGetVolumesIDParams().WithID(id)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Storage.GetVolumesID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Update ...
func (f *StorageClient) Update(id string, patchparms *storage.PatchVolumesIDParams) (*models.Volume, error) {
	params := storage.NewPatchVolumesIDParams().WithID(id).WithBody(patchparms.Body)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Storage.PatchVolumesID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}
