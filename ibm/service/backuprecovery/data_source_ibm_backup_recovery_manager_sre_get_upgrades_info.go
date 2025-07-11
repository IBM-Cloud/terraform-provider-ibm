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
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func DataSourceIbmBackupRecoveryManagerSreGetUpgradesInfo() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryManagerSreGetUpgradesInfoRead,

		Schema: map[string]*schema.Schema{
			"cluster_identifiers": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Fetch upgrade progress details for a list of cluster identifiers in format clusterId:clusterIncarnationId.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceIbmBackupRecoveryManagerSreGetUpgradesInfoRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	heliosSreApiClient, err := meta.(conns.ClientSession).BackupRecoveryManagerSreV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_sre_get_upgrades_info", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
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

	_, _, err = heliosSreApiClient.ClustersUpgradesInfoWithContext(context, clustersUpgradesInfoOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ClustersUpgradesInfoWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_manager_sre_get_upgrades_info", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBackupRecoveryManagerSreGetUpgradesInfoID(d))

	return nil
}

// dataSourceIbmBackupRecoveryManagerSreGetUpgradesInfoID returns a reasonable ID for the list.
func dataSourceIbmBackupRecoveryManagerSreGetUpgradesInfoID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
