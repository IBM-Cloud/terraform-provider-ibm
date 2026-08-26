// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.114.3-943fbc81-20260603-173645
*/

package brsmigration

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/BackupAndRecovery/brs-migration-orchestrator/brsmigrationv2"
)

func ResourceIbmBrsMigration() *schema.Resource {
	return &schema.Resource{
		CreateContext:   resourceIbmBrsMigrationCreate,
		ReadContext:     resourceIbmBrsMigrationRead,
		UpdateContext:   resourceIbmBrsMigrationUpdate,
		DeleteContext:   resourceIbmBrsMigrationDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: validate.InvokeValidator("ibm_brs_migration", "name"),
				Description: "Human-readable name for this migration project.",
			},
			"brs_crn": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: validate.InvokeValidator("ibm_brs_migration", "brs_crn"),
				Description: "CRN of the IBM Cloud Backup and Recovery instance backing this migration.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ValidateFunc: validate.InvokeValidator("ibm_brs_migration", "description"),
				Description: "Optional human-readable description.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Server-assigned CRN for this migration resource.",
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current lifecycle state of the migration project.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this migration was created.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of the last update to this migration.",
			},
		},
	}
}

func ResourceIbmBrsMigrationValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^.+$`,
			MinValueLength:             1,
			MaxValueLength:             1024,
		},
		validate.ValidateSchema{
			Identifier:                 "brs_crn",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^crn:.+$`,
			MinValueLength:             1,
			MaxValueLength:             1024,
		},
		validate.ValidateSchema{
			Identifier:                 "description",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^.+$`,
			MinValueLength:             1,
			MaxValueLength:             1024,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_brs_migration", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmBrsMigrationCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brsMigrationClient, err := meta.(conns.ClientSession).BrsMigrationV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createMigrationOptions := &brsmigrationv2.CreateMigrationOptions{}

	createMigrationOptions.SetName(d.Get("name").(string))
	createMigrationOptions.SetBrsCrn(d.Get("brs_crn").(string))
	if _, ok := d.GetOk("description"); ok {
		createMigrationOptions.SetDescription(d.Get("description").(string))
	}

	migration, _, err := brsMigrationClient.CreateMigrationWithContext(context, createMigrationOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateMigrationWithContext failed: %s", err.Error()), "ibm_brs_migration", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*migration.ID)

	return resourceIbmBrsMigrationRead(context, d, meta)
}

func resourceIbmBrsMigrationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brsMigrationClient, err := meta.(conns.ClientSession).BrsMigrationV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getMigrationOptions := &brsmigrationv2.GetMigrationOptions{}

	getMigrationOptions.SetMigrationID(d.Id())

	migration, response, err := brsMigrationClient.GetMigrationWithContext(context, getMigrationOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetMigrationWithContext failed: %s", err.Error()), "ibm_brs_migration", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("name", migration.Name); err != nil {
		err = fmt.Errorf("Error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration", "read", "set-name").GetDiag()
	}
	if err = d.Set("brs_crn", migration.BrsCrn); err != nil {
		err = fmt.Errorf("Error setting brs_crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration", "read", "set-brs_crn").GetDiag()
	}
	if !core.IsNil(migration.Description) {
		if err = d.Set("description", migration.Description); err != nil {
			err = fmt.Errorf("Error setting description: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration", "read", "set-description").GetDiag()
		}
	}
	if !core.IsNil(migration.Crn) {
		if err = d.Set("crn", migration.Crn); err != nil {
			err = fmt.Errorf("Error setting crn: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration", "read", "set-crn").GetDiag()
		}
	}
	if err = d.Set("state", migration.State); err != nil {
		err = fmt.Errorf("Error setting state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration", "read", "set-state").GetDiag()
	}
	if err = d.Set("created_at", flex.DateTimeToString(migration.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration", "read", "set-created_at").GetDiag()
	}
	if !core.IsNil(migration.UpdatedAt) {
		if err = d.Set("updated_at", flex.DateTimeToString(migration.UpdatedAt)); err != nil {
			err = fmt.Errorf("Error setting updated_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration", "read", "set-updated_at").GetDiag()
		}
	}

	return nil
}

func resourceIbmBrsMigrationUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brsMigrationClient, err := meta.(conns.ClientSession).BrsMigrationV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateMigrationOptions := &brsmigrationv2.UpdateMigrationOptions{}

	updateMigrationOptions.SetMigrationID(d.Id())

	hasChange := false

	if d.HasChange("name") {
		updateMigrationOptions.SetName(d.Get("name").(string))
		hasChange = true
	}
	if d.HasChange("description") {
		updateMigrationOptions.SetDescription(d.Get("description").(string))
		hasChange = true
	}

	if hasChange {
		_, _, err = brsMigrationClient.UpdateMigrationWithContext(context, updateMigrationOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateMigrationWithContext failed: %s", err.Error()), "ibm_brs_migration", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIbmBrsMigrationRead(context, d, meta)
}

func resourceIbmBrsMigrationDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brsMigrationClient, err := meta.(conns.ClientSession).BrsMigrationV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteMigrationOptions := &brsmigrationv2.DeleteMigrationOptions{}

	deleteMigrationOptions.SetMigrationID(d.Id())

	_, err = brsMigrationClient.DeleteMigrationWithContext(context, deleteMigrationOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteMigrationWithContext failed: %s", err.Error()), "ibm_brs_migration", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}
