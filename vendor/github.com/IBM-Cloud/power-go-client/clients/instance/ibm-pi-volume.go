package instance

import (
	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_volumes"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"log"
)

type IBMPIVolumeClient struct {
	session         *ibmpisession.IBMPISession
	powerinstanceid string
}

// NewPowerVolumeClient ...
func NewIBMPIVolumeClient(sess *ibmpisession.IBMPISession, powerinstanceid string) *IBMPIVolumeClient {
	return &IBMPIVolumeClient{
		sess, powerinstanceid,
	}
}

//Get information about a single volume only
func (f *IBMPIVolumeClient) Get(id, powerinstanceid string) (*models.Volume, error) {

	log.Printf("Calling the VolumeGet Method..")
	log.Printf("The input volume name is %s and  to the cloudinstance id %s", id, powerinstanceid)

	params := p_cloud_volumes.NewPcloudCloudinstancesVolumesGetParams().WithCloudInstanceID(powerinstanceid).WithVolumeID(id)
	resp, err := f.session.Power.PCloudVolumes.PcloudCloudinstancesVolumesGet(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

//Create

func (f *IBMPIVolumeClient) Create(volumename string, volumesize float64, volumetype string, volumeshareable bool, powerinstanceid string) (*models.Volume, error) {

	log.Printf("calling the PowerVolume Create Method")

	var body = models.CreateDataVolume{
		Name:      &volumename,
		Size:      &volumesize,
		DiskType:  volumetype,
		Shareable: &volumeshareable,
	}

	params := p_cloud_volumes.NewPcloudCloudinstancesVolumesPostParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid).WithBody(&body)
	resp, err := f.session.Power.PCloudVolumes.PcloudCloudinstancesVolumesPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Delete ...
func (f *IBMPIVolumeClient) Delete(id string, powerinstanceid string) error {
	//var cloudinstanceid = f.session.PowerServiceInstance
	params := p_cloud_volumes.NewPcloudCloudinstancesVolumesDeleteParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid).WithVolumeID(id)
	_, err := f.session.Power.PCloudVolumes.PcloudCloudinstancesVolumesDelete(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// Update..
func (f *IBMPIVolumeClient) Update(id, volumename string, volumesize float64, volumeshare bool, powerinstanceid string) (*models.Volume, error) {

	var patchbody = models.UpdateVolume{}
	patchbody.Name = &volumename
	patchbody.Size = volumesize
	patchbody.Shareable = &volumeshare
	params := p_cloud_volumes.NewPcloudCloudinstancesVolumesPutParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid).WithVolumeID(id).WithBody(&patchbody)

	resp, err := f.session.Power.PCloudVolumes.PcloudCloudinstancesVolumesPut(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Attach a volume

func (f *IBMPIVolumeClient) Attach(id, volumename string, powerinstanceid string) (models.Object, error) {

	log.Printf("Calling the Power Volume Attach method")

	params := p_cloud_volumes.NewPcloudPvminstancesVolumesPostParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid).WithPvmInstanceID(id).WithVolumeID(volumename)
	resp, err := f.session.Power.PCloudVolumes.PcloudPvminstancesVolumesPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return nil, errors.ToError(err)
	}
	log.Printf("Successfully attached the volume to the instance")

	return resp.Payload, nil

}

//Detach a volume

func (f *IBMPIVolumeClient) Detach(id, volumename string, powerinstanceid string) (models.Object, error) {
	log.Printf("Calling the Power Volume Detach method")

	params := p_cloud_volumes.NewPcloudPvminstancesVolumesDeleteParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid).WithPvmInstanceID(id).WithVolumeID(volumename)
	resp, err := f.session.Power.PCloudVolumes.PcloudPvminstancesVolumesDelete(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil

}

// All volumes part of an instance

func (f *IBMPIVolumeClient) GetAll(id, cloud_instance_id string) (*models.Volumes, error) {

	log.Printf("Calling the Power Volumes GetAll Method")
	params := p_cloud_volumes.NewPcloudPvminstancesVolumesGetallParamsWithTimeout(f.session.Timeout).WithPvmInstanceID(id).WithCloudInstanceID(cloud_instance_id)
	resp, err := f.session.Power.PCloudVolumes.PcloudPvminstancesVolumesGetall(params, ibmpisession.NewAuth(f.session, cloud_instance_id))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil

}

// Set a volume as the boot volume - PUT Operation

func (f *IBMPIVolumeClient) SetBootVolume(id, volumename, cloud_instance_id string) (models.Object, error) {
	log.Printf("Setting the Boot Volume for this %s instance as ", cloud_instance_id)
	params := p_cloud_volumes.NewPcloudPvminstancesVolumesSetbootPutParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(cloud_instance_id).WithPvmInstanceID(id).WithVolumeID(volumename)
	resp, err := f.session.Power.PCloudVolumes.PcloudPvminstancesVolumesSetbootPut(params, ibmpisession.NewAuth(f.session, cloud_instance_id))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}
