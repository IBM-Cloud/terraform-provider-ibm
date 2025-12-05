package backuprecovery

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	session "github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// get instanceId and region
func getInstanceIdAndRegion(d *schema.ResourceData) (string, string) {

	var region string
	var instanceId string

	if _, ok := d.GetOk("instance_id"); ok {
		instanceId = d.Get("instance_id").(string)
	}

	if _, ok := d.GetOk("region"); ok {
		region = d.Get("region").(string)
	}

	return instanceId, region

}

// Clone the base backup recovery client and set the API endpoint per the instance
func getClientWithInstanceEndpoint(originalClient *backuprecoveryv1.BackupRecoveryV1, bmxsession *session.Session, instanceId, region, endpointType string) *backuprecoveryv1.BackupRecoveryV1 {
	// build the api endpoint

	// default endpoint_type is set to public
	if instanceId == "" && region == "" {
		return originalClient
	}

	domain := "cloud.ibm.com"
	serviceName := "backup-recovery"

	endpointsFile := bmxsession.Config.EndpointsFile

	iamUrl := os.Getenv("IBMCLOUD_IAM_API_ENDPOINT")
	if iamUrl == "" {
		iamUrl = conns.FileFallBack(endpointsFile, endpointType, "IBMCLOUD_IAM_API_ENDPOINT", region, "https://iam.cloud.ibm.com")
	}

	if strings.Contains(iamUrl, "test") {
		domain = "test.cloud.ibm.com"
	}

	var endpoint string
	if endpointType == "private" {
		endpoint = fmt.Sprintf("https://%s.private.%s.%s.%s/v2", instanceId, region, serviceName, domain)
	} else {
		endpoint = fmt.Sprintf("https://%s.%s.%s.%s/v2", instanceId, region, serviceName, domain)
	}

	// clone the client and set endpoint
	newClient := &backuprecoveryv1.BackupRecoveryV1{
		Service: originalClient.Service.Clone(),
	}
	newClient.Service.SetServiceURL(endpoint)
	return newClient
}

// Add the fields needed for building the instance endpoint to the given schema
func AddInstanceFields(resource *schema.Resource) *schema.Resource {
	resource.Schema["instance_id"] = &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		ForceNew:     true,
		Description:  "The instnace ID of the Backup Recovery instance.",
		RequiredWith: []string{"region"},
	}
	resource.Schema["region"] = &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		Computed:     true,
		ForceNew:     true,
		Description:  "The region of the Backup Recovery instance.",
		RequiredWith: []string{"instance_id"},
	}
	resource.Schema["endpoint_type"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: "public or private.",
		Default:     "public",
	}

	return resource
}

func BackupRecoverManagerEnvFallBack(endpointsFile, endpointType, region, str string) string {
	if v := os.Getenv(str); v != "" {
		return v
	}
	return fileFallBack(endpointsFile, endpointType, region, "IBMCLOUD_BACKUP_RECOVERY_MANAGER_API_KEY")

}

func fileFallBack(f, visibility, region, key string) string {
	fileMap := map[string]interface{}{}
	err := json.Unmarshal([]byte(f), &fileMap)
	if err != nil {
		return ""
	}
	if val, ok := fileMap[key]; ok {
		if v, ok := val.(map[string]interface{})[visibility]; ok {
			if r, ok := v.(map[string]interface{})[region]; ok && r.(string) != "" {
				return r.(string)
			}
		}
	}
	return ""
}

// Clone the base backup recovery client and set the API endpoint per the instance
func setManagerClientAuth(originalClient *backuprecoveryv1.BackupRecoveryManagementSreApiV1, bmxsession *session.Session, region, endpointType string) (*backuprecoveryv1.BackupRecoveryManagementSreApiV1, error) {
	// build the api endpoint

	endpointsFile := bmxsession.Config.EndpointsFile

	apiKey := BackupRecoverManagerEnvFallBack(endpointsFile, endpointType, region, "IBMCLOUD_BACKUP_RECOVERY_MANAGER_API_KEY")

	if apiKey == "" {
		err := fmt.Errorf("IBMCLOUD_BACKUP_RECOVERY_MANAGER_API_KEY not set in env or endpoints file")
		return nil, err
	}
	authconfig := &backuprecoveryv1.ManagementSreAuthenticatorConfig{
		ApiKey: apiKey,
	}
	authenticator, err := backuprecoveryv1.NewManagementSreAuthenticator(authconfig)
	if err != nil {
		return nil, err
	}

	// clone the client and set endpoint
	newClient := &backuprecoveryv1.BackupRecoveryManagementSreApiV1{
		Service: originalClient.Service.Clone(),
	}
	newClient.Service.Options.Authenticator = authenticator
	return newClient, nil
}

// Clone the base backup recovery client and set the API endpoint per the instance
func getManagerClientWithInstanceEndpoint(originalClient *backuprecoveryv1.BackupRecoveryManagementSreApiV1, bmxsession *session.Session, instanceId, region, endpointType string) *backuprecoveryv1.BackupRecoveryManagementSreApiV1 {
	// build the api endpoint

	// default endpoint_type is set to public
	if instanceId == "" {
		return originalClient
	}

	domain := "cloud.ibm.com"
	serviceName := "backup-recovery"
	endpointsFile := bmxsession.Config.EndpointsFile

	iamUrl := os.Getenv("IBMCLOUD_IAM_API_ENDPOINT")
	if iamUrl == "" {
		iamUrl = conns.FileFallBack(endpointsFile, endpointType, "IBMCLOUD_IAM_API_ENDPOINT", region, "https://iam.cloud.ibm.com")
	}
	if strings.Contains(iamUrl, "test") {
		domain = "test.cloud.ibm.com"
	}

	var endpoint string
	if endpointType == "private" {
		endpoint = fmt.Sprintf("https://%s.sre.private.%s.%s.%s/v2", instanceId, region, serviceName, domain)
	} else {
		endpoint = fmt.Sprintf("https://%s.sre.%s.%s.%s/v2", instanceId, region, serviceName, domain)
	}

	// clone the client and set endpoint
	newClient := &backuprecoveryv1.BackupRecoveryManagementSreApiV1{
		Service: originalClient.Service.Clone(),
	}
	newClient.Service.SetServiceURL(endpoint)
	return newClient
}
