package instance

import (
	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_volumes"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"log"
	"time"
)

type IBMPICloneVolumeClient struct {
	session         *ibmpisession.IBMPISession
	powerinstanceid string
}

// NewSnapShotClient ...
func NewIBMPICloneVolumeClient(sess *ibmpisession.IBMPISession, powerinstanceid string) *IBMPICloneVolumeClient {
	return &IBMPICloneVolumeClient{
		sess, powerinstanceid,
	}
}

//Create a clone volume
func (f *IBMPICloneVolumeClient) Create(clone_params *p_cloud_volumes.PcloudVolumesClonePostParams, id, powerinstanceid string, timeout time.Duration) (*models.VolumesCloneResponse, error) {

	log.Printf("Calling the CloneVolume Create Method with provided time out value of [%f]", timeout.Minutes())
	log.Printf("The input clone name is %s and  to the cloudinstance id %s", id, powerinstanceid)
	params := p_cloud_volumes.NewPcloudVolumesClonePostParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithBody(clone_params.Body)

	resp, err := f.session.Power.PCloudVolumes.PcloudVolumesClonePost(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// Delete a volume that has been cloned
