// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.113.1-d76630af-20260320-135953
 */

package backuprecovery

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func ResourceIbmBackupRecoveryProtectionSourceRefresh() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmBackupRecoveryProtectionSourceRefreshCreate,
		ReadContext:   resourceIbmBackupRecoveryProtectionSourceRefreshRead,
		UpdateContext: resourceIbmBackupRecoveryProtectionSourceRefreshUpdate,
		DeleteContext: resourceIbmBackupRecoveryProtectionSourceRefreshDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"x_ibm_tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Specifies the unique id of the tenant.",
			},
			"backup_recovery_protection_source_id": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "protection source Id",
			},
		},
	}
}

func resourceIbmBackupRecoveryProtectionSourceRefreshCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_protection_source_refresh", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	endpointType := d.Get("endpoint_type").(string)
	instanceId, region, serviceName := getInstanceIdAndRegion(d)
	if instanceId != "" && region != "" {
		bmxsession, err := meta.(conns.ClientSession).BluemixSession()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("unable to get clientSession"), "ibm_backup_recovery_protection_source_refresh", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		backupRecoveryClient = getClientWithInstanceEndpoint(backupRecoveryClient, bmxsession, instanceId, region, endpointType, serviceName)
	}

	refreshProtectionSourceByIdOptions := &backuprecoveryv1.RefreshProtectionSourceByIdOptions{}

	refreshProtectionSourceByIdOptions.SetXIBMTenantID(d.Get("x_ibm_tenant_id").(string))
	refreshProtectionSourceByIdOptions.SetID(int64(d.Get("backup_recovery_protection_source_id").(int)))

	_, err = backupRecoveryClient.RefreshProtectionSourceByIDWithContext(context, refreshProtectionSourceByIdOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("RefreshProtectionSourceByIDWithContext failed: %s", err.Error()), "ibm_backup_recovery_protection_source_refresh", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	registrationId := strconv.Itoa(int(*refreshProtectionSourceByIdOptions.ID))

	d.SetId(registrationId)

	return resourceIbmBackupRecoveryProtectionSourceRefreshRead(context, d, meta)
}

func resourceIbmBackupRecoveryProtectionSourceRefreshRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceIbmBackupRecoveryProtectionSourceRefreshDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.

	var diags diag.Diagnostics
	warning := diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Delete Not Supported",
		Detail:   "The resource definition will be only be removed from the terraform statefile. This resource cannot be deleted from the backend. ",
	}
	diags = append(diags, warning)
	d.SetId("")
	return diags
}

func resourceIbmBackupRecoveryProtectionSourceRefreshUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "update" operation.
	var diags diag.Diagnostics
	warning := diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Resource update will only affect terraform state and not the actual backend resource",
		Detail:   "Update operation for this resource is not supported and will only affect the terraform statefile. No changes will be made to the backend resource.",
	}
	diags = append(diags, warning)
	return diags
}
