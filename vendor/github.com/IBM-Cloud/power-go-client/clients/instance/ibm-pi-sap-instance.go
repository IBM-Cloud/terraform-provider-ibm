package instance

import (
	"fmt"
	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_s_a_p"
	"github.com/IBM-Cloud/power-go-client/power/models"

	"log"
)

type IBMPISAPInstanceClient struct {
	session         *ibmpisession.IBMPISession
	powerinstanceid string
}

// NewIBMPIInstanceClient ...
func NewIBMPISAPInstanceClient(sess *ibmpisession.IBMPISession, powerinstanceid string) *IBMPISAPInstanceClient {
	return &IBMPISAPInstanceClient{
		session:         sess,
		powerinstanceid: powerinstanceid,
	}
}

//Create SAP System
func (f *IBMPISAPInstanceClient) Create(sapdef *p_cloud_s_a_p.PcloudSapPostParams, id, powerinstanceid string) (*models.PVMInstanceList, error) {

	params := p_cloud_s_a_p.NewPcloudSapPostParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid).WithBody(sapdef.Body)
	sapok, sapcreated, sapaccepted, err := f.session.Power.PCloudSAP.PcloudSapPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		log.Printf("failed to process the request..")
		return nil, errors.ToError(err)
	}

	if sapok != nil && len(sapok.Payload) > 0 {
		log.Printf("Looks like we have an instance created....")
		log.Printf("Checking if the instance name is right ")
		log.Printf("Printing the instanceid %s", *sapok.Payload[0].PvmInstanceID)
		return &sapok.Payload, nil
	}
	if sapcreated != nil && len(sapcreated.Payload) > 0 {
		log.Printf("Printing the instanceid %s", *sapcreated.Payload[0].PvmInstanceID)
		return &sapcreated.Payload, nil
	}
	if sapaccepted != nil && len(sapaccepted.Payload) > 0 {

		log.Printf("Printing the instanceid %s", *sapaccepted.Payload[0].PvmInstanceID)
		return &sapaccepted.Payload, nil
	}

	//return &postok.Payload, nil
	return nil, fmt.Errorf("No response Returned ")
}
