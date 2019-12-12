package geography

import (
	"github.ibm.com/Bluemix/riaas-go-client/riaas/client/geography"
	"github.ibm.com/Bluemix/riaas-go-client/riaas/models"

	"github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/Bluemix/riaas-go-client/session"
)

// ZoneClient ...
type ZoneClient struct {
	session *session.Session
}

// NewZoneClient ...
func NewZoneClient(sess *session.Session) *ZoneClient {
	return &ZoneClient{
		sess,
	}
}

// List ..
func (f *ZoneClient) List(region string) ([]*models.Zone, error) {
	params := geography.NewGetRegionsRegionNameZonesParamsWithTimeout(f.session.Timeout).WithRegionName(region)
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Geography.GetRegionsRegionNameZones(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload.Zones, nil
}

// Get ...
func (f *ZoneClient) Get(region, name string) (*models.Zone, error) {
	params := geography.NewGetRegionsRegionNameZonesZoneNameParamsWithTimeout(f.session.Timeout).WithRegionName(region).WithZoneName(name)
	params.Version = "2019-10-08"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Geography.GetRegionsRegionNameZonesZoneName(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}
