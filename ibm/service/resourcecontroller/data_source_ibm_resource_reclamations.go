// Copyright IBM Corp. 2024
// Licensed under the Mozilla Public License v2.0

package resourcecontroller

import (
	"context"
	"fmt"
	"log"
	"time"

	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
)

// DataSourceIBMResourceReclamations returns the schema for the ibm_resource_reclamations data source.
//
// Example Terraform usage:
//
//	data "ibm_resource_reclamations" "all" {}
//	output "reclamation_ids" {
//	  value = [for r in data.ibm_resource_reclamations.all.reclamations : r.id]
//	}
func DataSourceIBMResourceReclamations() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMResourceReclamationsRead,
		Description: "Retrieve the list of all resource reclamations in the account.",
		Schema: map[string]*schema.Schema{
			"reclamations": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of resource reclamations.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID associated with the reclamation.",
						},
						"entity_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The entity ID for this reclamation.",
						},
						"entity_type_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The entity type ID for this reclamation.",
						},
						"entity_crn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The full CRN associated with this reclamation.",
						},
						"resource_instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource instance ID associated with the reclamation.",
						},
						"resource_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource group ID.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account ID.",
						},
						"policy_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The policy ID for the reclamation.",
						},
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The state of this reclamation.",
						},
						"target_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "When the reclamation retention period ends (RFC3339).",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date/time when created (RFC3339).",
						},
						"created_by": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subject who created this reclamation.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date/time when last updated (RFC3339).",
						},
						"updated_by": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subject who updated this reclamation.",
						},
						"custom_properties": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Custom properties set on the reclamation.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
		},
	}
}

// dataSourceIBMResourceReclamationsRead fetches all reclamations using Resource Controller V2
// and sets them into the Terraform state.
// Errors are always reported with flex.DiscriminatedTerraformErrorf and flex.TerraformErrorf to aid in debugging.
func dataSourceIBMResourceReclamationsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(
			err,
			err.Error(),
			"(Data) ibm_resource_reclamations",
			"read",
			"initialize-client",
		)
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	opts := &rc.ListReclamationsOptions{}
	reclamationsList, _, err := rsConClient.ListReclamationsWithContext(context, opts)
	if err != nil {
		tfErr := flex.TerraformErrorf(
			err,
			fmt.Sprintf("ListReclamations failed: %s", err.Error()),
			"(Data) ibm_resource_reclamations",
			"read",
		)
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	var flattened []map[string]interface{}
	for _, rec := range reclamationsList.Resources {
		flattened = append(flattened, flattenReclamationForDS(&rec))
	}

	d.SetId(fmt.Sprintf("reclamations-%d", time.Now().Unix()))
	if err := d.Set("reclamations", flattened); err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(
			err,
			fmt.Sprintf("Error setting reclamations: %s", err),
			"(Data) ibm_resource_reclamations",
			"read",
			"set-reclamations",
		)
		return tfErr.GetDiag()
	}

	return nil
}

// flattenReclamationForDS transforms rc.Reclamation into a flat map compatible with Terraform schema.
func flattenReclamationForDS(rec *rc.Reclamation) map[string]interface{} {
	m := make(map[string]interface{})
	if rec.ID != nil {
		m["id"] = *rec.ID
	}
	if rec.EntityID != nil {
		m["entity_id"] = *rec.EntityID
	}
	if rec.EntityTypeID != nil {
		m["entity_type_id"] = *rec.EntityTypeID
	}
	if rec.EntityCRN != nil {
		m["entity_crn"] = *rec.EntityCRN
	}
	if rec.ResourceInstanceID != nil {
		m["resource_instance_id"] = *rec.ResourceInstanceID
	}
	if rec.ResourceGroupID != nil {
		m["resource_group_id"] = *rec.ResourceGroupID
	}
	if rec.AccountID != nil {
		m["account_id"] = *rec.AccountID
	}
	if rec.PolicyID != nil {
		m["policy_id"] = *rec.PolicyID
	}
	if rec.State != nil {
		m["state"] = *rec.State
	}
	if rec.TargetTime != nil {
		m["target_time"] = *rec.TargetTime
	}
	if rec.CreatedAt != nil {
		m["created_at"] = rec.CreatedAt.String()
	}
	if rec.CreatedBy != nil {
		m["created_by"] = *rec.CreatedBy
	}
	if rec.UpdatedAt != nil {
		m["updated_at"] = rec.UpdatedAt.String()
	}
	if rec.UpdatedBy != nil {
		m["updated_by"] = *rec.UpdatedBy
	}
	// Flex.Flatten will safely deal with nil and non-string values
	m["custom_properties"] = flex.Flatten(rec.CustomProperties)
	return m
}
