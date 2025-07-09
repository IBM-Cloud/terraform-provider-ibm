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

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func ResourceIbmBackupRecoveryConnectorAccessToken() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmBackupRecoveryConnectorAccessTokenCreate,
		ReadContext:   resourceIbmBackupRecoveryConnectorAccessTokenRead,
		DeleteContext: resourceIbmBackupRecoveryConnectorAccessTokenDelete,
		CustomizeDiff: checkDiffResourceIbmBackupRecoveryConnectorAccessToken,
		UpdateContext: resourceIbmBackupRecoveryConnectorAccessTokenUpdate,

		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the login name of the Cohesity user.",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Specifies the password of the Cohesity user account.",
			},
			"domain": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the domain the user is logging in to. For a local user the domain is LOCAL. For LDAP/AD user, the domain will map to a LDAP connection string. A user is uniquely identified by a combination of username and domain. LOCAL is the default domain.",
			},
			"access_token": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Created access token.",
				Sensitive:   true,
			},
			"privileges": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Privileges for the user.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"token_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Access token type.",
			},
		},
	}
}

func checkDiffResourceIbmBackupRecoveryConnectorAccessToken(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
	if d.Id() == "" {
		return nil
	}

	for fieldName := range ResourceIbmBackupRecoveryConnectorAccessToken().Schema {
		if d.HasChange(fieldName) {
			return fmt.Errorf("[ERROR] Resource ibm_backup_recovery_connector_access_token cannot be updated.")
		}
	}
	return nil
}

func resourceIbmBackupRecoveryConnectorAccessTokenCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryConnectorClient, err := meta.(conns.ClientSession).BackupRecoveryV1Connector()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_access_token", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if backupRecoveryConnectorClient.GetConnectorURL() == "" {
		tfErr := flex.DiscriminatedTerraformErrorf(nil, "No connector URL specified. Please set the `IBMCLOUD_BACKUP_RECOVERY_CONNECTOR_ENDPOINT` environment variable or specify the endpoint in endpoints.json file.", "ibm_backup_recovery_connector_access_token", "create", "initialize-client")
		return tfErr.GetDiag()
	}

	createAccessTokenOptions := &backuprecoveryv1.CreateAccessTokenOptions{}

	if _, ok := d.GetOk("username"); ok {
		createAccessTokenOptions.SetUsername(d.Get("username").(string))
	}
	if _, ok := d.GetOk("password"); ok {
		createAccessTokenOptions.SetPassword(d.Get("password").(string))
	}
	if _, ok := d.GetOk("domain"); ok {
		createAccessTokenOptions.SetDomain(d.Get("domain").(string))
	}

	result, _, err := backupRecoveryConnectorClient.CreateAccessTokenWithContext(context, createAccessTokenOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateAccessTokenWithContext failed: %s", err.Error()), "ibm_backup_recovery_connector_access_token", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(resourceIbmBackupRecoveryConnectionAccessTokenID(d))

	if !core.IsNil(result.AccessToken) {
		if err = d.Set("access_token", result.AccessToken); err != nil {
			err = fmt.Errorf("Error setting access_token: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_access_token", "read", "set-access_token").GetDiag()
		}
	}
	if !core.IsNil(result.Privileges) {
		if err = d.Set("privileges", result.Privileges); err != nil {
			err = fmt.Errorf("Error setting privileges: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_access_token", "read", "set-privileges").GetDiag()
		}
	}
	if !core.IsNil(result.TokenType) {
		if err = d.Set("token_type", result.TokenType); err != nil {
			err = fmt.Errorf("Error setting token_type: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_access_token", "read", "set-token_type").GetDiag()
		}
	}

	return resourceIbmBackupRecoveryConnectorAccessTokenRead(context, d, meta)
}

func resourceIbmBackupRecoveryConnectionAccessTokenID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func resourceIbmBackupRecoveryConnectorAccessTokenUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceIbmBackupRecoveryConnectorAccessTokenRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return nil
}

func resourceIbmBackupRecoveryConnectorAccessTokenDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
