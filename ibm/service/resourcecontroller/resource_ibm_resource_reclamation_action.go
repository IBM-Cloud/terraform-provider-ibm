// Copyright IBM Corp. 2024
// Licensed under the Mozilla Public License v2.0

package resourcecontroller

import (
	"context"
	"fmt"
	"log"
	"strings"

	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
)

// ResourceIBMResourceReclamationAction implements a Terraform resource to execute reclamation actions.
// Supported actions: "reclaim" (permanent delete), "restore" (recover resource).
//
// Example usage:
//
//	resource "ibm_resource_reclamation_action" "example" {
//	  id         = "reclamation-uuid"
//	  action     = "reclaim"        # or "restore"
//	  request_by = "user@example.com"  # optional
//	  comment    = "Performing reclamation action"  # optional
//	}
func ResourceIBMResourceReclamationAction() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMResourceReclamationActionCreate,
		ReadContext:   resourceIBMResourceReclamationActionRead,
		DeleteContext: resourceIBMResourceReclamationActionDelete,
		Description:   "Perform a reclamation action ('reclaim' or 'restore') on a resource reclamation.",

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The reclamation ID to perform the action on.",
			},
			"action": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "The reclamation action to perform: 'reclaim' or 'restore'.",
				ValidateFunc: validateReclamationAction,
			},
			"request_by": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The request initiator different from the authentication token (optional).",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A comment to describe the action (optional).",
			},

			// Computed fields from the returned reclamation
			"entity_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_type_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_crn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_instance_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"policy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"target_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"custom_properties": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

// validateReclamationAction validates the "action" input field.
func validateReclamationAction(val interface{}, key string) (warns []string, errs []error) {
	v := strings.ToLower(val.(string))
	if v != "reclaim" && v != "restore" {
		errs = append(errs, fmt.Errorf("%q must be either 'reclaim' or 'restore', got: %s", key, v))
	}
	return
}

// resourceIBMResourceReclamationActionCreate performs the reclamation action.
func resourceIBMResourceReclamationActionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(
			err,
			err.Error(),
			"(Resource) ibm_resource_reclamation_action",
			"create",
			"initialize-client",
		)
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	recID := d.Get("id").(string)
	action := strings.ToLower(d.Get("action").(string))

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

	reclamation, _, err := rsConClient.RunReclamationActionWithContext(ctx, actionOpts)
	if err != nil {
		tfErr := flex.TerraformErrorf(
			err,
			fmt.Sprintf("RunReclamationActionWithContext failed (%s): %s", action, err.Error()),
			"(Resource) ibm_resource_reclamation_action",
			"create",
		)
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if reclamation == nil || reclamation.ID == nil {
		d.SetId(recID)
	} else {
		d.SetId(*reclamation.ID)
	}

	flattenAndSetReclamation(d, reclamation)

	return nil
}

// resourceIBMResourceReclamationActionRead refreshes state by querying the reclamation out of the full list.
// If the reclamation is missing, removes it from state.
func resourceIBMResourceReclamationActionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(
			err,
			err.Error(),
			"(Resource) ibm_resource_reclamation_action",
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
		tfErr := flex.TerraformErrorf(
			err,
			fmt.Sprintf("ListReclamationsWithContext failed during read: %s", err.Error()),
			"(Resource) ibm_resource_reclamation_action",
			"read",
		)
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		// Assume resource gone, clear state
		d.SetId("")
		return nil
	}

	// Locate reclamation by ID
	for _, rec := range reclamationsList.Resources {
		if rec.ID != nil && *rec.ID == recID {
			flattenAndSetReclamation(d, &rec)
			return nil
		}
	}

	// Not found: resource likely deleted or reclaimed
	d.SetId("")
	return nil
}

// resourceIBMResourceReclamationActionDelete implements the Terraform Delete action by clearing state.
func resourceIBMResourceReclamationActionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		setField("id", *rec.ID)
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
