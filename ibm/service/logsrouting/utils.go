package logsrouting

import (
	"log"
	"strings"

	"github.com/IBM/logs-router-go-sdk/ibmcloudlogsroutingv0"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Clones the logs routing client and sets the correct URL. Public, private, or custom
func updateClientURLWithEndpoint(logsRoutingClient *ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0, d *schema.ResourceData) (*ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0, string, error) {

	region := d.Get("region").(string)
	originalConfigServiceURL := logsRoutingClient.GetServiceURL()
	newServiceURL := replaceRegion(originalConfigServiceURL, region)

	newClient := &ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0{
		Service: logsRoutingClient.Service.Clone(),
	}

	log.Printf("Constructing client with new service URL %s", newServiceURL)

	newClient.Service.SetServiceURL(newServiceURL)

	return newClient, region, nil
}

// Function to replace the region in the URL
func replaceRegion(url, region string) string {
	// Split the URL by "." to isolate the relevant parts
	parts := strings.Split(url, ".")

	// Check if the URL contains 'private'
	if len(parts) > 1 && parts[1] == "private" {
		parts[2] = region
	} else {
		parts[1] = region
	}

	// Join the parts back into a complete URL
	return strings.Join(parts, ".")
}
