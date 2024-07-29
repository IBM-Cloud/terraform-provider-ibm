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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func DataSourceIbmRecoveryFetchUptierData() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmRecoveryFetchUptierDataRead,

		Schema: map[string]*schema.Schema{
			"archive_u_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Archive UID of the current restore.",
			},
			"data_size": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Specifies the amount of data in bytes estimated to be uptiered as part of the current restore job.",
			},
		},
	}
}

func dataSourceIbmRecoveryFetchUptierDataRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	fetchUptierDataOptions := &backuprecoveryv1.FetchUptierDataOptions{}

	fetchUptierDataOptions.SetArchiveUID(d.Get("archive_u_id").(string))

	fetchUptierDataResponse, response, err := backupRecoveryClient.FetchUptierDataWithContext(context, fetchUptierDataOptions)
	if err != nil {
		log.Printf("[DEBUG] FetchUptierDataWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("FetchUptierDataWithContext failed %s\n%s", err, response))
	}

	d.SetId(dataSourceIbmRecoveryFetchUptierDataID(d))

	if err = d.Set("data_size", flex.IntValue(fetchUptierDataResponse.DataSize)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting data_size: %s", err))
	}

	return nil
}

// dataSourceIbmRecoveryFetchUptierDataID returns a reasonable ID for the list.
func dataSourceIbmRecoveryFetchUptierDataID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
