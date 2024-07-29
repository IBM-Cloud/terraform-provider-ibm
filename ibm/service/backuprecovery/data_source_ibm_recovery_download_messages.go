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

func DataSourceIbmRecoveryDownloadMessages() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmRecoveryDownloadMessagesRead,

		Schema: map[string]*schema.Schema{
			"recovery_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies a unique ID of a Recovery.",
			},
		},
	}
}

func dataSourceIbmRecoveryDownloadMessagesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getRecoveryErrorsReportOptions := &backuprecoveryv1.GetRecoveryErrorsReportOptions{}

	getRecoveryErrorsReportOptions.SetID(d.Get("recovery_id").(string))

	response, err := backupRecoveryClient.GetRecoveryErrorsReportWithContext(context, getRecoveryErrorsReportOptions)
	if err != nil {
		log.Printf("[DEBUG] GetRecoveryErrorsReportWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetRecoveryErrorsReportWithContext failed %s\n%s", err, response))
	}

	d.SetId(dataSourceIbmRecoveryDownloadMessagesID(d))

	return nil
}

// dataSourceIbmRecoveryDownloadMessagesID returns a reasonable ID for the list.
func dataSourceIbmRecoveryDownloadMessagesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
