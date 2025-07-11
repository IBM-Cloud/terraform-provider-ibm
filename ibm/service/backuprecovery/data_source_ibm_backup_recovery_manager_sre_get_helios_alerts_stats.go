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

func DataSourceIbmBackupRecoveryManagerSreGetHeliosAlertsStats() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryManagerSreGetHeliosAlertsStatsRead,

		Schema: map[string]*schema.Schema{
			"start_time_usecs": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Specifies the start time Unix time epoch in microseconds from which the active alerts stats are computed.",
			},
			"end_time_usecs": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Specifies the end time Unix time epoch in microseconds to which the active alerts stats are computed.",
			},
			"cluster_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the list of cluster IDs.",
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"region_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter by a list of region ids.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"exclude_stats_by_cluster": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specifies if stats of active alerts per cluster needs to be excluded. If set to false (default value), stats of active alerts per cluster is included in the response. If set to true, only aggregated stats summary will be present in the response.",
			},
			"alert_source": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies a list of alert origination source. If not specified, all alerts from all the sources are considered in the response.",
			},
			"tenant_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies a list of tenants.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"aggregated_alerts_stats": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the active alert statistics details.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"num_critical_alerts": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the count of active critical Alerts excluding alerts that belong to other bucket.",
						},
						"num_critical_alerts_categories": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the count of active critical alerts categories.",
						},
						"num_data_service_alerts": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the count of active service Alerts.",
						},
						"num_data_service_critical_alerts": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the count of active service critical Alerts.",
						},
						"num_data_service_info_alerts": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the count of active service info Alerts.",
						},
						"num_data_service_warning_alerts": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the count of active service warning Alerts.",
						},
						"num_hardware_alerts": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the count of active hardware Alerts.",
						},
						"num_hardware_critical_alerts": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the count of active hardware critical Alerts.",
						},
						"num_hardware_info_alerts": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the count of active hardware info Alerts.",
						},
						"num_hardware_warning_alerts": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the count of active hardware warning Alerts.",
						},
						"num_info_alerts": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the count of active info Alerts excluding alerts that belong to other bucket.",
						},
						"num_info_alerts_categories": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the count of active info alerts categories.",
						},
						"num_maintenance_alerts": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the count of active Alerts of maintenance bucket.",
						},
						"num_maintenance_critical_alerts": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the count of active other critical Alerts.",
						},
						"num_maintenance_info_alerts": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the count of active other info Alerts.",
						},
						"num_maintenance_warning_alerts": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the count of active other warning Alerts.",
						},
						"num_software_alerts": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the count of active software Alerts.",
						},
						"num_software_critical_alerts": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the count of active software critical Alerts.",
						},
						"num_software_info_alerts": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the count of active software info Alerts.",
						},
						"num_software_warning_alerts": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the count of active software warning Alerts.",
						},
						"num_warning_alerts": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the count of active warning Alerts excluding alerts that belong to other bucket.",
						},
						"num_warning_alerts_categories": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the count of active warning alerts categories.",
						},
					},
				},
			},
			"aggregated_cluster_stats": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the cluster statistics based on active alerts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"num_clusters_with_critical_alerts": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the count of clusters with at least one critical alert.",
						},
						"num_clusters_with_warning_alerts": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the count of clusters with at least one warning category alert and no critical alerts.",
						},
						"num_healthy_clusters": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the count of clusters with no warning or critical alerts.",
						},
					},
				},
			},
			"stats_by_cluster": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the active Alerts stats by clusters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alerts_stats": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the active alert statistics details.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"num_critical_alerts": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the count of active critical Alerts excluding alerts that belong to other bucket.",
									},
									"num_critical_alerts_categories": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the count of active critical alerts categories.",
									},
									"num_data_service_alerts": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the count of active service Alerts.",
									},
									"num_data_service_critical_alerts": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the count of active service critical Alerts.",
									},
									"num_data_service_info_alerts": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the count of active service info Alerts.",
									},
									"num_data_service_warning_alerts": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the count of active service warning Alerts.",
									},
									"num_hardware_alerts": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the count of active hardware Alerts.",
									},
									"num_hardware_critical_alerts": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the count of active hardware critical Alerts.",
									},
									"num_hardware_info_alerts": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the count of active hardware info Alerts.",
									},
									"num_hardware_warning_alerts": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the count of active hardware warning Alerts.",
									},
									"num_info_alerts": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the count of active info Alerts excluding alerts that belong to other bucket.",
									},
									"num_info_alerts_categories": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the count of active info alerts categories.",
									},
									"num_maintenance_alerts": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the count of active Alerts of maintenance bucket.",
									},
									"num_maintenance_critical_alerts": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the count of active other critical Alerts.",
									},
									"num_maintenance_info_alerts": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the count of active other info Alerts.",
									},
									"num_maintenance_warning_alerts": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the count of active other warning Alerts.",
									},
									"num_software_alerts": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the count of active software Alerts.",
									},
									"num_software_critical_alerts": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the count of active software critical Alerts.",
									},
									"num_software_info_alerts": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the count of active software info Alerts.",
									},
									"num_software_warning_alerts": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the count of active software warning Alerts.",
									},
									"num_warning_alerts": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the count of active warning Alerts excluding alerts that belong to other bucket.",
									},
									"num_warning_alerts_categories": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the count of active warning alerts categories.",
									},
								},
							},
						},
						"cluster_id": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the Cluster Id.",
						},
						"region_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the region id of cluster.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmBackupRecoveryManagerSreGetHeliosAlertsStatsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	heliosSreApiClient, err := meta.(conns.ClientSession).BackupRecoveryManagerSreV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_sre_get_helios_alerts_stats", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getHeliosAlertsStatsOptions := &backuprecoveryv1.GetHeliosAlertsStatsOptions{}

	getHeliosAlertsStatsOptions.SetStartTimeUsecs(int64(d.Get("start_time_usecs").(int)))
	getHeliosAlertsStatsOptions.SetEndTimeUsecs(int64(d.Get("end_time_usecs").(int)))
	if _, ok := d.GetOk("cluster_ids"); ok {
		var clusterIds []int64
		for _, v := range d.Get("cluster_ids").([]interface{}) {
			clusterIdsItem := int64(v.(int))
			clusterIds = append(clusterIds, clusterIdsItem)
		}
		getHeliosAlertsStatsOptions.SetClusterIds(clusterIds)
	}
	if _, ok := d.GetOk("region_ids"); ok {
		var regionIds []string
		for _, v := range d.Get("region_ids").([]interface{}) {
			regionIdsItem := v.(string)
			regionIds = append(regionIds, regionIdsItem)
		}
		getHeliosAlertsStatsOptions.SetRegionIds(regionIds)
	}
	if _, ok := d.GetOk("exclude_stats_by_cluster"); ok {
		getHeliosAlertsStatsOptions.SetExcludeStatsByCluster(d.Get("exclude_stats_by_cluster").(bool))
	}
	if _, ok := d.GetOk("alert_source"); ok {
		getHeliosAlertsStatsOptions.SetAlertSource(d.Get("alert_source").(string))
	}
	if _, ok := d.GetOk("tenant_ids"); ok {
		var tenantIds []string
		for _, v := range d.Get("tenant_ids").([]interface{}) {
			tenantIdsItem := v.(string)
			tenantIds = append(tenantIds, tenantIdsItem)
		}
		getHeliosAlertsStatsOptions.SetTenantIds(tenantIds)
	}

	mcmActiveAlertsStats, _, err := heliosSreApiClient.GetHeliosAlertsStatsWithContext(context, getHeliosAlertsStatsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetHeliosAlertsStatsWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_manager_sre_get_helios_alerts_stats", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBackupRecoveryManagerSreGetHeliosAlertsStatsID(d))

	if !core.IsNil(mcmActiveAlertsStats.AggregatedAlertsStats) {
		aggregatedAlertsStats := []map[string]interface{}{}
		aggregatedAlertsStatsMap, err := DataSourceIbmBackupRecoveryManagerSreGetHeliosAlertsStatsActiveAlertsStatsToMap(mcmActiveAlertsStats.AggregatedAlertsStats)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_sre_get_helios_alerts_stats", "read", "aggregated_alerts_stats-to-map").GetDiag()
		}
		aggregatedAlertsStats = append(aggregatedAlertsStats, aggregatedAlertsStatsMap)
		if err = d.Set("aggregated_alerts_stats", aggregatedAlertsStats); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting aggregated_alerts_stats: %s", err), "(Data) ibm_backup_recovery_manager_sre_get_helios_alerts_stats", "read", "set-aggregated_alerts_stats").GetDiag()
		}
	}

	if !core.IsNil(mcmActiveAlertsStats.AggregatedClusterStats) {
		aggregatedClusterStats := []map[string]interface{}{}
		aggregatedClusterStatsMap, err := DataSourceIbmBackupRecoveryManagerSreGetHeliosAlertsStatsClusterAlertStatsToMap(mcmActiveAlertsStats.AggregatedClusterStats)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_sre_get_helios_alerts_stats", "read", "aggregated_cluster_stats-to-map").GetDiag()
		}
		aggregatedClusterStats = append(aggregatedClusterStats, aggregatedClusterStatsMap)
		if err = d.Set("aggregated_cluster_stats", aggregatedClusterStats); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting aggregated_cluster_stats: %s", err), "(Data) ibm_backup_recovery_manager_sre_get_helios_alerts_stats", "read", "set-aggregated_cluster_stats").GetDiag()
		}
	}

	if !core.IsNil(mcmActiveAlertsStats.StatsByCluster) {
		statsByCluster := []map[string]interface{}{}
		for _, statsByClusterItem := range mcmActiveAlertsStats.StatsByCluster {
			statsByClusterItemMap, err := DataSourceIbmBackupRecoveryManagerSreGetHeliosAlertsStatsMcmActiveAlertsStatsByClusterToMap(&statsByClusterItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_sre_get_helios_alerts_stats", "read", "stats_by_cluster-to-map").GetDiag()
			}
			statsByCluster = append(statsByCluster, statsByClusterItemMap)
		}
		if err = d.Set("stats_by_cluster", statsByCluster); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting stats_by_cluster: %s", err), "(Data) ibm_backup_recovery_manager_sre_get_helios_alerts_stats", "read", "set-stats_by_cluster").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmBackupRecoveryManagerSreGetHeliosAlertsStatsID returns a reasonable ID for the list.
func dataSourceIbmBackupRecoveryManagerSreGetHeliosAlertsStatsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBackupRecoveryManagerSreGetHeliosAlertsStatsActiveAlertsStatsToMap(model *backuprecoveryv1.ActiveAlertsStats) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.NumCriticalAlerts != nil {
		modelMap["num_critical_alerts"] = flex.IntValue(model.NumCriticalAlerts)
	}
	if model.NumCriticalAlertsCategories != nil {
		modelMap["num_critical_alerts_categories"] = flex.IntValue(model.NumCriticalAlertsCategories)
	}
	if model.NumDataServiceAlerts != nil {
		modelMap["num_data_service_alerts"] = flex.IntValue(model.NumDataServiceAlerts)
	}
	if model.NumDataServiceCriticalAlerts != nil {
		modelMap["num_data_service_critical_alerts"] = flex.IntValue(model.NumDataServiceCriticalAlerts)
	}
	if model.NumDataServiceInfoAlerts != nil {
		modelMap["num_data_service_info_alerts"] = flex.IntValue(model.NumDataServiceInfoAlerts)
	}
	if model.NumDataServiceWarningAlerts != nil {
		modelMap["num_data_service_warning_alerts"] = flex.IntValue(model.NumDataServiceWarningAlerts)
	}
	if model.NumHardwareAlerts != nil {
		modelMap["num_hardware_alerts"] = flex.IntValue(model.NumHardwareAlerts)
	}
	if model.NumHardwareCriticalAlerts != nil {
		modelMap["num_hardware_critical_alerts"] = flex.IntValue(model.NumHardwareCriticalAlerts)
	}
	if model.NumHardwareInfoAlerts != nil {
		modelMap["num_hardware_info_alerts"] = flex.IntValue(model.NumHardwareInfoAlerts)
	}
	if model.NumHardwareWarningAlerts != nil {
		modelMap["num_hardware_warning_alerts"] = flex.IntValue(model.NumHardwareWarningAlerts)
	}
	if model.NumInfoAlerts != nil {
		modelMap["num_info_alerts"] = flex.IntValue(model.NumInfoAlerts)
	}
	if model.NumInfoAlertsCategories != nil {
		modelMap["num_info_alerts_categories"] = flex.IntValue(model.NumInfoAlertsCategories)
	}
	if model.NumMaintenanceAlerts != nil {
		modelMap["num_maintenance_alerts"] = flex.IntValue(model.NumMaintenanceAlerts)
	}
	if model.NumMaintenanceCriticalAlerts != nil {
		modelMap["num_maintenance_critical_alerts"] = flex.IntValue(model.NumMaintenanceCriticalAlerts)
	}
	if model.NumMaintenanceInfoAlerts != nil {
		modelMap["num_maintenance_info_alerts"] = flex.IntValue(model.NumMaintenanceInfoAlerts)
	}
	if model.NumMaintenanceWarningAlerts != nil {
		modelMap["num_maintenance_warning_alerts"] = flex.IntValue(model.NumMaintenanceWarningAlerts)
	}
	if model.NumSoftwareAlerts != nil {
		modelMap["num_software_alerts"] = flex.IntValue(model.NumSoftwareAlerts)
	}
	if model.NumSoftwareCriticalAlerts != nil {
		modelMap["num_software_critical_alerts"] = flex.IntValue(model.NumSoftwareCriticalAlerts)
	}
	if model.NumSoftwareInfoAlerts != nil {
		modelMap["num_software_info_alerts"] = flex.IntValue(model.NumSoftwareInfoAlerts)
	}
	if model.NumSoftwareWarningAlerts != nil {
		modelMap["num_software_warning_alerts"] = flex.IntValue(model.NumSoftwareWarningAlerts)
	}
	if model.NumWarningAlerts != nil {
		modelMap["num_warning_alerts"] = flex.IntValue(model.NumWarningAlerts)
	}
	if model.NumWarningAlertsCategories != nil {
		modelMap["num_warning_alerts_categories"] = flex.IntValue(model.NumWarningAlertsCategories)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerSreGetHeliosAlertsStatsClusterAlertStatsToMap(model *backuprecoveryv1.ClusterAlertStats) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.NumClustersWithCriticalAlerts != nil {
		modelMap["num_clusters_with_critical_alerts"] = flex.IntValue(model.NumClustersWithCriticalAlerts)
	}
	if model.NumClustersWithWarningAlerts != nil {
		modelMap["num_clusters_with_warning_alerts"] = flex.IntValue(model.NumClustersWithWarningAlerts)
	}
	if model.NumHealthyClusters != nil {
		modelMap["num_healthy_clusters"] = flex.IntValue(model.NumHealthyClusters)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerSreGetHeliosAlertsStatsMcmActiveAlertsStatsByClusterToMap(model *backuprecoveryv1.McmActiveAlertsStatsByCluster) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AlertsStats != nil {
		alertsStatsMap, err := DataSourceIbmBackupRecoveryManagerSreGetHeliosAlertsStatsActiveAlertsStatsToMap(model.AlertsStats)
		if err != nil {
			return modelMap, err
		}
		modelMap["alerts_stats"] = []map[string]interface{}{alertsStatsMap}
	}
	if model.ClusterID != nil {
		modelMap["cluster_id"] = flex.IntValue(model.ClusterID)
	}
	if model.RegionID != nil {
		modelMap["region_id"] = *model.RegionID
	}
	return modelMap, nil
}
