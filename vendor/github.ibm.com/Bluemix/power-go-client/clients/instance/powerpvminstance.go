package instance

import (
	"github.ibm.com/Bluemix/power-go-client/errors"
	"github.ibm.com/Bluemix/power-go-client/power/client/p_cloud_p_vm_instances"
	"github.ibm.com/Bluemix/power-go-client/power/models"
	"github.ibm.com/Bluemix/power-go-client/session"

	"log"
)

type PowerPvmClient struct {
	session *session.Session
}

// NewPowerPvmClient ...
func NewPowerPvmClient(sess *session.Session) *PowerPvmClient {
	return &PowerPvmClient{
		sess,
	}
}

//Get information about a single pvm only
func (f *PowerPvmClient) Get(id string) (*models.PVMInstance, error) {

	params := p_cloud_p_vm_instances.NewPcloudPvminstancesGetParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(f.session.PowerServiceInstance).WithPvmInstanceID(id)
	resp, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesGet(params, session.NewAuth(f.session))
	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

//Create

func (f *PowerPvmClient) Create(powerdef *p_cloud_p_vm_instances.PcloudPvminstancesPostParams) (*models.PVMInstanceList, *models.PVMInstanceList, *models.PVMInstanceList, error) {

	log.Printf("Calling the Power PVM Create Method")
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesPostParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(f.session.PowerServiceInstance).WithBody(powerdef.Body)

	log.Printf("Printing the params to be passed %+v", params)

	_, _, resp, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesPost(params, session.NewAuth(f.session))

	if err != nil {
		log.Printf("failed to process the request..")
		return nil, nil, nil, errors.ToError(err)
	}

	return &resp.Payload, nil, nil, nil
}

func (f *PowerPvmClient) Delete(id string) error {

	log.Printf("Calling the Power PVM Delete Method")
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesDeleteParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(f.session.PowerServiceInstance).WithPvmInstanceID(id)
	_, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesDelete(params, session.NewAuth(f.session))

	if err != nil {
		return errors.ToError(err)
	}

	return nil
}

func (f *PowerPvmClient) Update(id string, powerupdateparams *p_cloud_p_vm_instances.PcloudPvminstancesPutParams) (*models.PVMInstanceUpdateResponse, error) {

	log.Printf("Calling the Power PVM Update Instance Method")
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesPutParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(f.session.PowerServiceInstance).WithPvmInstanceID(id).WithBody(powerupdateparams.Body)
	resp, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesPut(params, session.NewAuth(f.session))

	if err != nil {
		return nil, errors.ToError(err)

	}
	return resp.Payload, nil
}

func (f *PowerPvmClient) Action(id string, poweractionparams *p_cloud_p_vm_instances.PcloudPvminstancesActionPostParams) (models.Object, error) {

	log.Printf("Calling the Power PVM Action Method")
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesActionPostParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(f.session.PowerServiceInstance).WithPvmInstanceID(id)
	postok, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesActionPost(params, session.NewAuth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return postok.Payload, nil

}
