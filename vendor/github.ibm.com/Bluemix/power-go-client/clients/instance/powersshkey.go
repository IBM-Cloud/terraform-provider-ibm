package instance

import (
	"github.com/go-openapi/strfmt"
	"github.ibm.com/Bluemix/power-go-client/errors"
	"github.ibm.com/Bluemix/power-go-client/power/client/p_cloud_tenants_ssh_keys"
	"github.ibm.com/Bluemix/power-go-client/power/models"
	"github.ibm.com/Bluemix/power-go-client/session"
	"log"
	"time"
)

type PowerSSHKeyClient struct {
	session *session.Session
}

// NewPowerNetworkClient ...
func NewPowerSSHKeyClient(sess *session.Session) *PowerSSHKeyClient {
	return &PowerSSHKeyClient{
		sess,
	}
}

func (f *PowerSSHKeyClient) Get(id string) (*models.SSHKey, error) {

	//var cloudinstanceid = f.session.PowerServiceInstance

	var tenantid = f.session.UserAccount

	params := p_cloud_tenants_ssh_keys.NewPcloudTenantsSshkeysGetParams().WithTenantID(tenantid).WithSshkeyName(id)
	resp, err := f.session.Power.PCloudTenantsSSHKeys.PcloudTenantsSshkeysGet(params, session.NewAuth(f.session))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

func (f *PowerSSHKeyClient) Create(name string, sshkey string) (*models.SSHKey, *models.SSHKey, error) {

	createDate := strfmt.DateTime(time.Now())
	//var postokCreated = ""
	var body = models.SSHKey{

		Name:         &name,
		SSHKey:       &sshkey,
		CreationDate: &createDate,
	}

	params := p_cloud_tenants_ssh_keys.NewPcloudTenantsSshkeysPostParamsWithTimeout(f.session.Timeout).WithTenantID(f.session.UserAccount).WithBody(&body)
	resp, postok, err := f.session.Power.PCloudTenantsSSHKeys.PcloudTenantsSshkeysPost(params, session.NewAuth(f.session))
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
