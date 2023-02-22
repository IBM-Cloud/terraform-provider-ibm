package secretsmanager

import (
	"fmt"
	"github.com/IBM/secrets-manager-go-sdk/secretsmanagerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"os"
	"strings"
)

// Clone the base secrets manager client and set the API endpoint per the instance
func getClientWithInstanceEndpoint(originalClient *secretsmanagerv2.SecretsManagerV2, d *schema.ResourceData) *secretsmanagerv2.SecretsManagerV2 {
	baseUrl := originalClient.Service.GetServiceURL()
	instanceId := d.Get("instance_id").(string)

	// find the region
	var region string
	_, ok := d.GetOk("region")
	if ok {
		region = d.Get("region").(string)
	} else {
		// extract region from base URL (provider config)
		// base url is like that : "https://<private.>secrets-manager.<region>.<rest of domain>"
		u := strings.Replace(baseUrl, "private.", "", 1)
		region = strings.Split(u, ".")[1]
	}

	log.Printf("[DEBUG] Secret Manager base URL: %s", baseUrl)

	// find the endpoint type
	var endpointType string
	_, ok = d.GetOk("endpoint_type")
	if ok {
		endpointType = d.Get("endpoint_type").(string)
		log.Printf("[DEBUG] Found endpoint type field")
	} else {
		if strings.Contains(baseUrl, "private.") {
			endpointType = "private"
		} else {
			endpointType = "public"
		}
	}

	// build the api endpoint
	domain := "appdomain.cloud"
	if strings.Contains(os.Getenv("IBMCLOUD_IAM_API_ENDPOINT"), "test") {
		domain = "test.appdomain.cloud"
	}
	var endpoint string
	if endpointType == "private" {
		endpoint = fmt.Sprintf("https://%s.private.%s.secrets-manager.%s/api", instanceId, region, domain)
	} else {
		endpoint = fmt.Sprintf("https://%s.%s.secrets-manager.%s/api", instanceId, region, domain)
	}

	// clone the client and set endpoint
	newClient := &secretsmanagerv2.SecretsManagerV2{
		Service: originalClient.Service.Clone(),
	}
	newClient.Service.SetServiceURL(endpoint)
	return newClient
}

// Add the fields needed for building the instance endpoint to the given schema
func AddInstanceFields(resource *schema.Resource) *schema.Resource {
	resource.Schema["instance_id"] = &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
		Description: "The ID of the Secrets Manager instance.",
	}
	resource.Schema["region"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		ForceNew:    true,
		Description: "The region of the Secrets Manager instance.",
	}
	resource.Schema["endpoint_type"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: "public or private.",
	}

	return resource
}
