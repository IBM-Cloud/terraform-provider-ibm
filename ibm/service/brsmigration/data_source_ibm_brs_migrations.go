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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/ibm-brs-migration-sdk-go/brsmigrationv1"
)

func DataSourceIbmBrsMigrations() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBrsMigrationsRead,

		Schema: map[string]*schema.Schema{
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter by migration state.",
			},
			"migrations": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of migration projects on this page.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Migration project ID (mgr-{uuid4} format).",
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
				},
			},
		},
	}
}

func dataSourceIbmBrsMigrationsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brsMigrationClient, err := meta.(conns.ClientSession).BrsMigrationV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_brs_migrations", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listMigrationsOptions := &brsmigrationv1.ListMigrationsOptions{}

	if _, ok := d.GetOk("state"); ok {
		listMigrationsOptions.SetState(d.Get("state").(string))
	}

	var pager *brsmigrationv1.MigrationsPager
	pager, err = brsMigrationClient.NewMigrationsPager(listMigrationsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_brs_migrations", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	allItems, err := pager.GetAll()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("MigrationsPager.GetAll() failed %s", err), "(Data) ibm_brs_migrations", "read")
		log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBrsMigrationsID(d))

	mapSlice := []map[string]interface{}{}
	for _, modelItem := range allItems {
		modelMap, err := DataSourceIbmBrsMigrationsMigrationToMap(&modelItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_brs_migrations", "read", "Migrations-to-map").GetDiag()
		}
		mapSlice = append(mapSlice, modelMap)
	}

	if err = d.Set("migrations", mapSlice); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting migrations %s", err), "(Data) ibm_brs_migrations", "read", "migrations-set").GetDiag()
	}

	return nil
}

// dataSourceIbmBrsMigrationsID returns a reasonable ID for the list.
func dataSourceIbmBrsMigrationsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBrsMigrationsMigrationToMap(model *brsmigrationv1.Migration) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	modelMap["brs_crn"] = *model.BrsCrn
	if model.Crn != nil {
		modelMap["crn"] = *model.Crn
	}
	modelMap["state"] = *model.State
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	modelMap["created_at"] = model.CreatedAt.String()
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	return modelMap, nil
}
