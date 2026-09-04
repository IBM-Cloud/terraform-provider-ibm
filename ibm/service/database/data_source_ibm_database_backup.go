// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
)

type dataSourceIBMDatabaseBackupBackend interface {
	Read(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics
}

func pickDataSourceBackupBackend(d *schema.ResourceData, meta interface{}) (dataSourceIBMDatabaseBackupBackend, error) {
	backupID := d.Get("backup_id").(string)

	// Gen2 only uses decoupled Independent Backup, identifiable from the CRN's
	// service-name segment (5th colon-delimited field) alone. SplitN is bounded
	// to 6 fields so we avoid splitting the remainder of the CRN (region, scope,
	// instance/resource IDs) that we don't need.
	parts := strings.SplitN(backupID, ":", 6)

	if len(parts) >= 5 && parts[4] == "databases-independent-backups" {
		return newDataSourceIBMDatabaseBackupGen2Backend(), nil
	}

	// Reject coupled backups whose source instance is Gen2.
	if err := rejectCoupledBackupFromGen2Instance(backupID, meta); err != nil {
		return nil, err
	}

	// All other backup IDs are Classic — route directly to the classic backend.
	return newDataSourceIBMDatabaseBackupClassicBackend(), nil
}

// rejectCoupledBackupFromGen2Instance errors if backupID belongs to a Gen2 instance.
// Unresolvable instances are allowed through; the ICD API rejects them server-side.
func rejectCoupledBackupFromGen2Instance(backupID string, meta interface{}) error {
	if isGen2CoupledBackup(backupID, meta) {
		return fmt.Errorf("Gen2 instances only support Independent Backup; use the Independent Backup CRN (databases-independent-backups) instead")
	}
	return nil
}

// isGen2CoupledBackup returns true only when the backup CRN can be resolved to a Gen2 instance.
// Returns false whenever any lookup step fails.
func isGen2CoupledBackup(backupID string, meta interface{}) bool {
	instanceCRN, err := instanceCRNFromCoupledBackupCRN(backupID)
	if err != nil {
		return false
	}

	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return false
	}

	instance, _, err := rsConClient.GetResourceInstance(&rc.GetResourceInstanceOptions{ID: &instanceCRN})
	if err != nil || instance.ResourcePlanID == nil {
		return false
	}

	return isGen2Plan(*instance.ResourcePlanID)
}

func DataSourceIBMDatabaseBackup() *schema.Resource {
	return &schema.Resource{
		ReadContext: DataSourceIBMDatabaseBackupRead,

		Schema: map[string]*schema.Schema{
			"backup_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Backup ID.",
			},
			"deployment_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the deployment this backup relates to.",
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of backup.",
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of this backup.",
			},
			"is_downloadable": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Is this backup available to download?.",
			},
			"is_restorable": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Can this backup be used to restore an instance?.",
			},
			"download_link": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI which is currently available for file downloading.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date and time when this backup was created.",
			},
		},
	}
}

func DataSourceIBMDatabaseBackupRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	b, err := pickDataSourceBackupBackend(d, meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_database_backup", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return b.Read(context, d, meta)
}

type dataSourceIBMDatabaseBackupClassicBackend struct{}

func newDataSourceIBMDatabaseBackupClassicBackend() dataSourceIBMDatabaseBackupBackend {
	return &dataSourceIBMDatabaseBackupClassicBackend{}
}

func (c *dataSourceIBMDatabaseBackupClassicBackend) Read(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cloudDatabasesClient, err := meta.(conns.ClientSession).CloudDatabasesV5()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_database_backup", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getBackupInfoOptions := &clouddatabasesv5.GetBackupInfoOptions{}
	getBackupInfoOptions.SetBackupID(d.Get("backup_id").(string))

	backup, response, err := cloudDatabasesClient.GetBackupInfoWithContext(context, getBackupInfoOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBackupInfoWithContext failed: %s\n%s", err.Error(), response), "(Data) ibm_database_backup", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*backup.Backup.ID)

	if err = d.Set("backup_id", backup.Backup.ID); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting backup_id: %s", err), "(Data) ibm_database_backup", "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("deployment_id", backup.Backup.DeploymentID); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting deployment_id: %s", err), "(Data) ibm_database_backup", "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("type", backup.Backup.Type); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting type: %s", err), "(Data) ibm_database_backup", "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("status", backup.Backup.Status); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_database_backup", "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("is_downloadable", backup.Backup.IsDownloadable); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting is_downloadable: %s", err), "(Data) ibm_database_backup", "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("is_restorable", backup.Backup.IsRestorable); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting is_restorable: %s", err), "(Data) ibm_database_backup", "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("download_link", backup.Backup.DownloadLink); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting download_link: %s", err), "(Data) ibm_database_backup", "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("created_at", flex.DateTimeToString(backup.Backup.CreatedAt)); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_database_backup", "read")
		return tfErr.GetDiag()
	}

	return nil
}
