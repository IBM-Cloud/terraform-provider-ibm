package appconfiguration

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.in/yaml.v3"
)

// Wrapper function around  deprecated GetOkExists function with same functionality
func GetFieldExists(d *schema.ResourceData, field string) (any, bool) {
	return d.GetOkExists(field)
}

func getAppConfigClient(meta any, guid string) (*appconfigurationv1.AppConfigurationV1, error) {
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

func formatValue(typ string, format any, value any) (any, error) {
	switch typ {
	case "BOOLEAN":
		convertedValue, ok := value.(bool)
		if !ok {
			return nil, flex.FmtErrorf("value not of type boolean")
		}
		return convertedValue, nil
	case "NUMERIC":
		convertedValue, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			return nil, flex.FmtErrorf("value not of type numeric: %s", err.Error())
		}
		return convertedValue, nil
	case "SECRETREF":
		stringValue := value.(string)
		config := map[string]any{}
		err := json.Unmarshal([]byte(stringValue), &config)
		if err != nil {
			return nil, flex.FmtErrorf("value not of type secret-reference: %s", err.Error())
		}
		return config, nil
	case "STRING":
		stringValue := value.(string)
		if formatString, ok := format.(string); ok {
			switch formatString {
			case "TEXT":
				return stringValue, nil
			case "JSON":
				config := map[string]any{}
				err := json.Unmarshal([]byte(stringValue), &config)
				if err != nil {
					return nil, flex.FmtErrorf("value not of type json: %s", err.Error())
				}
				return config, nil
			case "YAML":
				config := map[string]any{}
				err := yaml.Unmarshal([]byte(stringValue), &config)
				if err != nil {
					return nil, flex.FmtErrorf("value not of type yaml: %s", err.Error())
				}
				return config, nil
			}
		}
	}
	return nil, flex.FmtErrorf("invalid configuration of type and format")
}
