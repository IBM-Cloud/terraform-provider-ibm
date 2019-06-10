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
	params.Version = "2019-03-26"
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
	params.Version = "2019-03-26"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Compute.GetImagesID(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}
