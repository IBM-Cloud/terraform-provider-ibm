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
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func ResourceIbmBackupRecoveryConnectorPasswordReset() *schema.Resource {
	return &schema.Resource{
		CreateContext: ResourceIbmBackupRecoveryConnectorPasswordResetCreate,
		ReadContext:   ResourceIbmBackupRecoveryConnectorPasswordResetRead,
		DeleteContext: ResourceIbmBackupRecoveryConnectorPasswordResetDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Specifies the login name of the connector user.",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Specifies the current password of the connector user.",
			},
			"new_password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Specifies the new password for the connector user.",
			},
			"session_name_cookie": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Specifies the session name cookie of the Cohesity user.",
			},
		},
	}
}

func ResourceIbmBackupRecoveryConnectorPasswordResetCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	getUsersOptions := &backuprecoveryv1.GetUsersOptions{}

	// if _, ok := d.GetOk("username"); ok {
	// 	getUsersOptions.SetUsername(d.Get("username").(string))
	// }

	if _, ok := d.GetOk("session_name_cookie"); ok {
		headerMap := map[string]string{"Cookie": d.Get("session_name_cookie").(string)}
		getUsersOptions.SetHeaders(headerMap)
	}

	res, _, err := backupRecoveryConnectorClient.GetUsersWithContext(context, getUsersOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetUsersWithContext failed: %s", err.Error()), "ibm_backup_recovery_reset_password", "getUser")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateUserOptions := &backuprecoveryv1.UpdateUserOptions{}
	updateUserOptions.AdUserInfo = res[0].AdUserInfo
	updateUserOptions.AdditionalGroupNames = res[0].AdditionalGroupNames
	updateUserOptions.AllowDsoModify = res[0].AllowDsoModify
	updateUserOptions.AuditLogSettings = res[0].AuditLogSettings
	updateUserOptions.AuthenticationType = res[0].AuthenticationType
	updateUserOptions.ClusterIdentifiers = res[0].ClusterIdentifiers
	updateUserOptions.CreatedTimeMsecs = res[0].CreatedTimeMsecs
	currentPassword := d.Get("password").(string)
	updateUserOptions.CurrentPassword = &currentPassword
	updateUserOptions.Description = res[0].Description
	updateUserOptions.Domain = res[0].Domain
	updateUserOptions.EffectiveTimeMsecs = res[0].EffectiveTimeMsecs
	updateUserOptions.EmailAddress = res[0].EmailAddress
	forcepasswordchange := false
	updateUserOptions.ForcePasswordChange = &forcepasswordchange
	updateUserOptions.ExpiredTimeMsecs = res[0].ExpiredTimeMsecs
	updateUserOptions.GoogleAccount = res[0].GoogleAccount
	updateUserOptions.IdpUserInfo = res[0].IdpUserInfo
	updateUserOptions.IntercomMessengerToken = res[0].IntercomMessengerToken
	updateUserOptions.IsAccountLocked = res[0].IsAccountLocked
	updateUserOptions.IsActive = res[0].IsActive
	updateUserOptions.LastSuccessfulLoginTimeMsecs = res[0].LastSuccessfulLoginTimeMsecs
	updateUserOptions.LastUpdatedTimeMsecs = res[0].LastUpdatedTimeMsecs
	updateUserOptions.MfaInfo = res[0].MfaInfo
	updateUserOptions.MfaMethods = res[0].MfaMethods
	updateUserOptions.ObjectClass = res[0].ObjectClass
	newPassword := d.Get("new_password").(string)
	updateUserOptions.Password = &newPassword
	updateUserOptions.Preferences = res[0].Preferences
	updateUserOptions.PreviousLoginTimeMsecs = res[0].PreviousLoginTimeMsecs
	updateUserOptions.PrimaryGroupName = res[0].PrimaryGroupName
	updateUserOptions.PrivilegeIds = res[0].PrivilegeIds
	updateUserOptions.Profiles = res[0].Profiles
	updateUserOptions.Restricted = res[0].Restricted
	updateUserOptions.Roles = res[0].Roles
	updateUserOptions.S3AccessKeyID = res[0].S3AccessKeyID
	updateUserOptions.S3AccountID = res[0].S3AccountID
	updateUserOptions.S3SecretKey = res[0].S3SecretKey
	updateUserOptions.SalesforceAccount = res[0].SalesforceAccount
	updateUserOptions.Sid = res[0].Sid
	updateUserOptions.SpogContext = res[0].SpogContext
	updateUserOptions.SubscriptionInfo = res[0].SubscriptionInfo
	updateUserOptions.TenantAccesses = res[0].TenantAccesses
	updateUserOptions.TenantID = res[0].TenantID
	updateUserOptions.Username = res[0].Username

	if _, ok := d.GetOk("session_name_cookie"); ok {
		headerMap := map[string]string{"Cookie": d.Get("session_name_cookie").(string)}
		updateUserOptions.SetHeaders(headerMap)
	}

	_, _, err = backupRecoveryConnectorClient.UpdateUserWithContext(context, updateUserOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateUserWithContext failed: %s", err.Error()), "ibm_backup_recovery_reset_password", "updateUser")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(ResourceIbmBackupRecoveryConnectorPasswordResetID(d))

	return resourceIbmBackupRecoveryConnectorAccessTokenRead(context, d, meta)
}

func ResourceIbmBackupRecoveryConnectorPasswordResetID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func ResourceIbmBackupRecoveryConnectorPasswordResetRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return nil
}

func ResourceIbmBackupRecoveryConnectorPasswordResetDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	d.SetId("")
	return nil
}
