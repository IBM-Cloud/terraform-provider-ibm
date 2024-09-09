// Copyright IBM Corp. 2024 All Rights Reserved.
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
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func DataSourceIbmBaasConnectorLogs() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBaasConnectorLogsRead,

		Schema: map[string]*schema.Schema{
			"x_ibm_tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.",
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

func dataSourceIbmBaasConnectorLogsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_baas_connector_logs", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getDataSourceConnectorLogsOptions := &backuprecoveryv1.GetDataSourceConnectorLogsOptions{}

	getDataSourceConnectorLogsOptions.SetXIBMTenantID(d.Get("x_ibm_tenant_id").(string))

	dataSourceConnectorLogs, _, err := backupRecoveryClient.GetDataSourceConnectorLogsWithContext(context, getDataSourceConnectorLogsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDataSourceConnectorLogsWithContext failed: %s", err.Error()), "(Data) ibm_baas_connector_logs", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBaasConnectorLogsID(d))

	if !core.IsNil(dataSourceConnectorLogs.ConnectorLogs) {
		connectorLogs := []map[string]interface{}{}
		for _, connectorLogsItem := range dataSourceConnectorLogs.ConnectorLogs {
			connectorLogsItemMap, err := DataSourceIbmBaasConnectorLogsDataSourceConnectorLogToMap(&connectorLogsItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_baas_connector_logs", "read", "connector_logs-to-map").GetDiag()
			}
			connectorLogs = append(connectorLogs, connectorLogsItemMap)
		}
		if err = d.Set("connector_logs", connectorLogs); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting connector_logs: %s", err), "(Data) ibm_baas_connector_logs", "read", "set-connector_logs").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmBaasConnectorLogsID returns a reasonable ID for the list.
func dataSourceIbmBaasConnectorLogsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBaasConnectorLogsDataSourceConnectorLogToMap(model *backuprecoveryv1.DataSourceConnectorLog) (map[string]interface{}, error) {
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
