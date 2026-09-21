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

func DataSourceIbmBackupRecoveryConnectorAgents() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryConnectorAgentsRead,

		Schema: map[string]*schema.Schema{
			"x_ibm_tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.",
			},
			"tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the ID of the tenant for which the connector agents are to be fetched.",
			},
			"connection_names": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the connection names whose connector agents are to be fetched.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"connection_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the connection IDs whose connector agents are to be fetched.",
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"connector_agents": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connector_agent_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the unique ID of the connector agent.",
						},
						"connector_agent_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the name of the connector agent.",
						},
						"connection_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the ID of the connection to which this connector agent belongs.",
						},
						"connection_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the name of the connection to which this connector agent belongs.",
						},
						"software_version": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the connector agent's software version.",
						},
						"connectivity_status": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies connector agent connectivity statusinformation like current connectivity status to cluster,when it last connected to the cluster successfully and fromwhen it has been continuously connected to the cluster without anyinterruptions.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"is_connected": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies whether the connector agent is currently connected to the cluster.",
									},
									"message": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies error message when the connector agent is unable to connect to the cluster.",
									},
									"connected_since_timestamp_secs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "This denotes the timestamp in UNIX seconds since when this connector agent has been connected to its cluster without any interruptions. This property will not be present if this connector agent is not currently connected to its cluster.",
									},
									"last_known_health_ok_timestamp_secs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the most recent known timestamp in UNIX seconds at which this connector agent passed the health checks. This property can be present even if this connector agent is not currently connected to its cluster.",
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

func dataSourceIbmBackupRecoveryConnectorAgentsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_connector_agents", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	endpointType := d.Get("endpoint_type").(string)
	instanceId, region, serviceName := getInstanceIdAndRegion(d)
	if instanceId != "" && region != "" {
		bmxsession, err := meta.(conns.ClientSession).BluemixSession()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("unable to get clientSession"), "ibm_backup_recovery_connector_agents", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		backupRecoveryClient = getClientWithInstanceEndpoint(backupRecoveryClient, bmxsession, instanceId, region, endpointType, serviceName)
	}

	listConnectorAgentsOptions := &backuprecoveryv1.ListConnectorAgentsOptions{}
	listConnectorAgentsOptions.SetXIBMTenantID(d.Get("x_ibm_tenant_id").(string))

	listConnectorAgentsOptions.SetTenantID(d.Get("tenant_id").(string))
	if _, ok := d.GetOk("connection_names"); ok {
		var connectionNames []string
		for _, v := range d.Get("connection_names").([]interface{}) {
			connectionNamesItem := v.(string)
			connectionNames = append(connectionNames, connectionNamesItem)
		}
		listConnectorAgentsOptions.SetConnectionNames(connectionNames)
	}
	if _, ok := d.GetOk("connection_ids"); ok {
		var connectionIds []int64
		for _, v := range d.Get("connection_ids").([]interface{}) {
			connectionIdsItem := int64(v.(int))
			connectionIds = append(connectionIds, connectionIdsItem)
		}
		listConnectorAgentsOptions.SetConnectionIds(connectionIds)
	}

	connectorAgentsList, _, err := backupRecoveryClient.ListConnectorAgentsWithContext(context, listConnectorAgentsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListConnectorAgentsWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_connector_agents", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBackupRecoveryConnectorAgentsID(d))

	if !core.IsNil(connectorAgentsList.ConnectorAgents) {
		connectorAgents := []map[string]interface{}{}
		for _, connectorAgentsItem := range connectorAgentsList.ConnectorAgents {
			connectorAgentsItemMap, err := DataSourceIbmBackupRecoveryConnectorAgentsConnectorAgentToMap(&connectorAgentsItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_connector_agents", "read", "connector_agents-to-map").GetDiag()
			}
			connectorAgents = append(connectorAgents, connectorAgentsItemMap)
		}
		if err = d.Set("connector_agents", connectorAgents); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting connector_agents: %s", err), "(Data) ibm_backup_recovery_connector_agents", "read", "set-connector_agents").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmBackupRecoveryConnectorAgentsID returns a reasonable ID for the list.
func dataSourceIbmBackupRecoveryConnectorAgentsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBackupRecoveryConnectorAgentsConnectorAgentToMap(model *backuprecoveryv1.ConnectorAgent) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["connector_agent_id"] = *model.ConnectorAgentID
	modelMap["connector_agent_name"] = *model.ConnectorAgentName
	modelMap["connection_id"] = *model.ConnectionID
	if model.ConnectionName != nil {
		modelMap["connection_name"] = *model.ConnectionName
	}
	if model.SoftwareVersion != nil {
		modelMap["software_version"] = *model.SoftwareVersion
	}
	if model.ConnectivityStatus != nil {
		connectivityStatusMap, err := DataSourceIbmBackupRecoveryConnectorAgentsConnectorAgentConnectivityStatusToMap(model.ConnectivityStatus)
		if err != nil {
			return modelMap, err
		}
		modelMap["connectivity_status"] = []map[string]interface{}{connectivityStatusMap}
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryConnectorAgentsConnectorAgentConnectivityStatusToMap(model *backuprecoveryv1.ConnectorAgentConnectivityStatus) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["is_connected"] = *model.IsConnected
	if model.Message != nil {
		modelMap["message"] = *model.Message
	}
	if model.ConnectedSinceTimestampSecs != nil {
		modelMap["connected_since_timestamp_secs"] = flex.IntValue(model.ConnectedSinceTimestampSecs)
	}
	if model.LastKnownHealthOkTimestampSecs != nil {
		modelMap["last_known_health_ok_timestamp_secs"] = flex.IntValue(model.LastKnownHealthOkTimestampSecs)
	}
	return modelMap, nil
}
