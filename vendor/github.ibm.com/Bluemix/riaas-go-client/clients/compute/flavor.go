package compute

import (
	"github.ibm.com/riaas/rias-api/riaas/client/compute"
	"github.ibm.com/riaas/rias-api/riaas/models"

	"github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/Bluemix/riaas-go-client/session"
	"github.ibm.com/Bluemix/riaas-go-client/utils"
)

// FlavorClient ...
type FlavorClient struct {
	session *session.Session
}

// NewFlavorClient ...
func NewFlavorClient(sess *session.Session) *FlavorClient {
	return &FlavorClient{
		sess,
	}
}

// List ...
func (f *FlavorClient) List(start string) ([]*models.Flavor, string, error) {
	params := compute.NewGetFlavorsParams()
	if start != "" {
		params = params.WithStart(&start)
	}
	resp, err := f.session.Riaas.Compute.GetFlavors(params, session.Auth(f.session))

	if err != nil {
		return nil, "", errors.ToError(err)
	}
	return resp.Payload.Flavors, utils.GetPageLink(resp.Payload.Next), nil
}

// Get ...
func (f *FlavorClient) Get(name string) (*models.Flavor, error) {
	params := compute.NewGetFlavorsIDParams().WithID(name)
	resp, err := f.session.Riaas.Compute.GetFlavorsID(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}
