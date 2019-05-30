package geography

import (
	"github.ibm.com/riaas/rias-api/riaas/client/geography"
	"github.ibm.com/riaas/rias-api/riaas/models"

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
	params := geography.NewGetRegionsRegionNameZonesParams().WithRegionName(region)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Geography.GetRegionsRegionNameZones(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload.Zones, nil
}

// Get ...
func (f *ZoneClient) Get(region, name string) (*models.Zone, error) {
	params := geography.NewGetRegionsRegionNameZonesZoneNameParams().WithRegionName(region).WithZoneName(name)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.Geography.GetRegionsRegionNameZonesZoneName(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}
