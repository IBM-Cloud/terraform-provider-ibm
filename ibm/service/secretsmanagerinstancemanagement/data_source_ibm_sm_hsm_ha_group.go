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

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

func DataSourceIbmSmHsmHaGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmSmHsmHaGroupRead,

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
			"ha_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The identifier of the HA Group to operate on. Matches the `id` field in the `HSMHAGroup` response schema.",
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
	}
}

func dataSourceIbmSmHsmHaGroupRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_sm_hsm_ha_group", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getInstanceHsmDestinationHaGroupOptions := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationHaGroupOptions{}

	getInstanceHsmDestinationHaGroupOptions.SetInstanceID(d.Get("instance_id").(string))
	getInstanceHsmDestinationHaGroupOptions.SetDestinationID(d.Get("destination_id").(string))
	getInstanceHsmDestinationHaGroupOptions.SetHaGroupID(d.Get("ha_group_id").(string))

	hSMHAGroup, _, err := secretsManagerInstanceManagementClient.GetInstanceHsmDestinationHaGroupWithContext(context, getInstanceHsmDestinationHaGroupOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceHsmDestinationHaGroupWithContext failed: %s", err.Error()), "(Data) ibm_sm_hsm_ha_group", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", *getInstanceHsmDestinationHaGroupOptions.InstanceID, *getInstanceHsmDestinationHaGroupOptions.DestinationID, *getInstanceHsmDestinationHaGroupOptions.HaGroupID))

	if !core.IsNil(hSMHAGroup.Href) {
		if err = d.Set("href", hSMHAGroup.Href); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting href: %s", err), "(Data) ibm_sm_hsm_ha_group", "read", "set-href").GetDiag()
		}
	}

	if err = d.Set("name", hSMHAGroup.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting name: %s", err), "(Data) ibm_sm_hsm_ha_group", "read", "set-name").GetDiag()
	}

	if !core.IsNil(hSMHAGroup.Description) {
		if err = d.Set("description", hSMHAGroup.Description); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting description: %s", err), "(Data) ibm_sm_hsm_ha_group", "read", "set-description").GetDiag()
		}
	}

	if err = d.Set("provisioning_state", hSMHAGroup.ProvisioningState); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting provisioning_state: %s", err), "(Data) ibm_sm_hsm_ha_group", "read", "set-provisioning_state").GetDiag()
	}

	if err = d.Set("health_state", hSMHAGroup.HealthState); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting health_state: %s", err), "(Data) ibm_sm_hsm_ha_group", "read", "set-health_state").GetDiag()
	}

	if err = d.Set("health_as_of", flex.DateTimeToString(hSMHAGroup.HealthAsOf)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting health_as_of: %s", err), "(Data) ibm_sm_hsm_ha_group", "read", "set-health_as_of").GetDiag()
	}

	if err = d.Set("group_label", hSMHAGroup.GroupLabel); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting group_label: %s", err), "(Data) ibm_sm_hsm_ha_group", "read", "set-group_label").GetDiag()
	}

	if err = d.Set("group_serial", hSMHAGroup.GroupSerial); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting group_serial: %s", err), "(Data) ibm_sm_hsm_ha_group", "read", "set-group_serial").GetDiag()
	}

	memberSerials := []interface{}{}
	for _, memberSerialsItem := range hSMHAGroup.MemberSerials {
		memberSerials = append(memberSerials, memberSerialsItem)
	}
	if err = d.Set("member_serials", memberSerials); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting member_serials: %s", err), "(Data) ibm_sm_hsm_ha_group", "read", "set-member_serials").GetDiag()
	}

	if err = d.Set("created_at", flex.DateTimeToString(hSMHAGroup.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting created_at: %s", err), "(Data) ibm_sm_hsm_ha_group", "read", "set-created_at").GetDiag()
	}

	if !core.IsNil(hSMHAGroup.UpdatedAt) {
		if err = d.Set("updated_at", flex.DateTimeToString(hSMHAGroup.UpdatedAt)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting updated_at: %s", err), "(Data) ibm_sm_hsm_ha_group", "read", "set-updated_at").GetDiag()
		}
	}

	if err = d.Set("created_by", hSMHAGroup.CreatedBy); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting created_by: %s", err), "(Data) ibm_sm_hsm_ha_group", "read", "set-created_by").GetDiag()
	}

	if !core.IsNil(hSMHAGroup.UpdatedBy) {
		if err = d.Set("updated_by", hSMHAGroup.UpdatedBy); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting updated_by: %s", err), "(Data) ibm_sm_hsm_ha_group", "read", "set-updated_by").GetDiag()
		}
	}

	return nil
}
