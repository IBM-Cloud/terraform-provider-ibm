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
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/BackupAndRecovery/brs-migration-orchestrator/brsmigrationv2"
)

func DataSourceIbmBrsMigration() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBrsMigrationRead,

		Schema: map[string]*schema.Schema{
			"migration_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The migration project ID (mgr-{uuid4} format).",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Human-readable name for this migration project.",
			},
			"brs_crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "CRN of the IBM Cloud Backup and Recovery instance backing this migration.",
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
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Optional human-readable description.",
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

func dataSourceIbmBrsMigrationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brsMigrationClient, err := meta.(conns.ClientSession).BrsMigrationV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_brs_migration", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getMigrationOptions := &brsmigrationv2.GetMigrationOptions{}

	getMigrationOptions.SetMigrationID(d.Get("migration_id").(string))

	migration, _, err := brsMigrationClient.GetMigrationWithContext(context, getMigrationOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetMigrationWithContext failed: %s", err.Error()), "(Data) ibm_brs_migration", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*getMigrationOptions.MigrationID)

	if err = d.Set("name", migration.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_brs_migration", "read", "set-name").GetDiag()
	}

	if err = d.Set("brs_crn", migration.BrsCrn); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting brs_crn: %s", err), "(Data) ibm_brs_migration", "read", "set-brs_crn").GetDiag()
	}

	if !core.IsNil(migration.Crn) {
		if err = d.Set("crn", migration.Crn); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_brs_migration", "read", "set-crn").GetDiag()
		}
	}

	if err = d.Set("state", migration.State); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting state: %s", err), "(Data) ibm_brs_migration", "read", "set-state").GetDiag()
	}

	if !core.IsNil(migration.Description) {
		if err = d.Set("description", migration.Description); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting description: %s", err), "(Data) ibm_brs_migration", "read", "set-description").GetDiag()
		}
	}

	if err = d.Set("created_at", flex.DateTimeToString(migration.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_brs_migration", "read", "set-created_at").GetDiag()
	}

	if !core.IsNil(migration.UpdatedAt) {
		if err = d.Set("updated_at", flex.DateTimeToString(migration.UpdatedAt)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting updated_at: %s", err), "(Data) ibm_brs_migration", "read", "set-updated_at").GetDiag()
		}
	}

	return nil
}
