package geography

import (
	"github.ibm.com/riaas/rias-api/riaas/client/geography"
	"github.ibm.com/riaas/rias-api/riaas/models"

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
	params := geography.NewGetRegionsParams()
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Geography.GetRegions(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload.Regions, nil
}

// Get ...
func (f *RegionClient) Get(name string) (*models.Region, error) {
	params := geography.NewGetRegionsNameParams().WithName(name)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Geography.GetRegionsName(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}
