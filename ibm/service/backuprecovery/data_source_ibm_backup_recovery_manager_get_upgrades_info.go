// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.96.1-5136e54a-20241108-203028
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

func DataSourceIbmBackupRecoveryManagerGetUpgradesInfo() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryManagerGetUpgradesInfoRead,

		Schema: map[string]*schema.Schema{
			"cluster_identifiers": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Fetch upgrade progress details for a list of cluster identifiers in format clusterId:clusterIncarnationId.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"upgrades_info": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_id": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies cluster's id.",
						},
						"cluster_incarnation_id": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies cluster's incarnation id.",
						},
						"patch_software_version": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Patch software version against which these logs are generated. This is specified for Patch type only.",
						},
						"software_version": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Upgrade software version against which these logs are generated.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the type of upgrade on a cluster.",
						},
						"upgrade_logs": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Upgrade logs per node.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"logs": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Upgrade logs for the node.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"log": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "One log statement of the complete logs.",
												},
												"time_stamp": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Time at which this log got generated.",
												},
											},
										},
									},
									"node_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Id of the node.",
									},
								},
							},
						},
						"upgrade_percent_complete": &schema.Schema{
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Upgrade percentage complete so far.",
						},
						"upgrade_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Upgrade status.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmBackupRecoveryManagerGetUpgradesInfoRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	managementApiClient, err := meta.(conns.ClientSession).BackupRecoveryManagerV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_upgrades_info", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	bmxsession, err := meta.(conns.ClientSession).BluemixSession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("unable to get clientSession"), "ibm_backup_recovery_manager_get_upgrades_info", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	endpointType := d.Get("endpoint_type").(string)
	instanceId, region := getInstanceIdAndRegion(d)
	managementApiClient, err = setManagerClientAuth(managementApiClient, bmxsession, region, endpointType)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("unable to set authenticator for clientSession: %s", err), "ibm_backup_recovery_manager_get_upgrades_info", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if instanceId != "" {
		managementApiClient = getManagerClientWithInstanceEndpoint(managementApiClient, bmxsession, instanceId, region, endpointType)
	}

	clustersUpgradesInfoOptions := &backuprecoveryv1.ClustersUpgradesInfoOptions{}

	if _, ok := d.GetOk("cluster_identifiers"); ok {
		var clusterIdentifiers []string
		for _, v := range d.Get("cluster_identifiers").([]interface{}) {
			clusterIdentifiersItem := v.(string)
			clusterIdentifiers = append(clusterIdentifiers, clusterIdentifiersItem)
		}
		clustersUpgradesInfoOptions.SetClusterIdentifiers(clusterIdentifiers)
	}

	upgradesInfo, _, err := managementApiClient.ClustersUpgradesInfoWithContext(context, clustersUpgradesInfoOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ClustersUpgradesInfoWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_manager_get_upgrades_info", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBackupRecoveryManagerGetUpgradesInfoID(d))

	if !core.IsNil(upgradesInfo) {
		upgradesInfoResult := []map[string]interface{}{}
		for _, upgradesInfoItem := range upgradesInfo {
			upgradesInfoItemMap, err := DataSourceIbmBackupRecoveryManagerGetUpgradesInfoUpgradeInfoToMap(&upgradesInfoItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_upgrades_info", "read", "upgrades_info-to-map").GetDiag()
			}
			upgradesInfoResult = append(upgradesInfoResult, upgradesInfoItemMap)
		}
		if err = d.Set("upgrades_info", upgradesInfoResult); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting upgrades_info: %s", err), "(Data) ibm_backup_recovery_manager_get_upgrades_info", "read", "set-upgrades_info").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmBackupRecoveryManagerGetUpgradesInfoID returns a reasonable ID for the list.
func dataSourceIbmBackupRecoveryManagerGetUpgradesInfoID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBackupRecoveryManagerGetUpgradesInfoUpgradeInfoToMap(model *backuprecoveryv1.UpgradeInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ClusterID != nil {
		modelMap["cluster_id"] = flex.IntValue(model.ClusterID)
	}
	if model.ClusterIncarnationID != nil {
		modelMap["cluster_incarnation_id"] = flex.IntValue(model.ClusterIncarnationID)
	}
	if model.PatchSoftwareVersion != nil {
		modelMap["patch_software_version"] = *model.PatchSoftwareVersion
	}
	if model.SoftwareVersion != nil {
		modelMap["software_version"] = *model.SoftwareVersion
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.UpgradeLogs != nil {
		upgradeLogs := []map[string]interface{}{}
		for _, upgradeLogsItem := range model.UpgradeLogs {
			upgradeLogsItemMap, err := DataSourceIbmBackupRecoveryManagerGetUpgradesInfoNodeUpgradeLogToMap(&upgradeLogsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			upgradeLogs = append(upgradeLogs, upgradeLogsItemMap)
		}
		modelMap["upgrade_logs"] = upgradeLogs
	}
	if model.UpgradePercentComplete != nil {
		modelMap["upgrade_percent_complete"] = *model.UpgradePercentComplete
	}
	if model.UpgradeStatus != nil {
		modelMap["upgrade_status"] = *model.UpgradeStatus
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetUpgradesInfoNodeUpgradeLogToMap(model *backuprecoveryv1.NodeUpgradeLog) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Logs != nil {
		logs := []map[string]interface{}{}
		for _, logsItem := range model.Logs {
			logsItemMap, err := DataSourceIbmBackupRecoveryManagerGetUpgradesInfoUpgradeLogToMap(&logsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			logs = append(logs, logsItemMap)
		}
		modelMap["logs"] = logs
	}
	if model.NodeID != nil {
		modelMap["node_id"] = *model.NodeID
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetUpgradesInfoUpgradeLogToMap(model *backuprecoveryv1.UpgradeLog) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Log != nil {
		modelMap["log"] = *model.Log
	}
	if model.TimeStamp != nil {
		modelMap["time_stamp"] = flex.IntValue(model.TimeStamp)
	}
	return modelMap, nil
}
