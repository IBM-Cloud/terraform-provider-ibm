package instance

import (
	"fmt"
	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_p_vm_instances"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_s_a_p"
	"github.com/IBM-Cloud/power-go-client/power/models"

	"log"
)

type IBMPIInstanceClient struct {
	session         *ibmpisession.IBMPISession
	powerinstanceid string
}

// NewIBMPIInstanceClient ...
func NewIBMPIInstanceClient(sess *ibmpisession.IBMPISession, powerinstanceid string) *IBMPIInstanceClient {
	return &IBMPIInstanceClient{
		session:         sess,
		powerinstanceid: powerinstanceid,
	}
}

//Get information about a single pvm only
func (f *IBMPIInstanceClient) Get(id, powerinstanceid string) (*models.PVMInstance, error) {

	params := p_cloud_p_vm_instances.NewPcloudPvminstancesGetParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid).WithPvmInstanceID(id)
	resp, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesGet(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// Get Information about all the PVM Instances for a Client

func (f *IBMPIInstanceClient) GetAll(powerinstanceid string) (*models.PVMInstances, error) {

	params := p_cloud_p_vm_instances.NewPcloudPvminstancesGetallParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid)
	resp, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesGetall(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

//Create

func (f *IBMPIInstanceClient) Create(powerdef *p_cloud_p_vm_instances.PcloudPvminstancesPostParams, powerinstanceid string) (*models.PVMInstanceList, error) {

	log.Printf("Calling the Power PVM Create Method")
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesPostParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid).WithBody(powerdef.Body)

	log.Printf("Printing the params to be passed %+v", params)

	postok, postcreated, postAccepted, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil {
		log.Printf("failed to process the request..")
		return nil, errors.ToError(err)
	}

	if postok != nil && len(postok.Payload) > 0 {
		log.Printf("Looks like we have an instance created....")
		log.Printf("Checking if the instance name is right ")
		log.Printf("Printing the instanceid %s", *postok.Payload[0].PvmInstanceID)
		return &postok.Payload, nil
	}
	if postcreated != nil && len(postcreated.Payload) > 0 {
		log.Printf("Printing the instanceid %s", *postcreated.Payload[0].PvmInstanceID)
		return &postcreated.Payload, nil
	}
	if postAccepted != nil && len(postAccepted.Payload) > 0 {

		log.Printf("Printing the instanceid %s", *postAccepted.Payload[0].PvmInstanceID)
		return &postAccepted.Payload, nil
	}

	//return &postok.Payload, nil
	return nil, fmt.Errorf("No response Returned ")
}

// PVM Instances Delete
func (f *IBMPIInstanceClient) Delete(id, powerinstanceid string) error {

	log.Printf("Calling the Power PVM Delete Method")
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesDeleteParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid).WithPvmInstanceID(id)
	_, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesDelete(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil {
		return errors.ToError(err)
	}

	return nil
}

// PVM Instances Update
func (f *IBMPIInstanceClient) Update(id, powerinstanceid string, powerupdateparams *p_cloud_p_vm_instances.PcloudPvminstancesPutParams) (*models.PVMInstanceUpdateResponse, error) {

	log.Printf("Calling the Power PVM Update Instance Method")
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesPutParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid).WithPvmInstanceID(id).WithBody(powerupdateparams.Body)
	resp, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesPut(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil {
		return nil, errors.ToError(err)

	}
	return resp.Payload, nil
}

// PVM Instances Operations
func (f *IBMPIInstanceClient) Action(poweractionparams *p_cloud_p_vm_instances.PcloudPvminstancesActionPostParams, id, powerinstanceid string) (models.Object, error) {

	log.Printf("Calling the Power PVM Action Method")
	log.Printf("the params are %s - powerinstance id is %s", id, powerinstanceid)
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesActionPostParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid).WithPvmInstanceID(id).WithBody(poweractionparams.Body)

	log.Printf("printing the poweraction params %+v", params)

	postok, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesActionPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return postok.Payload, nil

}

// Generate the Console URL

func (f *IBMPIInstanceClient) PostConsoleURL(id, powerinstanceid string) (models.Object, error) {
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesConsolePostParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid).WithPvmInstanceID(id)

	postok, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesConsolePost(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil {
		return nil, errors.ToError(err)
	}
	return postok.Payload, nil
}

// Capture an instance

func (f *IBMPIInstanceClient) CaptureInstanceToImageCatalog(id, powerinstanceid string, picaptureparams *p_cloud_p_vm_instances.PcloudPvminstancesCapturePostParams) (models.Object, error) {

	params := p_cloud_p_vm_instances.NewPcloudPvminstancesCapturePostParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid).WithPvmInstanceID(id).WithBody(picaptureparams.Body)
	postok, _, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesCapturePost(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return postok.Payload, nil

}

// Create a snapshot of the instance
func (f *IBMPIInstanceClient) CreatePvmSnapShot(snapshotdef *p_cloud_p_vm_instances.PcloudPvminstancesSnapshotsPostParams, pvminstanceid, powerinstanceid string) (*models.SnapshotCreateResponse, error) {
	log.Printf("Calling the Power PVM Snaphshot Method and printing the following data %s - %s - %s", snapshotdef.Body.Description, snapshotdef.Body.Name, snapshotdef.Body.VolumeIds)
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesSnapshotsPostParamsWithTimeout(f.session.Timeout).WithPvmInstanceID(pvminstanceid).WithCloudInstanceID(powerinstanceid).WithBody(snapshotdef.Body)
	snapshotpostaccepted, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesSnapshotsPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	log.Printf("*** Printing the error from the snapshot call *** %s", err)
	if err != nil {

		log.Printf("Printing the snapshot data %s", snapshotpostaccepted.Payload)
		return nil, fmt.Errorf("Failed to Create the snapshot for the pvminstance [%s]", pvminstanceid)

	}
	log.Printf("Successfully executed the snapshot for the pvminstance [%s] with status of [%s]", pvminstanceid, snapshotpostaccepted.Payload.SnapshotID)

	return snapshotpostaccepted.Payload, nil
}

// Create a clone

func (f *IBMPIInstanceClient) CreateClone(clonedef *p_cloud_p_vm_instances.PcloudPvminstancesClonePostParams, pvminstanceid, powerinstanceid string) (*models.PVMInstance, error) {
	log.Printf("Calling the Power PVM Clone Method")
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesClonePostParamsWithTimeout(f.session.Timeout).WithPvmInstanceID(pvminstanceid).WithCloudInstanceID(powerinstanceid).WithBody(clonedef.Body)
	clonePost, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesClonePost(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {

		return nil, fmt.Errorf("Failed to create the clone of the pvm instance")
	}
	return clonePost.Payload, nil
}

// Get information about the snapshots for a vm

func (f *IBMPIInstanceClient) GetSnapShotVM(powerinstanceid, pvminstanceid string) (*models.Snapshots, error) {

	log.Printf("Calling the GetSnapshot for vm Method..")
	log.Printf("The input pvmid name is %s and  to the cloudinstance id %s", pvminstanceid, powerinstanceid)

	params := p_cloud_p_vm_instances.NewPcloudPvminstancesSnapshotsGetallParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid).WithPvmInstanceID(pvminstanceid)
	resp, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesSnapshotsGetall(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil

}

// Restore a snapshot

func (f *IBMPIInstanceClient) RestoreSnapShotVM(powerinstanceid, pvminstanceid, snapshotid, restoreAction string, restoreparams *p_cloud_p_vm_instances.PcloudPvminstancesSnapshotsRestorePostParams) (*models.Snapshot, error) {
	log.Printf("Performing the snapshot restore for lpar with instanceid [%s] and snapshotid [%s] ", pvminstanceid, snapshotid)
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesSnapshotsRestorePostParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid).WithPvmInstanceID(pvminstanceid).WithSnapshotID(snapshotid).WithRestoreFailAction(&restoreAction).WithBody(restoreparams.Body)
	resp, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesSnapshotsRestorePost(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil || resp.Payload.SnapshotID != nil {
		log.Printf("Failed to perform the snapshot restore operation")
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// Create SAP Systems

func (f *IBMPIInstanceClient) CreateSAP(powerdef *p_cloud_s_a_p.PcloudSapPostParams, powerinstanceid string) (*models.PVMInstanceList, error) {

	log.Printf("Calling the Power PVM Create Method For SAP")
	params := p_cloud_s_a_p.NewPcloudSapPostParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid).WithBody(powerdef.Body)

	log.Printf("Printing the params to be passed %+v", params)

	postok, postcreated, postAccepted, err := f.session.Power.PCloudSAP.PcloudSapPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil {
		log.Printf("failed to process the request..")
		return nil, errors.ToError(err)
	}

	if postok != nil && len(postok.Payload) > 0 {
		log.Printf("Looks like we have an instance created....")
		log.Printf("Checking if the instance name is right ")
		log.Printf("Printing the instanceid for SAP %s", *postok.Payload[0].PvmInstanceID)
		return &postok.Payload, nil
	}
	if postcreated != nil && len(postcreated.Payload) > 0 {
		log.Printf("Printing the instanceid %s", *postcreated.Payload[0].PvmInstanceID)
		return &postcreated.Payload, nil
	}
	if postAccepted != nil && len(postAccepted.Payload) > 0 {

		log.Printf("Printing the instanceid %s", *postAccepted.Payload[0].PvmInstanceID)
		return &postAccepted.Payload, nil
	}

	//return &postok.Payload, nil
	return nil, fmt.Errorf("No response Returned ")
}

// Get All SAP Profiles

func (f *IBMPIInstanceClient) GetSAPProfiles(powerinstanceid string) (*models.SAPProfiles, error) {

	params := p_cloud_s_a_p.NewPcloudSapGetallParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid)
	resp, err := f.session.Power.PCloudSAP.PcloudSapGetall(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}
