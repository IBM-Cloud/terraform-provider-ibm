package geography

import (
	"github.ibm.com/Bluemix/riaas-go-client/riaas/client/geography"
	"github.ibm.com/Bluemix/riaas-go-client/riaas/models"

	"github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/Bluemix/riaas-go-client/session"
)

// RegionClient ...
type RegionClient struct {
	session *session.Session
}

// NewRegionClient ...
func NewRegionClient(sess *session.Session) *RegionClient {
	return &RegionClient{
		sess,
	}
}

// List ...
func (f *RegionClient) List() ([]*models.Region, error) {
	params := geography.NewGetRegionsParamsWithTimeout(f.session.Timeout)
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Geography.GetRegions(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload.Regions, nil
}

// Get ...
func (f *RegionClient) Get(name string) (*models.Region, error) {
	params := geography.NewGetRegionsNameParamsWithTimeout(f.session.Timeout).WithName(name)
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Geography.GetRegionsName(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}
