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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func DataSourceIbmBackupRecoveryConnectorAgentConfig() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryConnectorAgentConfigRead,

		Schema: map[string]*schema.Schema{
			"x_ibm_tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.",
			},
			"registration_token": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Token that is used for authenticating the connector agent with the DataProtect cluster.",
			},
		},
	}
}

func dataSourceIbmBackupRecoveryConnectorAgentConfigRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_connector_agent_config", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	endpointType := d.Get("endpoint_type").(string)
	instanceId, region, serviceName := getInstanceIdAndRegion(d)
	if instanceId != "" && region != "" {
		bmxsession, err := meta.(conns.ClientSession).BluemixSession()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("unable to get clientSession"), "ibm_backup_recovery_connector_agent_config", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		backupRecoveryClient = getClientWithInstanceEndpoint(backupRecoveryClient, bmxsession, instanceId, region, endpointType, serviceName)
	}

	getConnectorAgentConfigOptions := &backuprecoveryv1.GetConnectorAgentConfigOptions{}
	getConnectorAgentConfigOptions.SetXIBMTenantID(d.Get("x_ibm_tenant_id").(string))

	connectorAgentConfig, _, err := backupRecoveryClient.GetConnectorAgentConfigWithContext(context, getConnectorAgentConfigOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetConnectorAgentConfigWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_connector_agent_config", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBackupRecoveryConnectorAgentConfigID(d))

	if !core.IsNil(connectorAgentConfig.RegistrationToken) {
		if err = d.Set("registration_token", connectorAgentConfig.RegistrationToken); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting registration_token: %s", err), "(Data) ibm_backup_recovery_connector_agent_config", "read", "set-registration_token").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmBackupRecoveryConnectorAgentConfigID returns a reasonable ID for the list.
func dataSourceIbmBackupRecoveryConnectorAgentConfigID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
