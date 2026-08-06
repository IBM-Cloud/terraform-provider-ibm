// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package appconfiguration

import (
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
)

func DataSourceIBMAppConfigWorkflowConfig() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIbmAppConfigWorkflowConfigRead,

		Schema: map[string]*schema.Schema{
			"guid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.",
			},
			"workflow_config_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique identifier of the workflow configuration.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the workflow configuration.",
			},
			"workflow_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique identifier for the workflow configuration.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the workflow is enabled.",
			},
			// Named "workflow_provider" instead of "provider" because "provider" is a
			// reserved meta-argument in Terraform's plugin SDK and cannot be used as a
			// schema field name. It maps to the "provider" field in the IBM App
			// Configuration API response.
			"workflow_provider": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Provider configuration for the workflow. Corresponds to the 'provider' field in the API.",
				Elem:        workflowProviderSchema(),
			},
			"scope": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Scope configuration for the workflow.",
				Elem:        workflowScopeSchema(),
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

func dataSourceIbmAppConfigWorkflowConfigRead(d *schema.ResourceData, meta interface{}) error {
	guid := d.Get("guid").(string)

	appconfigClient, err := getAppConfigClient(meta, guid)
	if err != nil {
		return flex.FmtErrorf("%s", err)
	}

	options := &appconfigurationv1.GetWorkflowConfigOptions{}
	options.SetWorkflowConfigID(d.Get("workflow_config_id").(string))

	result, response, err := appconfigClient.GetWorkflowConfig(options)
	if err != nil {
		return flex.FmtErrorf("[ERROR] GetWorkflowConfig failed %s\n%s", err, response)
	}

	d.SetId(fmt.Sprintf("%s/%s", guid, d.Get("workflow_config_id").(string)))

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

// flattenWorkflowProvider converts a *WorkflowProvider SDK struct into Terraform state.
func flattenWorkflowProvider(p *appconfigurationv1.WorkflowProvider) []map[string]interface{} {
	providerMap := map[string]interface{}{}
	if p.Type != nil {
		providerMap["type"] = p.Type
	}
	if p.Metadata != nil {
		if m, ok := p.Metadata.(*appconfigurationv1.WorkflowMetadata); ok {
			meta := map[string]interface{}{}
			if m.ApprovalExpiration != nil {
				meta["approval_expiration"] = flex.IntValue(m.ApprovalExpiration)
			}
			if m.ApprovalGroupName != nil {
				meta["approval_group_name"] = m.ApprovalGroupName
			}
			if m.SmSecretID != nil {
				meta["sm_secret_id"] = m.SmSecretID
			}
			if m.WorkflowURL != nil {
				meta["workflow_url"] = m.WorkflowURL
			}
			if m.CrnMask != nil {
				meta["crn_mask"] = m.CrnMask
			}
			if m.DefaultEmail != nil {
				meta["default_email"] = m.DefaultEmail
			}
			if m.SmSecretCrn != nil {
				meta["sm_secret_crn"] = m.SmSecretCrn
			}
			providerMap["metadata"] = []map[string]interface{}{meta}
		}
	}
	return []map[string]interface{}{providerMap}
}

// flattenWorkflowScope converts a *WorkflowScope SDK struct into Terraform state.
func flattenWorkflowScope(s *appconfigurationv1.WorkflowScope) []map[string]interface{} {
	scopeMap := map[string]interface{}{}
	if s.Collections != nil {
		scopeMap["collections"] = []map[string]interface{}{{"mode": s.Collections.Mode, "ids": s.Collections.Ids}}
	}
	if s.Segments != nil {
		scopeMap["segments"] = []map[string]interface{}{{"mode": s.Segments.Mode, "ids": s.Segments.Ids}}
	}
	if s.Environments != nil {
		envs := []map[string]interface{}{}
		for _, e := range s.Environments {
			envMap := map[string]interface{}{}
			if e.EnvironmentID != nil {
				envMap["environment_id"] = e.EnvironmentID
			}
			if e.Resources != nil {
				res := map[string]interface{}{}
				if e.Resources.Environment != nil && e.Resources.Environment.Enable != nil {
					res["environment"] = []map[string]interface{}{{"enable": e.Resources.Environment.Enable}}
				}
				if e.Resources.Features != nil {
					res["features"] = []map[string]interface{}{{"mode": e.Resources.Features.Mode, "ids": e.Resources.Features.Ids}}
				}
				if e.Resources.Properties != nil {
					res["properties"] = []map[string]interface{}{{"mode": e.Resources.Properties.Mode, "ids": e.Resources.Properties.Ids}}
				}
				envMap["resources"] = []map[string]interface{}{res}
			}
			envs = append(envs, envMap)
		}
		scopeMap["environments"] = envs
	}
	return []map[string]interface{}{scopeMap}
}
