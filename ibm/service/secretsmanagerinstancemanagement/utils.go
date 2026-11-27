package secretsmanagerinstancemanagement

import (
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
)

// Replaces the region part of an endpoint. It is assumed that endpoint is a valid https URL
// and the region is the first part of the host name
func ReplaceEndpointRegion(endpoint string, region string) string {
	schemeParts := strings.SplitN(endpoint, "://", 2)
	if len(schemeParts) != 2 { // shouldn't happen
		return endpoint
	}

	hostParts := strings.SplitN(schemeParts[1], ".", 2)
	if len(hostParts) != 2 { // shouldn't happen
		return endpoint
	}

	return schemeParts[0] + "://" + region + "." + hostParts[1]
}

func getInstanceManagementClient(d *schema.ResourceData, meta interface{}) (*secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2, error) {
	client, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		return nil, err
	}
	return getClientWithRegion(client, d), nil
}

// Clone the base SM instance management client and set the region specified in the resource / data source
func getClientWithRegion(originalClient *secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2,
	d *schema.ResourceData) *secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2 {
	_, ok := d.GetOk("region")
	if !ok {
		return originalClient // default region
	}

	region := d.Get("region").(string)

	endpoint := originalClient.Service.GetServiceURL()
	newEndpoint := ReplaceEndpointRegion(endpoint, region)
	// clone the client and set endpoint
	newClient := &secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2{
		Service: originalClient.Service.Clone(),
	}
	newClient.Service.SetServiceURL(newEndpoint)
	return newClient
}

func getRegion(originalClient *secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2, d *schema.ResourceData) string {
	_, ok := d.GetOk("region")
	if ok {
		return d.Get("region").(string)
	} else {
		// extract region from base URL (provider config)
		// base url is like that : "https://<private.>secrets-manager.<region>.<rest of domain>"
		baseUrl := originalClient.Service.GetServiceURL()
		schemeParts := strings.SplitN(baseUrl, "://", 2)
		if len(schemeParts) != 2 { // shouldn't happen
			return ""
		}

		hostParts := strings.SplitN(schemeParts[1], ".", 2)
		return hostParts[0]
	}
}
