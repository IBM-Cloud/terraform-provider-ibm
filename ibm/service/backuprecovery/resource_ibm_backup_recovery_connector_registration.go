// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
)

func ResourceIbmBackupRecoveryConnectorRegistration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmBackupRecoveryConnectorRegistrationCreate,
		ReadContext:   resourceIbmBackupRecoveryConnectorRegistrationRead,
		DeleteContext: resourceIbmBackupRecoveryConnectorRegistrationDelete,
		Importer:      &schema.ResourceImporter{},
		CustomizeDiff: checkDiffResourceIbmBackupRecoveryConnectorRegistration,
		UpdateContext: resourceIbmBackupRecoveryConnectorRegistrationUpdate,
		Schema: map[string]*schema.Schema{
			"connector_id": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The connector's ID to be used for registration. Two connectors belonging to the same tenant are guaranteed to have different IDs.",
			},
			"access_token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Token required to authenticate to the connector. Token can be obtained using ibm_backup_recovery_connector_access_token resource",
			},
			"registration_token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The registration token.",
			},
		},
	}
}

func checkDiffResourceIbmBackupRecoveryConnectorRegistration(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
	if d.Id() == "" {
		return nil
	}

	for fieldName := range ResourceIbmBackupRecoveryConnectorRegistration().Schema {
		if d.HasChange(fieldName) {
			return fmt.Errorf("[ERROR] Resource ibm_backup_recovery_connector_registration cannot be updated.")
		}
	}
	return nil
}

func resourceIbmBackupRecoveryConnectorRegistrationCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryConnectorClient, err := meta.(conns.ClientSession).BackupRecoveryV1Connector()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_registration", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if backupRecoveryConnectorClient.GetConnectorURL() == "" {
		tfErr := flex.DiscriminatedTerraformErrorf(nil, "No connector URL specified. Please set the `IBMCLOUD_BACKUP_RECOVERY_CONNECTOR_ENDPOINT` environment variable or specify the endpoint in endpoints.json file.", "ibm_backup_recovery_data_source_connector_registration", "create", "initialize-client")
		return tfErr.GetDiag()
	}

	accessToken := d.Get("access_token").(string)
	var auth core.Authenticator
	auth = &core.BearerTokenAuthenticator{BearerToken: accessToken}
	backupRecoveryConnectorClient.Service.Options.Authenticator = auth

	registerDataSourceConnectorOptions := &backuprecoveryv1.RegisterDataSourceConnectorOptions{}

	registerDataSourceConnectorOptions.SetRegistrationToken(d.Get("registration_token").(string))
	if _, ok := d.GetOk("connector_id"); ok {
		registerDataSourceConnectorOptions.SetConnectorID(int64(d.Get("connector_id").(int)))
	}

	_, err = backupRecoveryConnectorClient.RegisterDataSourceConnectorWithContext(context, registerDataSourceConnectorOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("RegisterDataSourceConnectorWithContext failed: %s", err.Error()), "ibm_backup_recovery_connector_registration", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(resourceIbmBackupRecoveryConnectorRegistrationId(d))

	return resourceIbmBackupRecoveryConnectorRegistrationRead(context, d, meta)
}

func resourceIbmBackupRecoveryConnectorRegistrationId(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func resourceIbmBackupRecoveryConnectorRegistrationUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	var diags diag.Diagnostics
	warning := diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Resource update will only affect terraform state and not the actual backend resource",
		Detail:   "Update operation for this resource is not supported and will only affect the terraform statefile. No changes will be made to the backend resource.",
	}
	diags = append(diags, warning)
	// d.SetId("")
	return diags
}

func resourceIbmBackupRecoveryConnectorRegistrationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return nil
}

func resourceIbmBackupRecoveryConnectorRegistrationDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
