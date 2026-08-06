// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
)

func DataSourceIbmLogsLogDataRetentionTags() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmLogsLogDataRetentionTagsRead,

		Schema: map[string]*schema.Schema{
			"tags": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of editable archive retention tags, excluding non-editable tags such as Default.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceIbmLogsLogDataRetentionTagsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsClient, err := meta.(conns.ClientSession).LogsV0()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_logs_log_data_retention_tags", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	region := getLogsInstanceRegion(logsClient, d)
	instanceId := d.Get("instance_id").(string)
	logsClient, err = getClientWithLogsInstanceEndpoint(logsClient, meta, instanceId, region, getLogsInstanceEndpointType(logsClient, d))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_logs_log_data_retention_tags", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getLogDataRetentionTagsOptions := &logsv0.GetLogDataRetentionTagsOptions{}

	logDataRetentionTags, response, err := logsClient.GetLogDataRetentionTagsWithContext(context, getLogDataRetentionTagsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetLogDataRetentionTagsWithContext failed: %s", err.Error()), "(Data) ibm_logs_log_data_retention_tags", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	_ = response

	d.SetId(fmt.Sprintf("%s/%s", region, instanceId))

	if !core.IsNil(logDataRetentionTags.Tags) {
		if err = d.Set("tags", logDataRetentionTags.Tags); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting tags: %s", err), "(Data) ibm_logs_log_data_retention_tags", "read")
			return tfErr.GetDiag()
		}
	}

	return nil
}
