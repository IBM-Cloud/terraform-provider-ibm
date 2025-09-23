package appconfiguration

import (
	"fmt"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Wrapper function around  deprecated GetOkExists function with same functionality
func GetFieldExists(d *schema.ResourceData, field string) (any, bool) {
	return d.GetOkExists(field)
}

func getAppConfigClient(meta interface{}, guid string) (*appconfigurationv1.AppConfigurationV1, error) {
	bluemixSession, err := meta.(conns.ClientSession).BluemixSession()
	if err != nil {
		return nil, err
	}
	appConfigURL := fmt.Sprintf("https://%s.apprapp.cloud.ibm.com", bluemixSession.Config.Region)
	url := fmt.Sprintf("%s/apprapp/feature/v1/instances/%s", conns.EnvFallBack([]string{"IBMCLOUD_APP_CONFIG_API_ENDPOINT"}, appConfigURL), guid)
	appconfigClient, err := meta.(conns.ClientSession).AppConfigurationV1()
	if err != nil {
		return nil, err
	}
	appconfigClient.Service.Options.URL = url
	return appconfigClient, nil
}
