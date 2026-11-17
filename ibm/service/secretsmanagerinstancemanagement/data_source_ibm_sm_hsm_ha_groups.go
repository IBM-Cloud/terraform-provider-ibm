// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.117.3-3e91ab11-20261028-182643
 */

package secretsmanagerinstancemanagement

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

func DataSourceIbmSmHsmHaGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmSmHsmHaGroupsRead,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Secrets Manager instance ID.",
			},
			"destination_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Destination ID.",
			},
			"ha_groups": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "HSM HA Groups for this destination. At most 20 HSM HA Groups may exist per destination.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "HA Group ID. Immutable. System-managed.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL of the HA Group resource.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Human-readable display name.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Optional description. Omitted if not set.",
						},
						"provisioning_state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Lifecycle state of the HA Group:- `completed`: HA Group has been created and its config has been staged in the HSM config file. Initial and only state.**Note:** New values may be added to this enumeration in future API versions. Clients MUST handle unknown values gracefully and MUST NOT fail on unrecognised states.",
						},
						"health_state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "AS OF health state. Written exclusively by POST …/health and POST …/apply:- `unknown`: Initial state on creation. No health check has been performed yet.- `evaluating`: A health check is in progress for this resource.- `in_progress`: An apply operation is running across the destination.- `healthy`: Configured AND all member HSM nodes are healthy.- `degraded`: Configured AND any one member is unhealthy (but not all).- `unhealthy`: Configured AND all members are unhealthy or unknown.- `not_configured`: Prerequisites not met (destination needs ≥1 path with active non-expired Server CA AND ≥1 active non-expired client certificate).",
						},
						"health_as_of": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health.",
						},
						"group_label": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "HA Group virtual token label. Immutable after creation.",
						},
						"group_serial": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "HA Group serial number. Immutable after creation.",
						},
						"member_serials": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Ordered list of HSM node member serial numbers for this HA Group. Max 100 unique members.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"created_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Computed field.",
						},
						"updated_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Computed field. Omitted until first update.",
						},
						"created_by": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Computed field.",
						},
						"updated_by": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Computed field. Omitted until first update.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmSmHsmHaGroupsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_sm_hsm_ha_groups", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listInstanceHsmDestinationHaGroupsOptions := &secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationHaGroupsOptions{}

	listInstanceHsmDestinationHaGroupsOptions.SetInstanceID(d.Get("instance_id").(string))
	listInstanceHsmDestinationHaGroupsOptions.SetDestinationID(d.Get("destination_id").(string))

	hSMHAGroupCollection, _, err := secretsManagerInstanceManagementClient.ListInstanceHsmDestinationHaGroupsWithContext(context, listInstanceHsmDestinationHaGroupsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListInstanceHsmDestinationHaGroupsWithContext failed: %s", err.Error()), "(Data) ibm_sm_hsm_ha_groups", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmSmHsmHaGroupsID(d))

	haGroups := []map[string]interface{}{}
	for _, haGroupsItem := range hSMHAGroupCollection.HaGroups {
		haGroupsItemMap, err := DataSourceIbmSmHsmHaGroupsHSMHAGroupToMap(&haGroupsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_sm_hsm_ha_groups", "read", "ha_groups-to-map").GetDiag()
		}
		haGroups = append(haGroups, haGroupsItemMap)
	}
	if err = d.Set("ha_groups", haGroups); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting ha_groups: %s", err), "(Data) ibm_sm_hsm_ha_groups", "read", "set-ha_groups").GetDiag()
	}

	return nil
}

// dataSourceIbmSmHsmHaGroupsID returns a reasonable ID for the list.
func dataSourceIbmSmHsmHaGroupsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmSmHsmHaGroupsHSMHAGroupToMap(model *secretsmanagerinstancemanagementv2.HSMHAGroup) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	modelMap["name"] = *model.Name
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	modelMap["provisioning_state"] = *model.ProvisioningState
	modelMap["health_state"] = *model.HealthState
	modelMap["health_as_of"] = model.HealthAsOf.String()
	modelMap["group_label"] = *model.GroupLabel
	modelMap["group_serial"] = *model.GroupSerial
	modelMap["member_serials"] = model.MemberSerials
	modelMap["created_at"] = model.CreatedAt.String()
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	modelMap["created_by"] = *model.CreatedBy
	if model.UpdatedBy != nil {
		modelMap["updated_by"] = *model.UpdatedBy
	}
	return modelMap, nil
}
