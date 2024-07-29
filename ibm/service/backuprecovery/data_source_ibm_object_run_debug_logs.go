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

func DataSourceIbmObjectRunDebugLogs() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmObjectRunDebugLogsRead,

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
				Description: "Specifies the id of the object for which debug logs are to be returned.",
			},
		},
	}
}

func dataSourceIbmObjectRunDebugLogsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getRunDebugLogsForObjectOptions := &backuprecoveryv1.GetRunDebugLogsForObjectOptions{}

	getRunDebugLogsForObjectOptions.SetID(d.Get("group_id").(string))
	getRunDebugLogsForObjectOptions.SetRunID(d.Get("run_id").(string))
	getRunDebugLogsForObjectOptions.SetObjectID(d.Get("object_id").(string))

	response, err := backupRecoveryClient.GetRunDebugLogsForObjectWithContext(context, getRunDebugLogsForObjectOptions)
	if err != nil {
		log.Printf("[DEBUG] GetRunDebugLogsForObjectWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetRunDebugLogsForObjectWithContext failed %s\n%s", err, response))
	}

	d.SetId(dataSourceIbmObjectRunDebugLogsID(d))

	return nil
}

// dataSourceIbmObjectRunDebugLogsID returns a reasonable ID for the list.
func dataSourceIbmObjectRunDebugLogsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
