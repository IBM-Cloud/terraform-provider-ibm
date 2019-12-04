package compute

import (
	"github.ibm.com/Bluemix/riaas-go-client/riaas/client/compute"
	"github.ibm.com/Bluemix/riaas-go-client/riaas/models"

	"github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/Bluemix/riaas-go-client/session"
	"github.ibm.com/Bluemix/riaas-go-client/utils"
)

// ImageClient ...
type ImageClient struct {
	session *session.Session
}

// NewImageClient ...
func NewImageClient(sess *session.Session) *ImageClient {
	return &ImageClient{
		sess,
	}
}

// ListWithFilter ...
func (f *ImageClient) ListWithFilter(visibility, start string) ([]*models.Image, string, error) {
	params := compute.NewGetImagesParamsWithTimeout(f.session.Timeout)

	if visibility != "" {
		params = params.WithVisibility(&visibility)
	}

	if start != "" {
		params = params.WithStart(&start)
	}
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation

	resp, err := f.session.Riaas.Compute.GetImages(params, session.Auth(f.session))

	if err != nil {
		return nil, "", errors.ToError(err)
	}

	return resp.Payload.Images, utils.GetNext(resp.Payload.Next), nil
}

// List ...
func (f *ImageClient) List(start string) ([]*models.Image, string, error) {
	return f.ListWithFilter("", start)
}

// Get ...
func (f *ImageClient) Get(id string) (*models.Image, error) {
	params := compute.NewGetImagesIDParamsWithTimeout(f.session.Timeout).WithID(id)
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Compute.GetImagesID(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Create ...
func (f *ImageClient) Create(href, name, operatingSystem string) (*models.Image, error) {
	var operatingSystemIdentity = models.OperatingSystemIdentity{
		Name: &operatingSystem,
	}
	var imageFileTemplate = models.ImageFileTemplate{
		Href: &href,
	}
	var imageTemplate = models.ImageTemplate{
		File:            &imageFileTemplate,
		Name:            name,
		OperatingSystem: &operatingSystemIdentity,
	}

	params := compute.NewPostImagesParamsWithTimeout(f.session.Timeout).WithBody(&imageTemplate)
	params.Version = "2019-11-22"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Compute.PostImages(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Delete ...
func (f *ImageClient) Delete(id string) error {
	params := compute.NewDeleteImagesIDParamsWithTimeout(f.session.Timeout).WithID(id)
	params.Version = "2019-11-22"
	params.Generation = f.session.Generation
	_, err := f.session.Riaas.Compute.DeleteImagesID(params, session.Auth(f.session))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// Update ...
func (f *ImageClient) Update(id, name string) (*models.Image, error) {
	var imagePatch = models.ImagePatch{
		Name: name,
	}
	params := compute.NewPatchImagesIDParamsWithTimeout(f.session.Timeout).WithID(id).WithRequestBody(&imagePatch)
	params.Version = "2019-11-22"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Compute.PatchImagesID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}
