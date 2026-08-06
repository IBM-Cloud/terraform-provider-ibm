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
		convertedValue, err := strconv.ParseBool(value.(string))
		if err != nil {
			return nil, flex.FmtErrorf("value not of type boolean: %s", err.Error())
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

// workflowScopeModeElem returns the schema.Resource for a WorkflowScopeMode (mode + ids).
func workflowScopeModeElem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mode": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Scope mode - only ALL is supported.",
			},
			"ids": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "List of IDs - always empty when mode is ALL.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

// workflowEnvironmentElem returns the schema.Resource for a WorkflowEnvironment entry.
func workflowEnvironmentElem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"environment_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Environment identifier.",
			},
			"resources": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Resources configuration for the environment.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"environment": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Environment-level workflow enable flag.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable": {
										Type:        schema.TypeBool,
										Optional:    true,
										Computed:    true,
										Description: "Whether environment-level workflow is enabled.",
									},
								},
							},
						},
						"features": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Scope mode for features.",
							Elem:        workflowScopeModeElem(),
						},
						"properties": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Scope mode for properties.",
							Elem:        workflowScopeModeElem(),
						},
					},
				},
			},
		},
	}
}

// workflowProviderSchema returns the Computed-only schema.Resource for a WorkflowProvider (used by data sources).
func workflowProviderSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Provider type (SERVICENOW_EXTERNAL or SERVICENOW_IBM).",
			},
			"metadata": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Metadata for the ServiceNow provider.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"approval_expiration": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Approval expiration time in days.",
						},
						"approval_group_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the approval group in ServiceNow.",
						},
						"sm_secret_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Secret Manager secret ID for ServiceNow credentials.",
						},
						"workflow_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ServiceNow instance URL.",
						},
						"crn_mask": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CRN mask for the service.",
						},
						"default_email": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The email id to which the change request will be assigned.",
						},
						"sm_secret_crn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Secret Manager secret CRN for ServiceNow credentials.",
						},
					},
				},
			},
		},
	}
}

// workflowScopeSchema returns the Computed-only schema.Resource for a WorkflowScope (used by data sources).
func workflowScopeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"collections": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Scope mode for collections.",
				Elem:        workflowScopeModeElem(),
			},
			"segments": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Scope mode for segments.",
				Elem:        workflowScopeModeElem(),
			},
			"environments": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of environments where the workflow applies.",
				Elem:        workflowEnvironmentElem(),
			},
		},
	}
}

// rolloutConfigurationSchema returns the shared schema element for rollout_configuration blocks.
func rolloutConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"duration_preset": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A predefined duration preset that sets the overall pace of the rollout. Use CUSTOM to define phases manually.",
			},
			"start_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The UTC date and time at which the rollout should start, in ISO 8601 format.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The current execution status of a rollout (QUEUED, RUNNING, STOPPED).",
			},
			"phases": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The ordered list of rollout phases.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"percentage": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The rollout percentage target for this phase.",
						},
						"duration": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The length of time to wait before advancing to the next phase.",
						},
						"duration_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unit of time for the duration field (minutes, hours, days).",
						},
					},
				},
			},
		},
	}
}
