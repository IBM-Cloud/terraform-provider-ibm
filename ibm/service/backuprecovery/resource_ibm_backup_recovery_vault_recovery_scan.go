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

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func ResourceIbmBackupRecoveryVaultRecoveryScan() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmBackupRecoveryVaultRecoveryScanCreate,
		ReadContext:   resourceIbmBackupRecoveryVaultRecoveryScanRead,
		UpdateContext: resourceIbmBackupRecoveryVaultRecoveryScanUpdate,
		DeleteContext: resourceIbmBackupRecoveryVaultRecoveryScanDelete,
		Importer:      &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			"x_ibm_tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Specifies the unique id of the tenant.",
			},
			"cloud_type": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Specifies the cloud type where the vault is registered for recovery scan. Currently, only 'ibm' is supported.",
			},
			"recovery_scan_request_params": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ForceNew:    true,
				Description: "Specifies the parameters specific to the Backup and Recovery instance. which is the vault.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vault_id": &schema.Schema{
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Specifies the unique id of the IBM Cloud Backup and Recovery instance for which the recovery scan is to be initiated.",
						},
					},
				},
			},
			"uid": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies the unique id of the recovery scan.",
			},
		},
	}
}

func ResourceIbmBackupRecoveryVaultRecoveryScanValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "cloud_type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "ibm",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_backup_recovery_vault_recovery_scan", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmBackupRecoveryVaultRecoveryScanCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_vault_recovery_scan", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	vaultRecoveryScanOptions := &backuprecoveryv1.VaultRecoveryScanOptions{}

	vaultRecoveryScanOptions.SetXIBMTenantID(d.Get("x_ibm_tenant_id").(string))
	vaultRecoveryScanOptions.SetCloudType(d.Get("cloud_type").(string))
	if _, ok := d.GetOk("recovery_scan_request_params"); ok {
		recoveryScanRequestParamsModel, err := ResourceIbmBackupRecoveryVaultRecoveryScanMapToRecoveryScanRequestParams(d.Get("recovery_scan_request_params.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_vault_recovery_scan", "create", "parse-recovery_scan_request_params").GetDiag()
		}
		vaultRecoveryScanOptions.SetRecoveryScanRequestParams(recoveryScanRequestParamsModel)
	}

	recoveryScan, _, err := backupRecoveryClient.VaultRecoveryScanWithContext(context, vaultRecoveryScanOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("VaultRecoveryScanWithContext failed: %s", err.Error()), "ibm_backup_recovery_vault_recovery_scan", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*recoveryScan.Uid)
	if err = d.Set("cloud_type", recoveryScan.CloudType); err != nil {
		err = fmt.Errorf("Error setting cloud_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_vault_recovery_scan", "read", "set-cloud_type").GetDiag()
	}
	if !core.IsNil(recoveryScan.RecoveryScanRequestParams) {
		recoveryScanRequestParamsMap, err := ResourceIbmBackupRecoveryVaultRecoveryScanRecoveryScanRequestParamsToMap(recoveryScan.RecoveryScanRequestParams)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_vault_recovery_scan", "read", "recovery_scan_request_params-to-map").GetDiag()
		}
		if err = d.Set("recovery_scan_request_params", []map[string]interface{}{recoveryScanRequestParamsMap}); err != nil {
			err = fmt.Errorf("Error setting recovery_scan_request_params: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_vault_recovery_scan", "read", "set-recovery_scan_request_params").GetDiag()
		}
	}
	if !core.IsNil(recoveryScan.Uid) {
		if err = d.Set("uid", recoveryScan.Uid); err != nil {
			err = fmt.Errorf("Error setting uid: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_vault_recovery_scan", "read", "set-uid").GetDiag()
		}
	}

	return resourceIbmBackupRecoveryVaultRecoveryScanRead(context, d, meta)
}

func resourceIbmBackupRecoveryVaultRecoveryScanRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return nil
}

func resourceIbmBackupRecoveryVaultRecoveryScanDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceIbmBackupRecoveryVaultRecoveryScanUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func ResourceIbmBackupRecoveryVaultRecoveryScanMapToRecoveryScanRequestParams(modelMap map[string]interface{}) (*backuprecoveryv1.RecoveryScanRequestParams, error) {
	model := &backuprecoveryv1.RecoveryScanRequestParams{}
	model.VaultID = core.Int64Ptr(int64(modelMap["vault_id"].(int)))
	return model, nil
}

func ResourceIbmBackupRecoveryVaultRecoveryScanRecoveryScanRequestParamsToMap(model *backuprecoveryv1.RecoveryScanRequestParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["vault_id"] = flex.IntValue(model.VaultID)
	return modelMap, nil
}
