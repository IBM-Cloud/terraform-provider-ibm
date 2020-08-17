package instance

import (
	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_tasks"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"log"
)

type IBMPITaskClient struct {
	session         *ibmpisession.IBMPISession
	powerinstanceid string
}

// NewPowerImageClient ...
func NewIBMPITaskClient(sess *ibmpisession.IBMPISession, powerinstanceid string) *IBMPITaskClient {
	return &IBMPITaskClient{
		session:         sess,
		powerinstanceid: powerinstanceid,
	}
}

func (f *IBMPITaskClient) Get(id, powerinstanceid string) (*models.Task, error) {

	params := p_cloud_tasks.NewPcloudTasksGetParamsWithTimeout(postTimeOut).WithTaskID(id)
	resp, err := f.session.Power.PCloudTasks.PcloudTasksGet(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to get the task id ... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

func (f *IBMPITaskClient) Delete(id, powerinstanceid string) (models.Object, error) {

	params := p_cloud_tasks.NewPcloudTasksDeleteParamsWithTimeout(postTimeOut).WithTaskID(id)
	resp, err := f.session.Power.PCloudTasks.PcloudTasksDelete(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to delete the task id ... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}
