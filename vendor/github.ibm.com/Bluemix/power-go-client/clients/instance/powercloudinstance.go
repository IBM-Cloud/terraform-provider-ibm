package instance

import (
	"github.ibm.com/Bluemix/power-go-client/errors"
	"github.ibm.com/Bluemix/power-go-client/power/client/p_cloud_instances"
	"github.ibm.com/Bluemix/power-go-client/power/models"
	"github.ibm.com/Bluemix/power-go-client/session"
	"log"
)

type PowerCloudInstanceClient struct {
	session *session.Session
}

// NewPowerVolumeClient ...
func NewPowerCloudInstanceClient(sess *session.Session) *PowerCloudInstanceClient {
	return &PowerCloudInstanceClient{
		sess,
	}
}

//Get information about a single volume only
func (f *PowerCloudInstanceClient) Get(id string) (*models.CloudInstance, error) {

	params := p_cloud_instances.NewPcloudCloudinstancesGetParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(id)
	log.Printf("Printing the data %s", params)
	//params := p_cloud_volumes.NewPcloudCloudinstancesVolumesGetParams().WithCloudInstanceID(cloudinstanceid).WithVolumeID(id)
	resp, err := f.session.Power.PCloudInstances.PcloudCloudinstancesGet(params, session.NewAuth(f.session))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

//Create
