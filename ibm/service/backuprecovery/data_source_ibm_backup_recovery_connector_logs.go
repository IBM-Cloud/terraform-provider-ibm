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

func DataSourceIbmBackupRecoveryConnectorLogs() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryConnectorLogsRead,

		Schema: map[string]*schema.Schema{
			"access_token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Token required to authenticate to the connector. Token can be obtained using ibm_backup_recovery_connector_access_token resource",
			},
			"connector_logs": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the data-source connector logs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"message": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the message of this event.",
						},
						"timestamp_msecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the time stamp in milliseconds of the event.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the severity of the event.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmBackupRecoveryConnectorLogsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryConnectorClient, err := meta.(conns.ClientSession).BackupRecoveryV1Connector()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_connector_logs", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if backupRecoveryConnectorClient.GetConnectorURL() == "" {
		tfErr := flex.DiscriminatedTerraformErrorf(nil, "No connector URL specified. Please set the `IBMCLOUD_BACKUP_RECOVERY_CONNECTOR_ENDPOINT` environment variable or specify the endpoint in endpoints.json file.", "ibm_backup_recovery_connector_logs", "read", "initialize-client")
		return tfErr.GetDiag()
	}

	accessToken := d.Get("access_token").(string)
	var auth core.Authenticator
	auth = &core.BearerTokenAuthenticator{BearerToken: accessToken}
	backupRecoveryConnectorClient.Service.Options.Authenticator = auth

	getDataSourceConnectorLogsOptions := &backuprecoveryv1.GetDataSourceConnectorLogsOptions{}

	dataSourceConnectorLogs, _, err := backupRecoveryConnectorClient.GetDataSourceConnectorLogsWithContext(context, getDataSourceConnectorLogsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDataSourceConnectorLogsWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_connector_logs", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBackupRecoveryConnectorLogsID(d))

	if !core.IsNil(dataSourceConnectorLogs) {
		if !core.IsNil(dataSourceConnectorLogs.ConnectorLogs) {
			connectorLogs := []map[string]interface{}{}
			for _, connectorLogsItem := range dataSourceConnectorLogs.ConnectorLogs {
				connectorLogsItemMap, err := DataSourceIbmBackupRecoveryConnectorLogsDataSourceConnectorLogToMap(&connectorLogsItem) // #nosec G601
				if err != nil {
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_connector_logs", "read", "connector_logs-to-map").GetDiag()
				}
				connectorLogs = append(connectorLogs, connectorLogsItemMap)
			}
			if err = d.Set("connector_logs", connectorLogs); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting connector_logs: %s", err), "(Data) ibm_backup_recovery_connector_logs", "read", "set-connector_logs").GetDiag()
			}
		}
	}

	return nil
}

// dataSourceIbmBackupRecoveryConnectorLogsID returns a reasonable ID for the list.
func dataSourceIbmBackupRecoveryConnectorLogsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBackupRecoveryConnectorLogsDataSourceConnectorLogToMap(model *backuprecoveryv1.DataSourceConnectorLog) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Message != nil {
		modelMap["message"] = *model.Message
	}
	if model.TimestampMsecs != nil {
		modelMap["timestamp_msecs"] = flex.IntValue(model.TimestampMsecs)
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	return modelMap, nil
}
