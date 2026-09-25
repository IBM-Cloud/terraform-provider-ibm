// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package appconfiguration

import (
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMAppConfigWorkflowConfig() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIbmAppConfigWorkflowConfigCreate,
		Read:     resourceIbmAppConfigWorkflowConfigRead,
		Update:   resourceIbmAppConfigWorkflowConfigUpdate,
		Delete:   resourceIbmAppConfigWorkflowConfigDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"guid": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.",
			},
			"workflow_config_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique identifier of the workflow configuration (derived from href after create).",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the workflow configuration.",
			},
			"workflow_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique identifier for the workflow configuration.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Whether the workflow is enabled.",
			},
			// Named "workflow_provider" instead of "provider" because "provider" is a
			// reserved meta-argument in Terraform's plugin SDK and cannot be used as a
			// schema field name. It maps to the "provider" field in the IBM App
			// Configuration API request/response.
			"workflow_provider": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "Provider configuration for the workflow. Corresponds to the 'provider' field in the API.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Provider type (SERVICENOW_EXTERNAL or SERVICENOW_IBM).",
						},
						"metadata": {
							Type:        schema.TypeList,
							Required:    true,
							MaxItems:    1,
							Description: "Metadata for the ServiceNow provider.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"approval_expiration": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Approval expiration time in days.",
									},
									"approval_group_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Name of the approval group in ServiceNow.",
									},
									"sm_secret_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Secret Manager secret ID for ServiceNow credentials.",
									},
									"workflow_url": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "ServiceNow instance URL.",
									},
									"crn_mask": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "CRN mask for the service.",
									},
									"default_email": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The email id to which the change request will be assigned.",
									},
									"sm_secret_crn": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Secret Manager secret CRN for ServiceNow credentials.",
									},
								},
							},
						},
					},
				},
			},
			"scope": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "Scope configuration for the workflow. At least one of collections, segments, or environments must be specified.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"collections": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Scope mode for collections.",
							Elem:        workflowScopeModeElem(),
						},
						"segments": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Scope mode for segments.",
							Elem:        workflowScopeModeElem(),
						},
						"environments": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of environments where the workflow applies.",
							Elem:        workflowEnvironmentElem(),
						},
					},
				},
			},
			"created_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when the workflow was created.",
			},
			"updated_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when the workflow was last updated.",
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL to access this workflow configuration.",
			},
		},
	}
}

func resourceIbmAppConfigWorkflowConfigCreate(d *schema.ResourceData, meta interface{}) error {
	guid := d.Get("guid").(string)
	appconfigClient, err := getAppConfigClient(meta, guid)
	if err != nil {
		return flex.FmtErrorf("%s", err)
	}

	provider, err := resourceIbmAppConfigMapToWorkflowProvider(d.Get("workflow_provider").([]interface{}))
	if err != nil {
		return err
	}
	scope, err := resourceIbmAppConfigMapToWorkflowScope(d.Get("scope").([]interface{}))
	if err != nil {
		return err
	}

	options := &appconfigurationv1.CreateWorkflowConfigsOptions{}
	options.SetName(d.Get("name").(string))
	options.SetWorkflowID(d.Get("workflow_id").(string))
	options.SetEnabled(d.Get("enabled").(bool))
	options.SetProvider(provider)
	options.SetScope(scope)

	result, response, err := appconfigClient.CreateWorkflowConfigs(options)
	if err != nil {
		return flex.FmtErrorf("[ERROR] CreateWorkflowConfigs failed %s\n%s", err, response)
	}

	// Extract workflow_config_id from the href returned by the API.
	workflowConfigID, err := extractWorkflowConfigIDFromHref(result.Href)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s", guid, workflowConfigID))
	d.Set("workflow_config_id", workflowConfigID)

	return resourceIbmAppConfigWorkflowConfigRead(d, meta)
}

func resourceIbmAppConfigWorkflowConfigRead(d *schema.ResourceData, meta interface{}) error {
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return nil
	}
	if len(parts) < 2 {
		return flex.FmtErrorf("[ERROR] Invalid resource ID, expected guid/workflow_config_id")
	}
	guid, workflowConfigID := parts[0], parts[1]

	appconfigClient, err := getAppConfigClient(meta, guid)
	if err != nil {
		return flex.FmtErrorf("%s", err)
	}

	options := &appconfigurationv1.GetWorkflowConfigOptions{}
	options.SetWorkflowConfigID(workflowConfigID)

	result, response, err := appconfigClient.GetWorkflowConfig(options)
	if err != nil {
		return flex.FmtErrorf("[ERROR] GetWorkflowConfig failed %s\n%s", err, response)
	}

	if err = d.Set("workflow_config_id", workflowConfigID); err != nil {
		return flex.FmtErrorf("[ERROR] Error setting workflow_config_id: %s", err)
	}
	if result.Name != nil {
		if err = d.Set("name", result.Name); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting name: %s", err)
		}
	}
	if result.WorkflowID != nil {
		if err = d.Set("workflow_id", result.WorkflowID); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting workflow_id: %s", err)
		}
	}
	if result.Enabled != nil {
		if err = d.Set("enabled", result.Enabled); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting enabled: %s", err)
		}
	}
	if result.Provider != nil {
		if err = d.Set("workflow_provider", flattenWorkflowProvider(result.Provider)); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting workflow_provider: %s", err)
		}
	}
	if result.Scope != nil {
		if err = d.Set("scope", flattenWorkflowScope(result.Scope)); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting scope: %s", err)
		}
	}
	if result.CreatedTime != nil {
		if err = d.Set("created_time", result.CreatedTime.String()); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting created_time: %s", err)
		}
	}
	if result.UpdatedTime != nil {
		if err = d.Set("updated_time", result.UpdatedTime.String()); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting updated_time: %s", err)
		}
	}
	if result.Href != nil {
		if err = d.Set("href", result.Href); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting href: %s", err)
		}
	}

	return nil
}

func resourceIbmAppConfigWorkflowConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return nil
	}
	if len(parts) < 2 {
		return flex.FmtErrorf("[ERROR] Invalid resource ID, expected guid/workflow_config_id")
	}
	guid, workflowConfigID := parts[0], parts[1]

	appconfigClient, err := getAppConfigClient(meta, guid)
	if err != nil {
		return flex.FmtErrorf("%s", err)
	}

	if !d.HasChanges("name", "workflow_id", "enabled", "workflow_provider", "scope") {
		return nil
	}

	provider, err := resourceIbmAppConfigMapToWorkflowProvider(d.Get("workflow_provider").([]interface{}))
	if err != nil {
		return err
	}
	scope, err := resourceIbmAppConfigMapToWorkflowScope(d.Get("scope").([]interface{}))
	if err != nil {
		return err
	}

	options := &appconfigurationv1.UpdateWorkflowConfigsOptions{}
	options.SetWorkflowConfigID(workflowConfigID)
	options.SetName(d.Get("name").(string))
	options.SetWorkflowID(d.Get("workflow_id").(string))
	options.SetEnabled(d.Get("enabled").(bool))
	options.SetProvider(provider)
	options.SetScope(scope)

	_, response, err := appconfigClient.UpdateWorkflowConfigs(options)
	if err != nil {
		return flex.FmtErrorf("[ERROR] UpdateWorkflowConfigs failed %s\n%s", err, response)
	}

	return resourceIbmAppConfigWorkflowConfigRead(d, meta)
}

func resourceIbmAppConfigWorkflowConfigDelete(d *schema.ResourceData, meta interface{}) error {
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return nil
	}
	if len(parts) < 2 {
		return flex.FmtErrorf("[ERROR] Invalid resource ID, expected guid/workflow_config_id")
	}
	guid, workflowConfigID := parts[0], parts[1]

	appconfigClient, err := getAppConfigClient(meta, guid)
	if err != nil {
		return flex.FmtErrorf("%s", err)
	}

	options := &appconfigurationv1.DeleteWorkflowConfigsOptions{}
	options.SetWorkflowConfigID(workflowConfigID)

	response, err := appconfigClient.DeleteWorkflowConfigs(options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return flex.FmtErrorf("[ERROR] DeleteWorkflowConfigs failed %s\n%s", err, response)
	}

	d.SetId("")
	return nil
}

// extractWorkflowConfigIDFromHref parses the workflow_config_id from the href URL.
// e.g. ".../workflow/configs/{workflow_config_id}"
func extractWorkflowConfigIDFromHref(href *string) (string, error) {
	if href == nil {
		return "", flex.FmtErrorf("[ERROR] href is nil, cannot determine workflow_config_id")
	}
	// The path is .../workflow/configs/<id>, last segment is the id.
	path := *href
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			return path[i+1:], nil
		}
	}
	return path, nil
}

// resourceIbmAppConfigMapToWorkflowProvider maps Terraform input to *WorkflowProvider.
// The metadata struct is discriminated by the provider type:
//   - SERVICENOW_EXTERNAL → WorkflowMetadataExternalServiceNowMetadata
//     (fields: approval_expiration, approval_group_name, sm_secret_id, workflow_url)
//   - SERVICENOW_IBM → WorkflowMetadataIBMServiceNowMetadata
//     (fields: approval_expiration, crn_mask, sm_secret_crn, default_email)
func resourceIbmAppConfigMapToWorkflowProvider(providerList []interface{}) (*appconfigurationv1.WorkflowProvider, error) {
	if len(providerList) == 0 {
		return nil, flex.FmtErrorf("[ERROR] provider block is required")
	}
	pm := providerList[0].(map[string]interface{})
	provider := &appconfigurationv1.WorkflowProvider{}

	providerType := ""
	if v, ok := pm["type"].(string); ok && v != "" {
		providerType = v
		provider.Type = core.StringPtr(v)
	}

	if metaList, ok := pm["metadata"].([]interface{}); ok && len(metaList) > 0 {
		mm := metaList[0].(map[string]interface{})

		switch providerType {
		case "SERVICENOW_IBM":
			meta := &appconfigurationv1.WorkflowMetadataIBMServiceNowMetadata{}
			if v, ok := mm["approval_expiration"].(int); ok && v != 0 {
				meta.ApprovalExpiration = core.Int64Ptr(int64(v))
			}
			if v, ok := mm["crn_mask"].(string); ok && v != "" {
				meta.CrnMask = core.StringPtr(v)
			}
			if v, ok := mm["sm_secret_crn"].(string); ok && v != "" {
				meta.SmSecretCrn = core.StringPtr(v)
			}
			if v, ok := mm["default_email"].(string); ok && v != "" {
				meta.DefaultEmail = core.StringPtr(v)
			}
			provider.Metadata = meta
		default: // SERVICENOW_EXTERNAL
			meta := &appconfigurationv1.WorkflowMetadataExternalServiceNowMetadata{}
			if v, ok := mm["approval_expiration"].(int); ok && v != 0 {
				meta.ApprovalExpiration = core.Int64Ptr(int64(v))
			}
			if v, ok := mm["approval_group_name"].(string); ok && v != "" {
				meta.ApprovalGroupName = core.StringPtr(v)
			}
			if v, ok := mm["sm_secret_id"].(string); ok && v != "" {
				meta.SmSecretID = core.StringPtr(v)
			}
			if v, ok := mm["workflow_url"].(string); ok && v != "" {
				meta.WorkflowURL = core.StringPtr(v)
			}
			provider.Metadata = meta
		}
	}
	return provider, nil
}

// resourceIbmAppConfigMapToWorkflowScope maps Terraform input to *WorkflowScope.
func resourceIbmAppConfigMapToWorkflowScope(scopeList []interface{}) (*appconfigurationv1.WorkflowScope, error) {
	if len(scopeList) == 0 {
		return nil, flex.FmtErrorf("[ERROR] scope block is required")
	}
	sm := scopeList[0].(map[string]interface{})
	scope := &appconfigurationv1.WorkflowScope{}

	if collList, ok := sm["collections"].([]interface{}); ok && len(collList) > 0 {
		scope.Collections = mapToWorkflowScopeMode(collList[0].(map[string]interface{}))
	}
	if segList, ok := sm["segments"].([]interface{}); ok && len(segList) > 0 {
		scope.Segments = mapToWorkflowScopeMode(segList[0].(map[string]interface{}))
	}
	if envList, ok := sm["environments"].([]interface{}); ok && len(envList) > 0 {
		for _, e := range envList {
			em := e.(map[string]interface{})
			env := appconfigurationv1.WorkflowEnvironment{}
			if v, ok := em["environment_id"].(string); ok && v != "" {
				env.EnvironmentID = core.StringPtr(v)
			}
			if resList, ok := em["resources"].([]interface{}); ok && len(resList) > 0 {
				rm := resList[0].(map[string]interface{})
				res := &appconfigurationv1.WorkflowEnvironmentResources{}
				if envResList, ok := rm["environment"].([]interface{}); ok && len(envResList) > 0 {
					erm := envResList[0].(map[string]interface{})
					envRes := &appconfigurationv1.WorkflowEnvironmentResourcesEnvironment{}
					if v, ok := erm["enable"].(bool); ok {
						envRes.Enable = core.BoolPtr(v)
					}
					res.Environment = envRes
				}
				if featList, ok := rm["features"].([]interface{}); ok && len(featList) > 0 {
					res.Features = mapToWorkflowScopeMode(featList[0].(map[string]interface{}))
				}
				if propList, ok := rm["properties"].([]interface{}); ok && len(propList) > 0 {
					res.Properties = mapToWorkflowScopeMode(propList[0].(map[string]interface{}))
				}
				env.Resources = res
			}
			scope.Environments = append(scope.Environments, env)
		}
	}
	return scope, nil
}

func mapToWorkflowScopeMode(m map[string]interface{}) *appconfigurationv1.WorkflowScopeMode {
	sm := &appconfigurationv1.WorkflowScopeMode{}
	if v, ok := m["mode"].(string); ok && v != "" {
		sm.Mode = core.StringPtr(v)
	}
	ids := []string{}
	if idList, ok := m["ids"].([]interface{}); ok {
		for _, id := range idList {
			if s, ok := id.(string); ok {
				ids = append(ids, s)
			}
		}
	}
	sm.Ids = ids
	return sm
}
