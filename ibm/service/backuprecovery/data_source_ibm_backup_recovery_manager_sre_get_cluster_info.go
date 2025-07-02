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

func DataSourceIbmBackupRecoveryManagerSreGetClusterInfo() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryManagerSreGetClusterInfoRead,

		Schema: map[string]*schema.Schema{
			"cohesity_clusters": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the array of clusters upgrade details.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_support_for_pkg_downloads": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If cluster can support authHeader for upgrade or not.",
						},
						"available_versions": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the release versions the cluster can upgrade to.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"notes": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies release's notes.",
									},
									"patch_details": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the details of the available patch release.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"notes": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies patch release's notes.",
												},
												"release_type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Release's type.",
												},
												"version": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies patch release's version.",
												},
											},
										},
									},
									"release_stage": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the stage of a release.",
									},
									"release_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Release's type e.g, LTS, Feature, Patch, MCM.",
									},
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the type of package or release.",
									},
									"version": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies release's version.",
									},
								},
							},
						},
						"cluster_id": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies cluster id.",
						},
						"cluster_incarnation_id": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies cluster incarnation id.",
						},
						"cluster_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies cluster's name.",
						},
						"current_patch_version": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies current patch version of the cluster.",
						},
						"current_version": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies if the cluster is connected to helios.",
						},
						"health": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the health of the cluster.",
						},
						"is_connected_to_helios": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Specifies if the cluster is connected to helios.",
						},
						"location": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the location of the cluster.",
						},
						"multi_tenancy_enabled": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Specifies if multi tenancy is enabled in the cluster.",
						},
						"node_ips": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies an array of node ips for the cluster.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"number_of_nodes": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the number of nodes in the cluster.",
						},
						"patch_target_upgrade_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the patch package URL for the cluster. This is populated for patch update only.",
						},
						"patch_target_version": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies target version to which clusters are upgrading. This is populated for patch update only.",
						},
						"provider_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the type of the cluster provider.",
						},
						"scheduled_timestamp": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Time at which an upgrade is scheduled.",
						},
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the upgrade status of the cluster.",
						},
						"target_upgrade_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the upgrade URL for the cluster. This is populated for upgrade only.",
						},
						"target_version": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies target version to which clusters are to be upgraded. This is populated for upgrade only.",
						},
						"total_capacity": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies how total memory capacity of the cluster.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the type of the cluster.",
						},
						"update_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the type of upgrade performed on a cluster. This is to be used with status field to know the status of the upgrade action performed on cluster.",
						},
						"used_capacity": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies how much of the cluster capacity is consumed.",
						},
					},
				},
			},
			"sp_clusters": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the array of clusters claimed from IBM Storage Protect environment.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_id": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies cluster id.",
						},
						"cluster_incarnation_id": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies cluster incarnation id.",
						},
						"cluster_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies cluster's name.",
						},
						"current_version": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the currently running version on cluster.",
						},
						"health": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the health of the cluster.",
						},
						"is_connected_to_helios": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Specifies if the cluster is connected to helios.",
						},
						"node_ips": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies an array of node ips for the cluster.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"number_of_nodes": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the number of nodes in the cluster.",
						},
						"provider_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the type of the cluster provider.",
						},
						"total_capacity": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies total capacity of the cluster in bytes.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the type of the SP cluster.",
						},
						"used_capacity": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies how much of the cluster capacity is consumed in bytes.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmBackupRecoveryManagerSreGetClusterInfoRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	heliosSreApiClient, err := meta.(conns.ClientSession).BackupRecoveryManagerSreV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_sre_get_cluster_info", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getClustersInfoOptions := &backuprecoveryv1.GetClustersInfoOptions{}

	clusterDetails, _, err := heliosSreApiClient.GetClustersInfoWithContext(context, getClustersInfoOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetClustersInfoWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_manager_sre_get_cluster_info", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBackupRecoveryManagerSreGetClusterInfoID(d))

	if !core.IsNil(clusterDetails.CohesityClusters) {
		cohesityClusters := []map[string]interface{}{}
		for _, cohesityClustersItem := range clusterDetails.CohesityClusters {
			cohesityClustersItemMap, err := DataSourceIbmBackupRecoveryManagerSreGetClusterInfoClusterInfoToMap(&cohesityClustersItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_sre_get_cluster_info", "read", "cohesity_clusters-to-map").GetDiag()
			}
			cohesityClusters = append(cohesityClusters, cohesityClustersItemMap)
		}
		if err = d.Set("cohesity_clusters", cohesityClusters); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting cohesity_clusters: %s", err), "(Data) ibm_backup_recovery_manager_sre_get_cluster_info", "read", "set-cohesity_clusters").GetDiag()
		}
	}

	if !core.IsNil(clusterDetails.SpClusters) {
		spClusters := []map[string]interface{}{}
		for _, spClustersItem := range clusterDetails.SpClusters {
			spClustersItemMap, err := DataSourceIbmBackupRecoveryManagerSreGetClusterInfoSPClusterInfoToMap(&spClustersItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_sre_get_cluster_info", "read", "sp_clusters-to-map").GetDiag()
			}
			spClusters = append(spClusters, spClustersItemMap)
		}
		if err = d.Set("sp_clusters", spClusters); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting sp_clusters: %s", err), "(Data) ibm_backup_recovery_manager_sre_get_cluster_info", "read", "set-sp_clusters").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmBackupRecoveryManagerSreGetClusterInfoID returns a reasonable ID for the list.
func dataSourceIbmBackupRecoveryManagerSreGetClusterInfoID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBackupRecoveryManagerSreGetClusterInfoClusterInfoToMap(model *backuprecoveryv1.ClusterInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AuthSupportForPkgDownloads != nil {
		modelMap["auth_support_for_pkg_downloads"] = *model.AuthSupportForPkgDownloads
	}
	if model.AvailableVersions != nil {
		availableVersions := []map[string]interface{}{}
		for _, availableVersionsItem := range model.AvailableVersions {
			availableVersionsItemMap, err := DataSourceIbmBackupRecoveryManagerSreGetClusterInfoAvailableReleaseVersionToMap(&availableVersionsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			availableVersions = append(availableVersions, availableVersionsItemMap)
		}
		modelMap["available_versions"] = availableVersions
	}
	if model.ClusterID != nil {
		modelMap["cluster_id"] = flex.IntValue(model.ClusterID)
	}
	if model.ClusterIncarnationID != nil {
		modelMap["cluster_incarnation_id"] = flex.IntValue(model.ClusterIncarnationID)
	}
	if model.ClusterName != nil {
		modelMap["cluster_name"] = *model.ClusterName
	}
	if model.CurrentPatchVersion != nil {
		modelMap["current_patch_version"] = *model.CurrentPatchVersion
	}
	if model.CurrentVersion != nil {
		modelMap["current_version"] = *model.CurrentVersion
	}
	if model.Health != nil {
		modelMap["health"] = *model.Health
	}
	if model.IsConnectedToHelios != nil {
		modelMap["is_connected_to_helios"] = *model.IsConnectedToHelios
	}
	if model.Location != nil {
		modelMap["location"] = *model.Location
	}
	if model.MultiTenancyEnabled != nil {
		modelMap["multi_tenancy_enabled"] = *model.MultiTenancyEnabled
	}
	if model.NodeIps != nil {
		modelMap["node_ips"] = model.NodeIps
	}
	if model.NumberOfNodes != nil {
		modelMap["number_of_nodes"] = flex.IntValue(model.NumberOfNodes)
	}
	if model.PatchTargetUpgradeURL != nil {
		modelMap["patch_target_upgrade_url"] = *model.PatchTargetUpgradeURL
	}
	if model.PatchTargetVersion != nil {
		modelMap["patch_target_version"] = *model.PatchTargetVersion
	}
	if model.ProviderType != nil {
		modelMap["provider_type"] = *model.ProviderType
	}
	if model.ScheduledTimestamp != nil {
		modelMap["scheduled_timestamp"] = flex.IntValue(model.ScheduledTimestamp)
	}
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	if model.TargetUpgradeURL != nil {
		modelMap["target_upgrade_url"] = *model.TargetUpgradeURL
	}
	if model.TargetVersion != nil {
		modelMap["target_version"] = *model.TargetVersion
	}
	if model.TotalCapacity != nil {
		modelMap["total_capacity"] = flex.IntValue(model.TotalCapacity)
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.UpdateType != nil {
		modelMap["update_type"] = *model.UpdateType
	}
	if model.UsedCapacity != nil {
		modelMap["used_capacity"] = flex.IntValue(model.UsedCapacity)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerSreGetClusterInfoAvailableReleaseVersionToMap(model *backuprecoveryv1.AvailableReleaseVersion) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Notes != nil {
		modelMap["notes"] = *model.Notes
	}
	if model.PatchDetails != nil {
		patchDetailsMap, err := DataSourceIbmBackupRecoveryManagerSreGetClusterInfoAvailablePatchReleaseToMap(model.PatchDetails)
		if err != nil {
			return modelMap, err
		}
		modelMap["patch_details"] = []map[string]interface{}{patchDetailsMap}
	}
	if model.ReleaseStage != nil {
		modelMap["release_stage"] = *model.ReleaseStage
	}
	if model.ReleaseType != nil {
		modelMap["release_type"] = *model.ReleaseType
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.Version != nil {
		modelMap["version"] = *model.Version
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerSreGetClusterInfoAvailablePatchReleaseToMap(model *backuprecoveryv1.AvailablePatchRelease) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Notes != nil {
		modelMap["notes"] = *model.Notes
	}
	if model.ReleaseType != nil {
		modelMap["release_type"] = *model.ReleaseType
	}
	if model.Version != nil {
		modelMap["version"] = *model.Version
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerSreGetClusterInfoSPClusterInfoToMap(model *backuprecoveryv1.SPClusterInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ClusterID != nil {
		modelMap["cluster_id"] = flex.IntValue(model.ClusterID)
	}
	if model.ClusterIncarnationID != nil {
		modelMap["cluster_incarnation_id"] = flex.IntValue(model.ClusterIncarnationID)
	}
	if model.ClusterName != nil {
		modelMap["cluster_name"] = *model.ClusterName
	}
	if model.CurrentVersion != nil {
		modelMap["current_version"] = *model.CurrentVersion
	}
	if model.Health != nil {
		modelMap["health"] = *model.Health
	}
	if model.IsConnectedToHelios != nil {
		modelMap["is_connected_to_helios"] = *model.IsConnectedToHelios
	}
	if model.NodeIps != nil {
		modelMap["node_ips"] = model.NodeIps
	}
	if model.NumberOfNodes != nil {
		modelMap["number_of_nodes"] = flex.IntValue(model.NumberOfNodes)
	}
	if model.ProviderType != nil {
		modelMap["provider_type"] = *model.ProviderType
	}
	if model.TotalCapacity != nil {
		modelMap["total_capacity"] = flex.IntValue(model.TotalCapacity)
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.UsedCapacity != nil {
		modelMap["used_capacity"] = flex.IntValue(model.UsedCapacity)
	}
	return modelMap, nil
}
