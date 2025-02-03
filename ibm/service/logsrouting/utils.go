package logsrouting

import (
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/logs-router-go-sdk/ibmcloudlogsroutingv0"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Clones the logs routing client and sets the correct URL. Public, private, or custom
func updateClientURLWithEndpoint(logsRoutingClient *ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0, d *schema.ResourceData) (*ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0, string, error) {

	region := d.Get("region").(string)

	originalConfigServiceURL := logsRoutingClient.GetServiceURL()

	f := conns.EnvFallBack([]string{"IBMCLOUD_ENDPOINTS_FILE_PATH", "IC_ENDPOINTS_FILE_PATH"}, "")
	visibility := conns.EnvFallBack([]string{"IBMCLOUD_VISIBILITY", "IC_VISIBILITY"}, "")

	var logsrouterClientURL string
	var logsrouterURLErr error

	if f != "" && visibility != "public-and-private" {
		logsrouterClientURL = conns.FileFallBack(f, visibility, "IBMCLOUD_LOGS_ROUTING_API_ENDPOINT", region, ibmcloudlogsroutingv0.DefaultServiceURL)
	} else if visibility == "private" {
		logsrouterClientURL, logsrouterURLErr = ibmcloudlogsroutingv0.GetServiceURLForRegion("private." + region)
	} else {
		logsrouterClientURL, logsrouterURLErr = ibmcloudlogsroutingv0.GetServiceURLForRegion(region)
	}
	if logsrouterURLErr != nil {
		logsrouterClientURL = originalConfigServiceURL
	}

	newClient := &ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0{
		Service: logsRoutingClient.Service.Clone(),
	}

	log.Printf("Constructing client with new service URL %s", logsrouterClientURL)

	newClient.Service.SetServiceURL(logsrouterClientURL)

	return newClient, region, nil
}
