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

func DataSourceIbmRecoveryDownloadFiles() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmRecoveryDownloadFilesRead,

		Schema: map[string]*schema.Schema{
			"recovery_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the id of a Recovery.",
			},
			"start_offset": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies the start offset of file chunk to be downloaded.",
			},
			"length": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies the length of bytes to download. This can not be greater than 8MB (8388608 byets).",
			},
			"file_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the downloaded type, i.e: error, success_files_list.",
			},
			"source_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the name of the source on which restore is done.",
			},
			"start_time": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the start time of restore task.",
			},
			"include_tenants": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specifies if objects of all the organizations under the hierarchy of the logged in user's organization should be returned.",
			},
		},
	}
}

func dataSourceIbmRecoveryDownloadFilesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	downloadFilesFromRecoveryOptions := &backuprecoveryv1.DownloadFilesFromRecoveryOptions{}

	downloadFilesFromRecoveryOptions.SetID(d.Get("recovery_id").(string))
	if _, ok := d.GetOk("start_offset"); ok {
		downloadFilesFromRecoveryOptions.SetStartOffset(int64(d.Get("start_offset").(int)))
	}
	if _, ok := d.GetOk("length"); ok {
		downloadFilesFromRecoveryOptions.SetLength(int64(d.Get("length").(int)))
	}
	if _, ok := d.GetOk("file_type"); ok {
		downloadFilesFromRecoveryOptions.SetFileType(d.Get("file_type").(string))
	}
	if _, ok := d.GetOk("source_name"); ok {
		downloadFilesFromRecoveryOptions.SetSourceName(d.Get("source_name").(string))
	}
	if _, ok := d.GetOk("start_time"); ok {
		downloadFilesFromRecoveryOptions.SetStartTime(d.Get("start_time").(string))
	}
	if _, ok := d.GetOk("include_tenants"); ok {
		downloadFilesFromRecoveryOptions.SetIncludeTenants(d.Get("include_tenants").(bool))
	}

	response, err := backupRecoveryClient.DownloadFilesFromRecoveryWithContext(context, downloadFilesFromRecoveryOptions)
	if err != nil {
		log.Printf("[DEBUG] DownloadFilesFromRecoveryWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DownloadFilesFromRecoveryWithContext failed %s\n%s", err, response))
	}

	d.SetId(dataSourceIbmRecoveryDownloadFilesID(d))

	return nil
}

// dataSourceIbmRecoveryDownloadFilesID returns a reasonable ID for the list.
func dataSourceIbmRecoveryDownloadFilesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
