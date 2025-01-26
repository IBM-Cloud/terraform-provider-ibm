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

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func ResourceIbmBackupRecoveryConnectorRegistration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmBackupRecoveryConnectorRegistrationCreate,
		ReadContext:   resourceIbmBackupRecoveryConnectorRegistrationRead,
		DeleteContext: resourceIbmBackupRecoveryConnectorRegistrationDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"connector_id": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "The connector's ID to be used for registration. Two connectors belonging to the same tenant are guaranteed to have different IDs.",
			},
			"access_token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Token required to authenticate to the connector",
			},
			"registration_token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The registration token.",
			},
		},
	}
}

func resourceIbmBackupRecoveryConnectorRegistrationCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryConnectorClient, err := meta.(conns.ClientSession).BackupRecoveryV1Connector()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_registration", "read", "initialize-client")
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
	backupRecoveryConnectorClient.SetAuthenticator(&auth)

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

func resourceIbmBackupRecoveryConnectorRegistrationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return nil
}

func resourceIbmBackupRecoveryConnectorRegistrationDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	d.SetId("")
	return nil
}
