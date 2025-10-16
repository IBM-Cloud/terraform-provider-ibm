package backuprecovery

import (
	"fmt"

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
func getClientWithInstanceEndpoint(originalClient *backuprecoveryv1.BackupRecoveryV1, instanceId, region, endpointType string) *backuprecoveryv1.BackupRecoveryV1 {
	// build the api endpoint

	// default endpoint_type is set to public
	if instanceId == "" && region == "" {
		return originalClient
	}

	domain := "cloud.ibm.com"
	serviceName := "backup-recovery"

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
