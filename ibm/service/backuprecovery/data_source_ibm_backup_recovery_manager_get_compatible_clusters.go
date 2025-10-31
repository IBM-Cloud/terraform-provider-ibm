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

func DataSourceIbmBackupRecoveryManagerGetCompatibleClusters() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryManagerGetCompatibleClustersRead,

		Schema: map[string]*schema.Schema{
			"release_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"compatible_clusters": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
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
							Description: "Specifies the current version of the cluster.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmBackupRecoveryManagerGetCompatibleClustersRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	managementSreApiClient, err := meta.(conns.ClientSession).BackupRecoveryManagerV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_compatible_clusters", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	compatibleClustersForReleaseOptions := &backuprecoveryv1.CompatibleClustersForReleaseOptions{}

	if _, ok := d.GetOk("release_version"); ok {
		compatibleClustersForReleaseOptions.SetReleaseVersion(d.Get("release_version").(string))
	}

	compatibleClusters, _, err := managementSreApiClient.CompatibleClustersForReleaseWithContext(context, compatibleClustersForReleaseOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CompatibleClustersForReleaseWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_manager_get_compatible_clusters", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBackupRecoveryManagerGetCompatibleClustersID(d))

	if !core.IsNil(compatibleClusters) {
		compatibleClustersResult := []map[string]interface{}{}
		for _, compatibleClustersItem := range compatibleClusters {
			compatibleClustersItemMap, err := DataSourceIbmBackupRecoveryManagerGetCompatibleClustersCompatibleClusterToMap(&compatibleClustersItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_compatible_clusters", "read", "compatible_clusters-to-map").GetDiag()
			}
			compatibleClustersResult = append(compatibleClustersResult, compatibleClustersItemMap)
		}
		if err = d.Set("compatible_clusters", compatibleClusters); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting compatible_clusters: %s", err), "(Data) ibm_backup_recovery_manager_get_compatible_clusters", "read", "set-compatible_clusters").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmBackupRecoveryManagerGetCompatibleClustersID returns a reasonable ID for the list.
func dataSourceIbmBackupRecoveryManagerGetCompatibleClustersID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBackupRecoveryManagerGetCompatibleClustersCompatibleClusterToMap(model *backuprecoveryv1.CompatibleCluster) (map[string]interface{}, error) {
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
	return modelMap, nil
}
