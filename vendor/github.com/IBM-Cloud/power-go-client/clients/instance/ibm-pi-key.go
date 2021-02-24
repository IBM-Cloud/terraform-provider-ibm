package instance

import (
	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_tenants_ssh_keys"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/go-openapi/strfmt"
	"log"
	"time"
)

type IBMPIKeyClient struct {
	session         *ibmpisession.IBMPISession
	powerinstanceid string
}

func NewIBMPIKeyClient(sess *ibmpisession.IBMPISession, powerinstanceid string) *IBMPIKeyClient {
	return &IBMPIKeyClient{sess, powerinstanceid}
}

/*
This was a change requested by the IBM cloud Team to move the powerinstanceid out from the provider and pass it in the call
The Power-IAAS API requires the crn to be passed in the header.
*/
func (f *IBMPIKeyClient) Get(id, powerinstanceid string) (*models.SSHKey, error) {

	var tenantid = f.session.UserAccount
	log.Printf("Calling the Get code with the following params %s -  %s -  %s", id, powerinstanceid, tenantid)

	params := p_cloud_tenants_ssh_keys.NewPcloudTenantsSshkeysGetParams().WithTenantID(tenantid).WithSshkeyName(id)
	resp, err := f.session.Power.PCloudTenantsSSHKeys.PcloudTenantsSshkeysGet(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

func (f *IBMPIKeyClient) Create(name string, sshkey, powerinstanceid string) (*models.SSHKey, *models.SSHKey, error) {

	createDate := strfmt.DateTime(time.Now())
	var body = models.SSHKey{

		Name:         &name,
		SSHKey:       &sshkey,
		CreationDate: &createDate,
	}

	params := p_cloud_tenants_ssh_keys.NewPcloudTenantsSshkeysPostParamsWithTimeout(f.session.Timeout).WithTenantID(f.session.UserAccount).WithBody(&body)
	resp, postok, err := f.session.Power.PCloudTenantsSSHKeys.PcloudTenantsSshkeysPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return nil, nil, errors.ToError(err)
	}

	if resp != nil {
		log.Printf("Failed to get the key ")
	}
	if postok != nil {
		log.Print("Request failed ")
	}

	return nil, nil, nil

}

// Delete ...
func (f *IBMPIKeyClient) Delete(id string, powerinstanceid string) error {
	var tenantid = f.session.UserAccount
	params := p_cloud_tenants_ssh_keys.NewPcloudTenantsSshkeysDeleteParamsWithTimeout(f.session.Timeout).WithTenantID(tenantid).WithSshkeyName(id)
	_, err := f.session.Power.PCloudTenantsSSHKeys.PcloudTenantsSshkeysDelete(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}
