// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.104.0-b4a47c49-20250418-184351
 */

package logs

import (
	"context"
	"fmt"
	"log"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
)

func DataSourceIbmLogsPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmLogsPolicyRead,

		Schema: map[string]*schema.Schema{
			"logs_policy_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of policy.",
			},
			"before": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Policy ID.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Policy name.",
						},
					},
				},
			},
			"company_id": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Company ID.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of policy.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of policy.",
			},
			"priority": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The data pipeline sources that match the policy rules will go through.",
			},
			"deleted": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Soft deletion flag.",
			},
			"enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enabled flag.",
			},
			"order": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Order of policy in relation to other policies.",
			},
			"application_rule": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Rule for matching with application.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_type_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Identifier of the rule.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Value of the rule. Multiple values can be provided as comma separated string of values.",
						},
					},
				},
			},
			"subsystem_rule": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Rule for matching with application.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_type_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Identifier of the rule.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Value of the rule. Multiple values can be provided as comma separated string of values.",
						},
					},
				},
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Created at date at utc+0.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Updated at date at utc+0.",
			},
			"archive_retention": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Archive retention definition.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "References archive retention definition.",
						},
					},
				},
			},
			"log_rules": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Log rules.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"severities": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Source severities to match with.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmLogsPolicyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsClient, err := meta.(conns.ClientSession).LogsV0()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_policy", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	region := getLogsInstanceRegion(logsClient, d)
	instanceId := d.Get("instance_id").(string)
	logsClient, err = getClientWithLogsInstanceEndpoint(logsClient, meta, instanceId, region, getLogsInstanceEndpointType(logsClient, d))
	if err != nil {
		return diag.FromErr(fmt.Errorf("Unable to get updated logs instance client"))
	}

	getPolicyOptions := &logsv0.GetPolicyOptions{}

	getPolicyOptions.SetID(core.UUIDPtr(strfmt.UUID(d.Get("logs_policy_id").(string))))

	policyIntf, _, err := logsClient.GetPolicyWithContext(context, getPolicyOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetPolicyWithContext failed: %s", err.Error()), "(Data) ibm_logs_policy", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	policy := policyIntf.(*logsv0.Policy)

	d.SetId(fmt.Sprintf("%s", *getPolicyOptions.ID))

	if !core.IsNil(policy.Before) {
		before := []map[string]interface{}{}
		beforeMap, err := DataSourceIbmLogsPolicyPolicyBeforeToMap(policy.Before)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_policy", "read", "before-to-map").GetDiag()
		}
		before = append(before, beforeMap)
		if err = d.Set("before", before); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting before: %s", err), "(Data) ibm_logs_policy", "read", "set-before").GetDiag()
		}
	}

	if err = d.Set("company_id", flex.IntValue(policy.CompanyID)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting company_id: %s", err), "(Data) ibm_logs_policy", "read", "set-company_id").GetDiag()
	}

	if err = d.Set("name", policy.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_logs_policy", "read", "set-name").GetDiag()
	}

	if err = d.Set("description", policy.Description); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting description: %s", err), "(Data) ibm_logs_policy", "read", "set-description").GetDiag()
	}

	if !core.IsNil(policy.Priority) {
		if err = d.Set("priority", policy.Priority); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting priority: %s", err), "(Data) ibm_logs_policy", "read", "set-priority").GetDiag()
		}
	}

	if !core.IsNil(policy.Deleted) {
		if err = d.Set("deleted", policy.Deleted); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting deleted: %s", err), "(Data) ibm_logs_policy", "read", "set-deleted").GetDiag()
		}
	}

	if !core.IsNil(policy.Enabled) {
		if err = d.Set("enabled", policy.Enabled); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting enabled: %s", err), "(Data) ibm_logs_policy", "read", "set-enabled").GetDiag()
		}
	}

	if err = d.Set("order", flex.IntValue(policy.Order)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting order: %s", err), "(Data) ibm_logs_policy", "read", "set-order").GetDiag()
	}

	if !core.IsNil(policy.ApplicationRule) {
		applicationRule := []map[string]interface{}{}
		applicationRuleMap, err := DataSourceIbmLogsPolicyQuotaV1RuleToMap(policy.ApplicationRule)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_policy", "read", "application_rule-to-map").GetDiag()
		}
		applicationRule = append(applicationRule, applicationRuleMap)
		if err = d.Set("application_rule", applicationRule); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting application_rule: %s", err), "(Data) ibm_logs_policy", "read", "set-application_rule").GetDiag()
		}
	}

	if !core.IsNil(policy.SubsystemRule) {
		subsystemRule := []map[string]interface{}{}
		subsystemRuleMap, err := DataSourceIbmLogsPolicyQuotaV1RuleToMap(policy.SubsystemRule)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_policy", "read", "subsystem_rule-to-map").GetDiag()
		}
		subsystemRule = append(subsystemRule, subsystemRuleMap)
		if err = d.Set("subsystem_rule", subsystemRule); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting subsystem_rule: %s", err), "(Data) ibm_logs_policy", "read", "set-subsystem_rule").GetDiag()
		}
	}

	if err = d.Set("created_at", policy.CreatedAt); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_logs_policy", "read", "set-created_at").GetDiag()
	}

	if err = d.Set("updated_at", policy.UpdatedAt); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting updated_at: %s", err), "(Data) ibm_logs_policy", "read", "set-updated_at").GetDiag()
	}

	if !core.IsNil(policy.ArchiveRetention) {
		archiveRetention := []map[string]interface{}{}
		archiveRetentionMap, err := DataSourceIbmLogsPolicyQuotaV1ArchiveRetentionToMap(policy.ArchiveRetention)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_policy", "read", "archive_retention-to-map").GetDiag()
		}
		archiveRetention = append(archiveRetention, archiveRetentionMap)
		if err = d.Set("archive_retention", archiveRetention); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting archive_retention: %s", err), "(Data) ibm_logs_policy", "read", "set-archive_retention").GetDiag()
		}
	}

	if !core.IsNil(policy.LogRules) {
		logRules := []map[string]interface{}{}
		logRulesMap, err := DataSourceIbmLogsPolicyQuotaV1LogRulesToMap(policy.LogRules)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_policy", "read", "log_rules-to-map").GetDiag()
		}
		logRules = append(logRules, logRulesMap)
		if err = d.Set("log_rules", logRules); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting log_rules: %s", err), "(Data) ibm_logs_policy", "read", "set-log_rules").GetDiag()
		}
	}

	return nil
}

func DataSourceIbmLogsPolicyPolicyBeforeToMap(model *logsv0.PolicyBefore) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID.String()
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	return modelMap, nil
}

func DataSourceIbmLogsPolicyQuotaV1RuleToMap(model *logsv0.QuotaV1Rule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["rule_type_id"] = *model.RuleTypeID
	modelMap["name"] = *model.Name
	return modelMap, nil
}

func DataSourceIbmLogsPolicyQuotaV1ArchiveRetentionToMap(model *logsv0.QuotaV1ArchiveRetention) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID.String()
	return modelMap, nil
}

func DataSourceIbmLogsPolicyQuotaV1LogRulesToMap(model *logsv0.QuotaV1LogRules) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Severities != nil {
		modelMap["severities"] = model.Severities
	}
	return modelMap, nil
}
