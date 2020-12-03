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
func (f *IBMPICloneVolumeClient) Create(cloneParams *p_cloud_volumes.PcloudV2VolumesClonePostParams, timeout time.Duration) (*models.CloneTaskReference, error) {

	log.Printf("Calling the P2 CloneVolume Create Method with provided time out value of [%f]", timeout.Minutes())
	log.Printf("The input clone name is %s and  to the cloudinstance id %s", cloneParams.Body.Name, cloneParams.CloudInstanceID)
	params := p_cloud_volumes.NewPcloudV2VolumesClonePostParamsWithTimeout(timeout).WithCloudInstanceID(cloneParams.CloudInstanceID).WithBody(cloneParams.Body)

	resp, err := f.session.Power.PCloudVolumes.PcloudV2VolumesClonePost(params, ibmpisession.NewAuth(f.session, cloneParams.CloudInstanceID))

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

func (f *IBMPICloneVolumeClient) Get(powerinstanceid, clonetaskid string, timeout time.Duration) (*models.CloneTaskStatus, error) {

	log.Printf("Calling the P2 CloneVolume Get Method with provided time out value of [%f]", timeout.Minutes())

	params := p_cloud_volumes.NewPcloudV2VolumesClonetasksGetParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithCloneTaskID(clonetaskid)

	resp, err := f.session.Power.PCloudVolumes.PcloudV2VolumesClonetasksGet(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the get operation for clones... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

//Start a clone
func (f *IBMPICloneVolumeClient) StartClone(powerinstanceid, volume_clone_id string, timeout time.Duration) (*models.VolumesClone, error) {

	log.Printf("Calling the P2 CloneVolume Start Method with provided time out value of [%f]", timeout.Minutes())

	params := p_cloud_volumes.NewPcloudV2VolumescloneStartPostParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithVolumesCloneID(volume_clone_id)

	resp, err := f.session.Power.PCloudVolumes.PcloudV2VolumescloneStartPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the start operation for clones... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}
