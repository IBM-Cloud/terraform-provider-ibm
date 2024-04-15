// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/observability-c/dragonlog-logs-go-sdk/logsv0"
)

func DataSourceIbmLogsPolicies() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmLogsPoliciesRead,

		Schema: map[string]*schema.Schema{
			"enabled_only": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "optionally filter only enabled policies.",
			},
			"source_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "logs",
				Description: "Source type to filter policies by.",
			},
			"policies": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "company policies.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "policy id.",
						},
						"company_id": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "company id.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "name of policy.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "description of policy.",
						},
						"priority": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "the data pipeline sources that match the policy rules will go through.",
						},
						"deleted": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "soft deletion flag.",
						},
						"enabled": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "enabled flag.",
						},
						"order": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "order of policy in relation to other policies.",
						},
						"application_rule": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "rule for matching with application.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rule_type_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "identifier of the rule.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "value of the rule.",
									},
								},
							},
						},
						"subsystem_rule": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "rule for matching with application.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rule_type_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "identifier of the rule.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "value of the rule.",
									},
								},
							},
						},
						"created_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "created at timestamp.",
						},
						"updated_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "updated at timestamp.",
						},
						"archive_retention": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "archive retention definition.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "references archive retention definition.",
									},
								},
							},
						},
						"log_rules": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "log rules.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"severities": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "source severities to match with.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmLogsPoliciesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsClient, err := meta.(conns.ClientSession).LogsV0()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_logs_policies", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	region := getLogsInstanceRegion(logsClient, d)
	instanceId := d.Get("instance_id").(string)
	logsClient = getClientWithLogsInstanceEndpoint(logsClient, instanceId, region, getLogsInstanceEndpointType(logsClient, d))

	getCompanyPoliciesOptions := &logsv0.GetCompanyPoliciesOptions{}

	if _, ok := d.GetOk("enabled_only"); ok {
		getCompanyPoliciesOptions.SetEnabledOnly(d.Get("enabled_only").(bool))
	}
	if _, ok := d.GetOk("source_type"); ok {
		getCompanyPoliciesOptions.SetSourceType(d.Get("source_type").(string))
	}

	policyCollection, _, err := logsClient.GetCompanyPoliciesWithContext(context, getCompanyPoliciesOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetCompanyPoliciesWithContext failed: %s", err.Error()), "(Data) ibm_logs_policies", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmLogsPoliciesID(d))

	policies := []map[string]interface{}{}
	if policyCollection.Policies != nil {
		for _, modelItem := range policyCollection.Policies {
			modelMap, err := DataSourceIbmLogsPoliciesPolicyToMap(modelItem)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_logs_policies", "read")
				return tfErr.GetDiag()
			}
			policies = append(policies, modelMap)
		}
	}
	if err = d.Set("policies", policies); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting policies: %s", err), "(Data) ibm_logs_policies", "read")
		return tfErr.GetDiag()
	}

	return nil
}

// dataSourceIbmLogsPoliciesID returns a reasonable ID for the list.
func dataSourceIbmLogsPoliciesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmLogsPoliciesPolicyToMap(model logsv0.PolicyIntf) (map[string]interface{}, error) {
	if _, ok := model.(*logsv0.PolicyQuotaV1PolicySourceTypeRulesLogRules); ok {
		return DataSourceIbmLogsPoliciesPolicyQuotaV1PolicySourceTypeRulesLogRulesToMap(model.(*logsv0.PolicyQuotaV1PolicySourceTypeRulesLogRules))
	} else if _, ok := model.(*logsv0.Policy); ok {
		modelMap := make(map[string]interface{})
		model := model.(*logsv0.Policy)
		if model.ID != nil {
			modelMap["id"] = *model.ID
		}
		if model.CompanyID != nil {
			modelMap["company_id"] = flex.IntValue(model.CompanyID)
		}
		if model.Name != nil {
			modelMap["name"] = *model.Name
		}
		if model.Description != nil {
			modelMap["description"] = *model.Description
		}
		if model.Priority != nil {
			modelMap["priority"] = *model.Priority
		}
		if model.Deleted != nil {
			modelMap["deleted"] = *model.Deleted
		}
		if model.Enabled != nil {
			modelMap["enabled"] = *model.Enabled
		}
		if model.Order != nil {
			modelMap["order"] = flex.IntValue(model.Order)
		}
		if model.ApplicationRule != nil {
			applicationRuleMap, err := DataSourceIbmLogsPoliciesQuotaV1RuleToMap(model.ApplicationRule)
			if err != nil {
				return modelMap, err
			}
			modelMap["application_rule"] = []map[string]interface{}{applicationRuleMap}
		}
		if model.SubsystemRule != nil {
			subsystemRuleMap, err := DataSourceIbmLogsPoliciesQuotaV1RuleToMap(model.SubsystemRule)
			if err != nil {
				return modelMap, err
			}
			modelMap["subsystem_rule"] = []map[string]interface{}{subsystemRuleMap}
		}
		if model.CreatedAt != nil {
			modelMap["created_at"] = *model.CreatedAt
		}
		if model.UpdatedAt != nil {
			modelMap["updated_at"] = *model.UpdatedAt
		}
		if model.ArchiveRetention != nil {
			archiveRetentionMap, err := DataSourceIbmLogsPoliciesQuotaV1ArchiveRetentionToMap(model.ArchiveRetention)
			if err != nil {
				return modelMap, err
			}
			modelMap["archive_retention"] = []map[string]interface{}{archiveRetentionMap}
		}
		if model.LogRules != nil {
			logRulesMap, err := DataSourceIbmLogsPoliciesQuotaV1LogRulesToMap(model.LogRules)
			if err != nil {
				return modelMap, err
			}
			modelMap["log_rules"] = []map[string]interface{}{logRulesMap}
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized logsv0.PolicyIntf subtype encountered")
	}
}

func DataSourceIbmLogsPoliciesQuotaV1RuleToMap(model *logsv0.QuotaV1Rule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["rule_type_id"] = *model.RuleTypeID
	modelMap["name"] = *model.Name
	return modelMap, nil
}

func DataSourceIbmLogsPoliciesQuotaV1ArchiveRetentionToMap(model *logsv0.QuotaV1ArchiveRetention) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	return modelMap, nil
}

func DataSourceIbmLogsPoliciesQuotaV1LogRulesToMap(model *logsv0.QuotaV1LogRules) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Severities != nil {
		modelMap["severities"] = model.Severities
	}
	return modelMap, nil
}

func DataSourceIbmLogsPoliciesPolicyQuotaV1PolicySourceTypeRulesLogRulesToMap(model *logsv0.PolicyQuotaV1PolicySourceTypeRulesLogRules) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.CompanyID != nil {
		modelMap["company_id"] = flex.IntValue(model.CompanyID)
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.Priority != nil {
		modelMap["priority"] = *model.Priority
	}
	if model.Deleted != nil {
		modelMap["deleted"] = *model.Deleted
	}
	if model.Enabled != nil {
		modelMap["enabled"] = *model.Enabled
	}
	if model.Order != nil {
		modelMap["order"] = flex.IntValue(model.Order)
	}
	if model.ApplicationRule != nil {
		applicationRuleMap, err := DataSourceIbmLogsPoliciesQuotaV1RuleToMap(model.ApplicationRule)
		if err != nil {
			return modelMap, err
		}
		modelMap["application_rule"] = []map[string]interface{}{applicationRuleMap}
	}
	if model.SubsystemRule != nil {
		subsystemRuleMap, err := DataSourceIbmLogsPoliciesQuotaV1RuleToMap(model.SubsystemRule)
		if err != nil {
			return modelMap, err
		}
		modelMap["subsystem_rule"] = []map[string]interface{}{subsystemRuleMap}
	}
	if model.CreatedAt != nil {
		modelMap["created_at"] = *model.CreatedAt
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = *model.UpdatedAt
	}
	if model.ArchiveRetention != nil {
		archiveRetentionMap, err := DataSourceIbmLogsPoliciesQuotaV1ArchiveRetentionToMap(model.ArchiveRetention)
		if err != nil {
			return modelMap, err
		}
		modelMap["archive_retention"] = []map[string]interface{}{archiveRetentionMap}
	}
	if model.LogRules != nil {
		logRulesMap, err := DataSourceIbmLogsPoliciesQuotaV1LogRulesToMap(model.LogRules)
		if err != nil {
			return modelMap, err
		}
		modelMap["log_rules"] = []map[string]interface{}{logRulesMap}
	}
	return modelMap, nil
}
