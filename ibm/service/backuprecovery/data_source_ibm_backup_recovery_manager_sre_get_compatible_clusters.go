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

func DataSourceIbmBackupRecoveryManagerSreGetCompatibleClusters() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryManagerSreGetCompatibleClustersRead,

		Schema: map[string]*schema.Schema{
			"release_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceIbmBackupRecoveryManagerSreGetCompatibleClustersRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	heliosSreApiClient, err := meta.(conns.ClientSession).BackupRecoveryManagerSreV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_sre_get_compatible_clusters", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	compatibleClustersForReleaseOptions := &backuprecoveryv1.CompatibleClustersForReleaseOptions{}

	if _, ok := d.GetOk("release_version"); ok {
		compatibleClustersForReleaseOptions.SetReleaseVersion(d.Get("release_version").(string))
	}

	_, _, err = heliosSreApiClient.CompatibleClustersForReleaseWithContext(context, compatibleClustersForReleaseOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CompatibleClustersForReleaseWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_manager_sre_get_compatible_clusters", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBackupRecoveryManagerSreGetCompatibleClustersID(d))

	return nil
}

// dataSourceIbmBackupRecoveryManagerSreGetCompatibleClustersID returns a reasonable ID for the list.
func dataSourceIbmBackupRecoveryManagerSreGetCompatibleClustersID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
