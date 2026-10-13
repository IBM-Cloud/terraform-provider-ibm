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
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func DataSourceIbmBackupRecoveryConnectorAgentsUpgradeStatus() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryConnectorAgentsUpgradeStatusRead,

		Schema: map[string]*schema.Schema{
			"upgrade_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the ID of the upgrade run, as returned by POST /connector-agents/upgrade.",
			},
			"tenant_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filters the returned tenantStatuses to only these tenants. Tenants not part of this upgrade run are silently ignored. If omitted, every tenant in the run is returned.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"connection_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filters the connectionStatuses within each returned tenant to only these connections. Connections not part of this upgrade run are silently ignored. If omitted, every connection under each returned tenant is returned.",
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies the overall state of this upgrade run, derived from the states of every tenant in the run. 'NotStarted' indicates the run has not yet begun. 'InProgress' indicates at least one tenant is still upgrading. 'Failed' indicates no tenant is in progress and at least one has failed. 'Succeeded' indicates every tenant has succeeded. Unlike tenantStatuses, this field is not affected by the tenantIds or connectionIds filters.",
			},
			"tenant_statuses": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the upgrade status for every tenant included in this upgrade run, filtered by the tenantIds query parameter, if specified in the request.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tenant_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the ID of the tenant.",
						},
						"info": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the upgrade progress and outcome for a tenant.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"state": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the state of the upgrade run for this tenant. 'NotStarted' indicates the run has not yet begun. 'InProgress' indicates at least one connection is still upgrading. 'Failed' indicates every connection has reached a terminal state and at least one has failed. 'Succeeded' indicates every connection has succeeded.",
									},
									"target_version": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the version every connector agent for this tenant is being upgraded to. Absent if not specified on the originating upgrade request.",
									},
									"retry_forever": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies whether this tenant's upgrade retries indefinitely (true) or is capped at a fixed number of attempts (false), as set on the originating upgrade request. Absent if not specified on the originating upgrade request.",
									},
									"total_connections": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the total number of connections in scope for this tenant's upgrade run. Not affected by the connectionIds filter. Absent until the run is initialized.",
									},
									"succeeded_connections": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the number of connections for this tenant that upgrade successfully to the target version. Not affected by the connectionIds filter. Absent until the run is initialized.",
									},
									"failed_connections": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the number of connections for this tenant that exhaust their retry budget without upgrading successfully. Not affected by the connectionIds filter. Absent until the run is initialized.",
									},
									"created_timestamp_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the timestamp in UNIX time (microseconds) when this tenant's upgrade marker is created. Absent if not yet recorded.",
									},
									"finished_timestamp_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the timestamp in UNIX time (microseconds) when every connection for this tenant reached a terminal state. Absent while this tenant's upgrade is still in progress.",
									},
									"connection_statuses": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the upgrade status of every individual connection for this tenant included in this upgrade run, filtered by the connectionIds query parameter, if specified in the request. Absent if the run has no connections.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"connection_id": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the ID of the connection.",
												},
												"status": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the last fetched upgrade status of the connector agent on this connection. 'NotStarted' indicates the upgrade has not yet begun. 'InProgress' indicates the upgrade is underway. 'Succeeded' indicates the connection upgraded successfully. 'Failed' indicates the upgrade attempt failed.",
												},
												"attempt_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the number of upgrade attempts made so far for this connection in the current upgrade run. Absent if no attempt has been made.",
												},
												"last_attempt_timestamp_msecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the timestamp in UNIX time (milliseconds) of the most recent upgrade attempt for this connection. Absent if no attempt has been made.",
												},
												"last_status_fetched_timestamp_msecs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the timestamp in UNIX time (milliseconds) when this connection's upgrade status is last fetched from the connector agent. Absent if the status has not yet been fetched.",
												},
												"error_message": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the error message when the status is 'Failed'. Absent otherwise.",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmBackupRecoveryConnectorAgentsUpgradeStatusRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_connector_agents_upgrade_status", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getConnectorAgentsUpgradeStatusOptions := &backuprecoveryv1.GetConnectorAgentsUpgradeStatusOptions{}

	getConnectorAgentsUpgradeStatusOptions.SetUpgradeID(core.UUIDPtr(strfmt.UUID(d.Get("upgrade_id").(string))))
	if _, ok := d.GetOk("tenant_ids"); ok {
		var tenantIds []string
		for _, v := range d.Get("tenant_ids").([]interface{}) {
			tenantIdsItem := v.(string)
			tenantIds = append(tenantIds, tenantIdsItem)
		}
		getConnectorAgentsUpgradeStatusOptions.SetTenantIds(tenantIds)
	}
	if _, ok := d.GetOk("connection_ids"); ok {
		var connectionIds []int64
		for _, v := range d.Get("connection_ids").([]interface{}) {
			connectionIdsItem := int64(v.(int))
			connectionIds = append(connectionIds, connectionIdsItem)
		}
		getConnectorAgentsUpgradeStatusOptions.SetConnectionIds(connectionIds)
	}

	connectorAgentsUpgradeStatus, _, err := backupRecoveryClient.GetConnectorAgentsUpgradeStatusWithContext(context, getConnectorAgentsUpgradeStatusOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetConnectorAgentsUpgradeStatusWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_connector_agents_upgrade_status", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBackupRecoveryConnectorAgentsUpgradeStatusID(d))

	if err = d.Set("status", connectorAgentsUpgradeStatus.Status); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_backup_recovery_connector_agents_upgrade_status", "read", "set-status").GetDiag()
	}

	tenantStatuses := []map[string]interface{}{}
	for _, tenantStatusesItem := range connectorAgentsUpgradeStatus.TenantStatuses {
		tenantStatusesItemMap, err := DataSourceIbmBackupRecoveryConnectorAgentsUpgradeStatusConnectorAgentTenantUpgradeStatusToMap(&tenantStatusesItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_connector_agents_upgrade_status", "read", "tenant_statuses-to-map").GetDiag()
		}
		tenantStatuses = append(tenantStatuses, tenantStatusesItemMap)
	}
	if err = d.Set("tenant_statuses", tenantStatuses); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting tenant_statuses: %s", err), "(Data) ibm_backup_recovery_connector_agents_upgrade_status", "read", "set-tenant_statuses").GetDiag()
	}

	return nil
}

// dataSourceIbmBackupRecoveryConnectorAgentsUpgradeStatusID returns a reasonable ID for the list.
func dataSourceIbmBackupRecoveryConnectorAgentsUpgradeStatusID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBackupRecoveryConnectorAgentsUpgradeStatusConnectorAgentTenantUpgradeStatusToMap(model *backuprecoveryv1.ConnectorAgentTenantUpgradeStatus) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["tenant_id"] = *model.TenantID
	infoMap, err := DataSourceIbmBackupRecoveryConnectorAgentsUpgradeStatusConnectorAgentTenantUpgradeInfoToMap(model.Info)
	if err != nil {
		return modelMap, err
	}
	modelMap["info"] = []map[string]interface{}{infoMap}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryConnectorAgentsUpgradeStatusConnectorAgentTenantUpgradeInfoToMap(model *backuprecoveryv1.ConnectorAgentTenantUpgradeInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["state"] = *model.State
	if model.TargetVersion != nil {
		modelMap["target_version"] = *model.TargetVersion
	}
	if model.RetryForever != nil {
		modelMap["retry_forever"] = *model.RetryForever
	}
	if model.TotalConnections != nil {
		modelMap["total_connections"] = flex.IntValue(model.TotalConnections)
	}
	if model.SucceededConnections != nil {
		modelMap["succeeded_connections"] = flex.IntValue(model.SucceededConnections)
	}
	if model.FailedConnections != nil {
		modelMap["failed_connections"] = flex.IntValue(model.FailedConnections)
	}
	if model.CreatedTimestampUsecs != nil {
		modelMap["created_timestamp_usecs"] = flex.IntValue(model.CreatedTimestampUsecs)
	}
	if model.FinishedTimestampUsecs != nil {
		modelMap["finished_timestamp_usecs"] = flex.IntValue(model.FinishedTimestampUsecs)
	}
	if model.ConnectionStatuses != nil {
		connectionStatuses := []map[string]interface{}{}
		for _, connectionStatusesItem := range model.ConnectionStatuses {
			connectionStatusesItemMap, err := DataSourceIbmBackupRecoveryConnectorAgentsUpgradeStatusConnectorAgentUpgradeConnectionStatusToMap(&connectionStatusesItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			connectionStatuses = append(connectionStatuses, connectionStatusesItemMap)
		}
		modelMap["connection_statuses"] = connectionStatuses
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryConnectorAgentsUpgradeStatusConnectorAgentUpgradeConnectionStatusToMap(model *backuprecoveryv1.ConnectorAgentUpgradeConnectionStatus) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["connection_id"] = flex.IntValue(model.ConnectionID)
	modelMap["status"] = *model.Status
	if model.AttemptCount != nil {
		modelMap["attempt_count"] = flex.IntValue(model.AttemptCount)
	}
	if model.LastAttemptTimestampMsecs != nil {
		modelMap["last_attempt_timestamp_msecs"] = flex.IntValue(model.LastAttemptTimestampMsecs)
	}
	if model.LastStatusFetchedTimestampMsecs != nil {
		modelMap["last_status_fetched_timestamp_msecs"] = flex.IntValue(model.LastStatusFetchedTimestampMsecs)
	}
	if model.ErrorMessage != nil {
		modelMap["error_message"] = *model.ErrorMessage
	}
	return modelMap, nil
}
