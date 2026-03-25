package logsrouting

import (
	"fmt"
	"log"
	"strings"

	bxsession "github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/logs-router-go-sdk/ibmcloudlogsroutingv0"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Clones the logs routing client and sets the correct URL. Public, private, or custom
func updateClientURLWithEndpoint(logsRoutingClient *ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0, d *schema.ResourceData, sess *bxsession.Session) (*ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0, string, error) {

	var newServiceURL string
	originalConfigServiceURL := logsRoutingClient.GetServiceURL()
	endpointsFile := sess.Config.EndpointsFile
	visibility := sess.Config.Visibility
	privateEndpointType := sess.Config.PrivateEndpointType
	region := d.Get("region").(string)

	log.Printf("[DEBUG] Logs Routing Visibility:: %s, PrivateEndpointType: %s, Region: %s", visibility, privateEndpointType, region)

	if endpointsFile != "" && visibility != "public-and-private" {
		newServiceURL = conns.FileFallBack(endpointsFile, visibility, "IBMCLOUD_LOGS_ROUTING_API_ENDPOINT", region, originalConfigServiceURL)
	} else {
		newServiceURL = buildEndpointURL(originalConfigServiceURL, region, visibility, privateEndpointType)
	}

	newClient := &ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0{
		Service: logsRoutingClient.Service.Clone(),
	}

	log.Printf("[DEBUG] Constructing client with new service URL %s", newServiceURL)

	newClient.Service.SetServiceURL(newServiceURL)

	return newClient, region, nil
}

func buildEndpointURL(originalURL, region, visibility, privateEndpointType string) string {
	// Default endpoint format: https://management.<region>.logs-router.cloud.ibm.com/v1
	// Private VPE format: https://management.private.<region>.logs-router.cloud.ibm.com/v1
	// Private CSE format: https://management.private.<region>.logs-router.cloud.ibm.com:3443/v1

	if visibility == "private" {
		if privateEndpointType == "vpe" {
			return fmt.Sprintf("https://management.private.%s.logs-router.cloud.ibm.com/v1", region)
		} else {
			return fmt.Sprintf("https://management.private.%s.logs-router.cloud.ibm.com:3443/v1", region)
		}
	} else if visibility == "public-and-private" {
		return replaceRegion(originalURL, region)
	} else {
		return fmt.Sprintf("https://management.%s.logs-router.cloud.ibm.com/v1", region)
	}
}

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
