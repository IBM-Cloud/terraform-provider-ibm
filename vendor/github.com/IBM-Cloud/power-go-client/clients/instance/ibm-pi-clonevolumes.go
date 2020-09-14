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

//Create a clone volume using V2 of the API - This creates a clone
func (f *IBMPICloneVolumeClient) Create(cloneParams *p_cloud_volumes.PcloudV2VolumesClonePostParams, id, cloudinstance string, timeout time.Duration) (*models.CloneTaskReference, error) {

	log.Printf("Calling the P2 CloneVolume Create Method with provided time out value of [%f]", timeout.Minutes())
	log.Printf("The input clone name is %s and  to the cloudinstance id %s", id, cloudinstance)
	params := p_cloud_volumes.NewPcloudV2VolumesClonePostParamsWithTimeout(timeout).WithCloudInstanceID(cloudinstance).WithBody(cloneParams.Body)

	resp, err := f.session.Power.PCloudVolumes.PcloudV2VolumesClonePost(params, ibmpisession.NewAuth(f.session, cloudinstance))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

//Delete a clone
func (f *IBMPICloneVolumeClient) DeleteClone(cloneParams *p_cloud_volumes.PcloudV2VolumescloneDeleteParams, id, cloudinstance string, timeout time.Duration) (models.Object, error) {

	log.Printf("Calling the P2 CloneVolume Delele Method with provided time out value of [%f]", timeout.Minutes())
	log.Printf("The input clone name is %s and  to the cloudinstance id %s", id, cloudinstance)
	params := p_cloud_volumes.NewPcloudV2VolumescloneDeleteParamsWithTimeout(timeout).WithCloudInstanceID(cloudinstance).WithVolumesCloneID(id)

	resp, err := f.session.Power.PCloudVolumes.PcloudV2VolumescloneDelete(params, ibmpisession.NewAuth(f.session, cloudinstance))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// Cancel a Clone

// Get status of a clone request
/* PcloudV2VolumesClonetasksGet gets the status of a volumes clone request for the specified clone task ID -

This is from the post to start a clone - it returns a clone_task_id which will be used to query
*/
