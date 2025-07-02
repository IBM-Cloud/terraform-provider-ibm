// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.96.1-5136e54a-20241108-203028
 */

package backuprecovery

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func DataSourceIbmBackupRecoveryManagerGetResources() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryManagerGetResourcesRead,

		Schema: map[string]*schema.Schema{
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the type of the resource.",
			},
			"external_targets": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the list of External Targets.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the id of the External target.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the name of the External target.",
						},
						"system_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the id of the System.",
						},
						"system_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the name of the System.",
						},
						"target_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the type of the External target.",
						},
					},
				},
			},
			"message_code_mappings": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the list of Message codes.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"message_code": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the message code.",
						},
						"message_guid": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the message GUID.",
						},
					},
				},
			},
			"policies": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the list of Protection Groups.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the id of the Protection Policy.",
						},
						"is_global_policy": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Specifies whether this is a Global Policy.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the name of the Protection Policy.",
						},
						"system_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the id of the System.",
						},
						"system_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the name of the System.",
						},
					},
				},
			},
			"protection_groups": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the list of Protection Groups.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the id of the Protection Group.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the name of the Protection Group.",
						},
						"system_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the id of the System.",
						},
						"system_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the name of the System.",
						},
					},
				},
			},
			"sources": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the list of Registered sources.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"environments": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies list of all environments discovered as part of this source.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the name of the registered source.",
						},
						"uuid": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the unique identifier of registered source.",
						},
					},
				},
			},
			"tenants": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the list of Tenants.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the id of the Tenant.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the name of the Tenant.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmBackupRecoveryManagerGetResourcesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryManagerV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_resources", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getResourcesOptions := &backuprecoveryv1.GetResourcesOptions{}

	getResourcesOptions.SetResourceType(d.Get("resource_type").(string))

	resources, _, err := backupRecoveryClient.GetResourcesWithContext(context, getResourcesOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetResourcesWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_manager_get_resources", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBackupRecoveryManagerGetResourcesID(d))

	if !core.IsNil(resources.ExternalTargets) {
		externalTargets := []map[string]interface{}{}
		for _, externalTargetsItem := range resources.ExternalTargets {
			externalTargetsItemMap, err := DataSourceIbmBackupRecoveryManagerGetResourcesExternalTargetToMap(&externalTargetsItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_resources", "read", "external_targets-to-map").GetDiag()
			}
			externalTargets = append(externalTargets, externalTargetsItemMap)
		}
		if err = d.Set("external_targets", externalTargets); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting external_targets: %s", err), "(Data) ibm_backup_recovery_manager_get_resources", "read", "set-external_targets").GetDiag()
		}
	}

	if !core.IsNil(resources.MessageCodeMappings) {
		messageCodeMappings := []map[string]interface{}{}
		for _, messageCodeMappingsItem := range resources.MessageCodeMappings {
			messageCodeMappingsItemMap, err := DataSourceIbmBackupRecoveryManagerGetResourcesMessageCodeMappingToMap(&messageCodeMappingsItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_resources", "read", "message_code_mappings-to-map").GetDiag()
			}
			messageCodeMappings = append(messageCodeMappings, messageCodeMappingsItemMap)
		}
		if err = d.Set("message_code_mappings", messageCodeMappings); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting message_code_mappings: %s", err), "(Data) ibm_backup_recovery_manager_get_resources", "read", "set-message_code_mappings").GetDiag()
		}
	}

	if !core.IsNil(resources.Policies) {
		policies := []map[string]interface{}{}
		for _, policiesItem := range resources.Policies {
			policiesItemMap, err := DataSourceIbmBackupRecoveryManagerGetResourcesPolicyToMap(&policiesItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_resources", "read", "policies-to-map").GetDiag()
			}
			policies = append(policies, policiesItemMap)
		}
		if err = d.Set("policies", policies); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting policies: %s", err), "(Data) ibm_backup_recovery_manager_get_resources", "read", "set-policies").GetDiag()
		}
	}

	if !core.IsNil(resources.ProtectionGroups) {
		protectionGroups := []map[string]interface{}{}
		for _, protectionGroupsItem := range resources.ProtectionGroups {
			protectionGroupsItemMap, err := DataSourceIbmBackupRecoveryManagerGetResourcesProtectionGroupToMap(&protectionGroupsItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_resources", "read", "protection_groups-to-map").GetDiag()
			}
			protectionGroups = append(protectionGroups, protectionGroupsItemMap)
		}
		if err = d.Set("protection_groups", protectionGroups); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting protection_groups: %s", err), "(Data) ibm_backup_recovery_manager_get_resources", "read", "set-protection_groups").GetDiag()
		}
	}

	if !core.IsNil(resources.Sources) {
		sources := []map[string]interface{}{}
		for _, sourcesItem := range resources.Sources {
			sourcesItemMap, err := DataSourceIbmBackupRecoveryManagerGetResourcesRegisteredSourceToMap(&sourcesItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_resources", "read", "sources-to-map").GetDiag()
			}
			sources = append(sources, sourcesItemMap)
		}
		if err = d.Set("sources", sources); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting sources: %s", err), "(Data) ibm_backup_recovery_manager_get_resources", "read", "set-sources").GetDiag()
		}
	}

	if !core.IsNil(resources.Tenants) {
		tenants := []map[string]interface{}{}
		for _, tenantsItem := range resources.Tenants {
			tenantsItemMap, err := DataSourceIbmBackupRecoveryManagerGetResourcesTenantToMap(&tenantsItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_resources", "read", "tenants-to-map").GetDiag()
			}
			tenants = append(tenants, tenantsItemMap)
		}
		if err = d.Set("tenants", tenants); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting tenants: %s", err), "(Data) ibm_backup_recovery_manager_get_resources", "read", "set-tenants").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmBackupRecoveryManagerGetResourcesID returns a reasonable ID for the list.
func dataSourceIbmBackupRecoveryManagerGetResourcesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBackupRecoveryManagerGetResourcesExternalTargetToMap(model *backuprecoveryv1.ExternalTarget) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.SystemID != nil {
		modelMap["system_id"] = *model.SystemID
	}
	if model.SystemName != nil {
		modelMap["system_name"] = *model.SystemName
	}
	if model.TargetType != nil {
		modelMap["target_type"] = *model.TargetType
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetResourcesMessageCodeMappingToMap(model *backuprecoveryv1.MessageCodeMapping) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MessageCode != nil {
		modelMap["message_code"] = *model.MessageCode
	}
	if model.MessageGuid != nil {
		modelMap["message_guid"] = *model.MessageGuid
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetResourcesPolicyToMap(model *backuprecoveryv1.Policy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.IsGlobalPolicy != nil {
		modelMap["is_global_policy"] = *model.IsGlobalPolicy
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.SystemID != nil {
		modelMap["system_id"] = *model.SystemID
	}
	if model.SystemName != nil {
		modelMap["system_name"] = *model.SystemName
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetResourcesProtectionGroupToMap(model *backuprecoveryv1.ProtectionGroup) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.SystemID != nil {
		modelMap["system_id"] = *model.SystemID
	}
	if model.SystemName != nil {
		modelMap["system_name"] = *model.SystemName
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetResourcesRegisteredSourceToMap(model *backuprecoveryv1.RegisteredSource) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Environments != nil {
		modelMap["environments"] = model.Environments
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.UUID != nil {
		modelMap["uuid"] = *model.UUID
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetResourcesTenantToMap(model *backuprecoveryv1.Tenant) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	return modelMap, nil
}
