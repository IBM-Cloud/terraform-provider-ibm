package instance

import (
	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_images"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"log"
)

type IBMPIImageClient struct {
	session         *ibmpisession.IBMPISession
	powerinstanceid string
}

// NewPowerImageClient ...
func NewIBMPIImageClient(sess *ibmpisession.IBMPISession, powerinstanceid string) *IBMPIImageClient {
	return &IBMPIImageClient{
		session:         sess,
		powerinstanceid: powerinstanceid,
	}
}

func (f *IBMPIImageClient) Get(id, powerinstanceid string) (*models.Image, error) {

	params := p_cloud_images.NewPcloudCloudinstancesImagesGetParams().WithCloudInstanceID(powerinstanceid).WithImageID(id)
	resp, err := f.session.Power.PCloudImages.PcloudCloudinstancesImagesGet(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

func (f *IBMPIImageClient) GetAll(powerinstanceid string) (*models.Images, error) {

	params := p_cloud_images.NewPcloudImagesGetallParams()
	resp, err := f.session.Power.PCloudImages.PcloudImagesGetall(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil

}
