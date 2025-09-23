package logs_test

import (
	"fmt"
	"os"
	"strings"

	"github.com/IBM/logs-go-sdk/logsv0"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
)

const (
	cloudEndpoint     = "cloud.ibm.com"
	testCloudEndpoint = "test.cloud.ibm.com"
)

// <instance_id>.api.eu-gb.logs.test.cloud.ibm.com
// Clone the base logs client and set the API endpoint per the instance
func getTestClientWithLogsInstanceEndpoint(originalClient *logsv0.LogsV0) *logsv0.LogsV0 {
	// build the api endpoint
	domain := cloudEndpoint
	if strings.Contains(os.Getenv("IBMCLOUD_IAM_API_ENDPOINT"), "test") {
		domain = testCloudEndpoint
	}
	var endpoint string

	endpoint = fmt.Sprintf("https://%s.api.%s.logs.%s", acc.LogsInstanceId, acc.LogsInstanceRegion, domain)

	// clone the client and set endpoint
	newClient := &logsv0.LogsV0{
		Service: originalClient.Service.Clone(),
	}

	endpoint = conns.EnvFallBack([]string{"IBMCLOUD_LOGS_API_ENDPOINT"}, endpoint)

	newClient.Service.SetServiceURL(endpoint)

	return newClient
}
