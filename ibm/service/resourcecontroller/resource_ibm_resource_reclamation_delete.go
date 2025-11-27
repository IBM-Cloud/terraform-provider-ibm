// Copyright IBM Corp. 2024
// Licensed under the Mozilla Public License v2.0

package resourcecontroller

import (
	"context"
	"fmt"
	"log"

	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
)

// ResourceIBMResourceReclamationDelete implements a Terraform resource to permanently delete reclaimed resources.
// This resource performs the equivalent of "ibmcloud resource reclamation-delete" command.
//
// Example usage:
//
//	resource "ibm_resource_reclamation_delete" "example" {
//	  reclamation_id = "reclamation-uuid"
//	  request_by     = "user@example.com"      # optional
//	  comment        = "Permanent deletion"    # optional
//	}
func ResourceIBMResourceReclamationDelete() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMResourceReclamationDeleteCreate,
		ReadContext:   resourceIBMResourceReclamationDeleteRead,
		DeleteContext: resourceIBMResourceReclamationDeleteDelete,
		Description:   "Permanently delete a reclaimed resource. This action is irreversible and equivalent to 'ibmcloud resource reclamation-delete'.",

		Schema: map[string]*schema.Schema{
			"reclamation_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The reclamation ID to permanently delete. This is the reclamation ID, not the resource instance ID.",
			},
			"request_by": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The request initiator different from the authentication token (optional).",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "A comment to describe the deletion action (optional).",
			},

			// Computed fields from the returned reclamation response
			"entity_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The entity ID of the reclaimed resource.",
			},
			"entity_type_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The entity type ID of the reclaimed resource.",
			},
			"entity_crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN of the reclaimed resource.",
			},
			"resource_instance_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource instance ID of the reclaimed resource.",
			},
			"resource_group_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group ID of the reclaimed resource.",
			},
			"account_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The account ID that owns the reclaimed resource.",
			},
			"policy_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The policy ID associated with the reclamation.",
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The current state of the reclamation.",
			},
			"target_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The target time when the reclamation was scheduled.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp when the reclamation was created.",
			},
			"created_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user who created the reclamation.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp when the reclamation was last updated.",
			},
			"updated_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user who last updated the reclamation.",
			},
			"custom_properties": {
				Type:        schema.TypeMap,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Custom properties associated with the reclamation.",
			},
		},
	}
}

// resourceIBMResourceReclamationDeleteCreate permanently deletes the reclaimed resource.
func resourceIBMResourceReclamationDeleteCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(
			err,
			err.Error(),
			"(Resource) ibm_resource_reclamation_delete",
			"create",
			"initialize-client",
		)
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	recID := d.Get("reclamation_id").(string)
	action := "reclaim" // Always "reclaim" for permanent deletion

	actionOpts := &rc.RunReclamationActionOptions{
		ID:         &recID,
		ActionName: &action,
	}

	if v, ok := d.GetOk("request_by"); ok {
		str := v.(string)
		actionOpts.RequestBy = &str
	}
	if v, ok := d.GetOk("comment"); ok {
		str := v.(string)
		actionOpts.Comment = &str
	}

	log.Printf("[INFO] Permanently deleting reclamation: %s", recID)

	reclamation, _, err := rsConClient.RunReclamationActionWithContext(ctx, actionOpts)
	if err != nil {
		tfErr := flex.TerraformErrorf(
			err,
			fmt.Sprintf("RunReclamationActionWithContext failed for permanent deletion: %s", err.Error()),
			"(Resource) ibm_resource_reclamation_delete",
			"create",
		)
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if reclamation == nil || reclamation.ID == nil {
		tfErr := flex.TerraformErrorf(
			fmt.Errorf("reclamation response is nil or missing ID"),
			"Permanent deletion response was empty",
			"(Resource) ibm_resource_reclamation_delete",
			"create",
		)
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*reclamation.ID)
	flattenAndSetReclamation(d, reclamation)

	log.Printf("[INFO] Successfully deleted reclamation: %s", *reclamation.ID)

	return nil
}

// resourceIBMResourceReclamationDeleteRead refreshes state by querying the reclamation.
// Since this is an action resource, we keep state even if reclamation is gone (expected after deletion).
func resourceIBMResourceReclamationDeleteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(
			err,
			err.Error(),
			"(Resource) ibm_resource_reclamation_delete",
			"read",
			"initialize-client",
		)
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	recID := d.Id()
	opts := &rc.ListReclamationsOptions{}

	reclamationsList, _, err := rsConClient.ListReclamationsWithContext(ctx, opts)
	if err != nil {
		log.Printf("[WARN] ListReclamationsWithContext failed during read: %s", err.Error())
		// Keep state - the deletion action was performed, API error doesn't change that
		return nil
	}

	// Search for the reclamation by ID
	for _, rec := range reclamationsList.Resources {
		if rec.ID != nil && *rec.ID == recID {
			flattenAndSetReclamation(d, &rec)
			return nil
		}
	}

	// Not found: This is expected - the deletion action succeeded
	// Keep the state to prevent Terraform from trying to recreate
	log.Printf("[INFO] Reclamation %s not found (expected after successful deletion), keeping state", recID)
	return nil
}

// resourceIBMResourceReclamationDeleteDelete implements the Terraform Delete action.
// Since this resource represents a deletion action, we just clear the state.
func resourceIBMResourceReclamationDeleteDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// The actual deletion happened during Create
	// This Delete operation just removes the resource from state
	log.Printf("[INFO] Removing reclamation delete resource from state: %s", d.Id())
	d.SetId("")
	return nil
}

// flattenAndSetReclamation sets the reclamation info fields from SDK struct into the Terraform state.
func flattenAndSetReclamation(d *schema.ResourceData, rec *rc.Reclamation) {
	if rec == nil {
		return
	}

	setField := func(fieldName string, value interface{}) {
		if err := d.Set(fieldName, value); err != nil {
			log.Printf("[WARN] Failed to set field %q: %v", fieldName, err)
		}
	}

	if rec.ID != nil {
		setField("reclamation_id", *rec.ID)
	}
	if rec.EntityID != nil {
		setField("entity_id", *rec.EntityID)
	}
	if rec.EntityTypeID != nil {
		setField("entity_type_id", *rec.EntityTypeID)
	}
	if rec.EntityCRN != nil {
		setField("entity_crn", *rec.EntityCRN)
	}
	if rec.ResourceInstanceID != nil {
		setField("resource_instance_id", *rec.ResourceInstanceID)
	}
	if rec.ResourceGroupID != nil {
		setField("resource_group_id", *rec.ResourceGroupID)
	}
	if rec.AccountID != nil {
		setField("account_id", *rec.AccountID)
	}
	if rec.PolicyID != nil {
		setField("policy_id", *rec.PolicyID)
	}
	if rec.State != nil {
		setField("state", *rec.State)
	}
	if rec.TargetTime != nil {
		setField("target_time", *rec.TargetTime)
	}
	if rec.CreatedAt != nil {
		setField("created_at", rec.CreatedAt.String())
	}
	if rec.CreatedBy != nil {
		setField("created_by", *rec.CreatedBy)
	}
	if rec.UpdatedAt != nil {
		setField("updated_at", rec.UpdatedAt.String())
	}
	if rec.UpdatedBy != nil {
		setField("updated_by", *rec.UpdatedBy)
	}

	setField("custom_properties", flex.Flatten(rec.CustomProperties))
}
