// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.114.3-943fbc81-20260603-173645
 */

package backuprecovery

import (
	"context"
	"fmt"
	"log"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func ResourceIbmBackupRecoveryConnectorAgentUpgrade() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmBackupRecoveryConnectorAgentUpgradeCreate,
		ReadContext:   resourceIbmBackupRecoveryConnectorAgentUpgradeRead,
		UpdateContext: resourceIbmBackupRecoveryConnectorAgentUpgradeUpdate,
		DeleteContext: resourceIbmBackupRecoveryConnectorAgentUpgradeDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"tenant_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Specifies the tenants whose connector agents should be upgraded (all connections under each listed tenant).",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"connection_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Specifies the specific connections whose connector agents should be upgraded.",
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			"retry_forever": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				ForceNew:    true,
				Description: "When true or omitted, the upgrade is retried indefinitely until it succeeds. When false, the cluster retries up to 3 times before giving up.",
			},
			"upgrade_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "UUID assigned to this upgrade run.",
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies the overall state of this upgrade run. 'NotStarted' indicates the run has not yet begun. 'InProgress' indicates at least one tenant is still upgrading. 'Failed' indicates no tenant is in progress and at least one has failed. 'Succeeded' indicates every tenant has succeeded.",
			},
		},
	}
}

func resourceIbmBackupRecoveryConnectorAgentUpgradeCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_agent_upgrade", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	upgradeConnectorAgentsOptions := &backuprecoveryv1.UpgradeConnectorAgentsOptions{}

	if _, ok := d.GetOk("tenant_ids"); ok {
		var tenantIds []string
		for _, v := range d.Get("tenant_ids").([]interface{}) {
			tenantIds = append(tenantIds, v.(string))
		}
		upgradeConnectorAgentsOptions.SetTenantIds(tenantIds)
	}
	if _, ok := d.GetOk("connection_ids"); ok {
		var connectionIds []int64
		for _, v := range d.Get("connection_ids").([]interface{}) {
			connectionIds = append(connectionIds, int64(v.(int)))
		}
		upgradeConnectorAgentsOptions.SetConnectionIds(connectionIds)
	}
	upgradeConnectorAgentsOptions.SetRetryForever(d.Get("retry_forever").(bool))

	upgradeConnectorAgentsResponseBody, _, err := backupRecoveryClient.UpgradeConnectorAgentsWithContext(context, upgradeConnectorAgentsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpgradeConnectorAgentsWithContext failed: %s", err.Error()), "ibm_backup_recovery_connector_agent_upgrade", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(string(*upgradeConnectorAgentsResponseBody.UpgradeID))

	return resourceIbmBackupRecoveryConnectorAgentUpgradeRead(context, d, meta)
}

func resourceIbmBackupRecoveryConnectorAgentUpgradeRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_agent_upgrade", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getConnectorAgentsUpgradeStatusOptions := &backuprecoveryv1.GetConnectorAgentsUpgradeStatusOptions{}
	getConnectorAgentsUpgradeStatusOptions.SetUpgradeID(core.UUIDPtr(strfmt.UUID(d.Id())))

	connectorAgentsUpgradeStatus, response, err := backupRecoveryClient.GetConnectorAgentsUpgradeStatusWithContext(context, getConnectorAgentsUpgradeStatusOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetConnectorAgentsUpgradeStatusWithContext failed: %s", err.Error()), "ibm_backup_recovery_connector_agent_upgrade", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("upgrade_id", d.Id()); err != nil {
		err = fmt.Errorf("Error setting upgrade_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_agent_upgrade", "read", "set-upgrade_id").GetDiag()
	}
	if err = d.Set("status", connectorAgentsUpgradeStatus.Status); err != nil {
		err = fmt.Errorf("Error setting status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_agent_upgrade", "read", "set-status").GetDiag()
	}

	return nil
}

func resourceIbmBackupRecoveryConnectorAgentUpgradeDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceIbmBackupRecoveryConnectorAgentUpgradeUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
