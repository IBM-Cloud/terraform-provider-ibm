package instance

import (
	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_system_pools"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"log"
)

type IBMPISystemPoolClient struct {
	session         *ibmpisession.IBMPISession
	powerinstanceid string
}

// NewSnapShotClient ...
func NewIBMPISystemPoolClient(sess *ibmpisession.IBMPISession, powerinstanceid string) *IBMPISystemPoolClient {
	return &IBMPISystemPoolClient{
		sess, powerinstanceid,
	}
}

//Get the System Pools
func (f *IBMPISystemPoolClient) Get(powerinstanceid string) (models.SystemPools, error) {

	log.Printf("Calling the System Pools Capacity Method..")
	log.Printf("The input cloudinstance id is %s", powerinstanceid)
	params := p_cloud_system_pools.NewPcloudSystempoolsGetParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid)

	resp, err := f.session.Power.PCloudSystemPools.PcloudSystempoolsGet(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}
