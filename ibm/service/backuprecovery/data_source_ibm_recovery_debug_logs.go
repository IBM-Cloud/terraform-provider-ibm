// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func DataSourceIbmRecoveryDebugLogs() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmRecoveryDebugLogsRead,

		Schema: map[string]*schema.Schema{
			"recovery_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the id of a Recovery job.",
			},
		},
	}
}

func dataSourceIbmRecoveryDebugLogsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getRecoveryDebugLogsOptions := &backuprecoveryv1.GetRecoveryDebugLogsOptions{}

	getRecoveryDebugLogsOptions.SetID(d.Get("recovery_id").(string))

	response, err := backupRecoveryClient.GetRecoveryDebugLogsWithContext(context, getRecoveryDebugLogsOptions)
	if err != nil {
		log.Printf("[DEBUG] GetRecoveryDebugLogsWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetRecoveryDebugLogsWithContext failed %s\n%s", err, response))
	}

	d.SetId(dataSourceIbmRecoveryDebugLogsID(d))

	return nil
}

// dataSourceIbmRecoveryDebugLogsID returns a reasonable ID for the list.
func dataSourceIbmRecoveryDebugLogsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
