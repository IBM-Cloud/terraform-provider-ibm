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

func ResourceIbmBackupRecoveryVaultFailover() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmBackupRecoveryVaultFailoverCreate,
		ReadContext:   resourceIbmBackupRecoveryVaultFailoverRead,
		DeleteContext: resourceIbmBackupRecoveryVaultFailoverDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"x_ibm_tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Specifies the unique id of the tenant.",
			},
			"cloud_type": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_backup_recovery_vault_failover", "cloud_type"),
				Description:  "Specifies the type of the Backup and Recovery instance. Currently, only 'ibm' is supported.",
			},
			"failover_request_params": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ForceNew:    true,
				Description: "Specifies the parameters specific to the Backup and Recovery instance. viz the vault.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vault_id": &schema.Schema{
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Specifies the unique id of the IBM Cloud Backup and Recovery instance for which the failover is to be initiated.",
						},
					},
				},
			},
			"uid": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies the unique id of the failover.",
			},
		},
	}
}

func ResourceIbmBackupRecoveryVaultFailoverValidator() *validate.ResourceValidator {
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

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_backup_recovery_vault_failover", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmBackupRecoveryVaultFailoverCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_vault_failover", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	initVaultsFailoverOptions := &backuprecoveryv1.InitVaultsFailoverOptions{}

	initVaultsFailoverOptions.SetXIBMTenantID(d.Get("x_ibm_tenant_id").(string))
	initVaultsFailoverOptions.SetCloudType(d.Get("cloud_type").(string))
	if _, ok := d.GetOk("failover_request_params"); ok {
		failoverRequestParamsModel, err := ResourceIbmBackupRecoveryVaultFailoverMapToFailoverRequestParams(d.Get("failover_request_params.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_vault_failover", "create", "parse-failover_request_params").GetDiag()
		}
		initVaultsFailoverOptions.SetFailoverRequestParams(failoverRequestParamsModel)
	}

	vaultFailover, _, err := backupRecoveryClient.InitVaultsFailoverWithContext(context, initVaultsFailoverOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("InitVaultsFailoverWithContext failed: %s", err.Error()), "ibm_backup_recovery_vault_failover", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*vaultFailover.Uid)

	if err = d.Set("cloud_type", vaultFailover.CloudType); err != nil {
		err = fmt.Errorf("Error setting cloud_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_vault_failover", "read", "set-cloud_type").GetDiag()
	}
	if !core.IsNil(vaultFailover.FailoverRequestParams) {
		failoverRequestParamsMap, err := ResourceIbmBackupRecoveryVaultFailoverFailoverRequestParamsToMap(vaultFailover.FailoverRequestParams)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_vault_failover", "read", "failover_request_params-to-map").GetDiag()
		}
		if err = d.Set("failover_request_params", []map[string]interface{}{failoverRequestParamsMap}); err != nil {
			err = fmt.Errorf("Error setting failover_request_params: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_vault_failover", "read", "set-failover_request_params").GetDiag()
		}
	}

	return resourceIbmBackupRecoveryVaultFailoverRead(context, d, meta)
}

func resourceIbmBackupRecoveryVaultFailoverRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_vault_failover", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getBatchVaultsFailoverStatusOptions := &backuprecoveryv1.GetBatchVaultsFailoverStatusOptions{}

	getBatchVaultsFailoverStatusOptions.SetXIBMTenantID(d.Get("x_ibm_tenant_id").(string))

	getBatchVaultFailoverStatus, response, err := backupRecoveryClient.GetBatchVaultsFailoverStatusWithContext(context, getBatchVaultsFailoverStatusOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBatchVaultsFailoverStatusWithContext failed: %s", err.Error()), "ibm_backup_recovery_vault_failover", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if !core.IsNil(getBatchVaultFailoverStatus.ErrorMessage) {
		if err = d.Set("error_message", getBatchVaultFailoverStatus.ErrorMessage); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting error_message: %s", err), "(Data) ibm_backup_recovery_vault_failover_status", "read", "set-error_message").GetDiag()
		}
	}

	if !core.IsNil(getBatchVaultFailoverStatus.FailoverStatuses) {
		failoverStatuses := []map[string]interface{}{}
		for _, failoverStatusesItem := range getBatchVaultFailoverStatus.FailoverStatuses {
			failoverStatusesItemMap, err := DataSourceIbmBackupRecoveryVaultFailoverStatusBatchVaultFailoverStatusToMap(&failoverStatusesItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_vault_failover_status", "read", "failover_statuses-to-map").GetDiag()
			}
			failoverStatuses = append(failoverStatuses, failoverStatusesItemMap)
		}
		if err = d.Set("failover_statuses", failoverStatuses); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting failover_statuses: %s", err), "(Data) ibm_backup_recovery_vault_failover_status", "read", "set-failover_statuses").GetDiag()
		}
	}

	return nil
}

func resourceIbmBackupRecoveryVaultFailoverDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func ResourceIbmBackupRecoveryVaultFailoverMapToFailoverRequestParams(modelMap map[string]interface{}) (*backuprecoveryv1.FailoverRequestParams, error) {
	model := &backuprecoveryv1.FailoverRequestParams{}
	model.VaultID = core.Int64Ptr(int64(modelMap["vault_id"].(int)))
	return model, nil
}

func ResourceIbmBackupRecoveryVaultFailoverFailoverRequestParamsToMap(model *backuprecoveryv1.FailoverRequestParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["vault_id"] = flex.IntValue(model.VaultID)
	return modelMap, nil
}
