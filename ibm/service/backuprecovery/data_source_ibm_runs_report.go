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

func DataSourceIbmRunsReport() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmRunsReportRead,

		Schema: map[string]*schema.Schema{
			"group_id": {
				Type:     schema.TypeString,
				Required: true,
				// ValidateFunc: validate.InvokeValidator("ibm_create_protection_group_run_request", "run_type"),
				Description: "Protection group id",
			},
			"run_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies a unique run id of the Protection Group run.",
			},
			"object_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the id of the object for which errors/warnings are to be returned.",
			},
			"file_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the downloaded type, i.e: success_files_list, default: success_files_list.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the name of the source being backed up.",
			},
		},
	}
}

func dataSourceIbmRunsReportRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getRunsReportOptions := &backuprecoveryv1.GetRunsReportOptions{}

	getRunsReportOptions.SetID(d.Get("group_id").(string))
	getRunsReportOptions.SetRunID(d.Get("run_id").(string))
	getRunsReportOptions.SetObjectID(d.Get("object_id").(string))
	if _, ok := d.GetOk("file_type"); ok {
		getRunsReportOptions.SetFileType(d.Get("file_type").(string))
	}
	if _, ok := d.GetOk("name"); ok {
		getRunsReportOptions.SetName(d.Get("name").(string))
	}

	response, err := backupRecoveryClient.GetRunsReportWithContext(context, getRunsReportOptions)
	if err != nil {
		log.Printf("[DEBUG] GetRunsReportWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetRunsReportWithContext failed %s\n%s", err, response))
	}

	d.SetId(dataSourceIbmRunsReportID(d))

	return nil
}

// dataSourceIbmRunsReportID returns a reasonable ID for the list.
func dataSourceIbmRunsReportID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
