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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
)

func ResourceIbmLogsLogDataRetentionTags() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmLogsLogDataRetentionTagsCreate,
		ReadContext:   resourceIbmLogsLogDataRetentionTagsRead,
		UpdateContext: resourceIbmLogsLogDataRetentionTagsUpdate,
		DeleteContext: resourceIbmLogsLogDataRetentionTagsDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"tags": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				Description: "List of editable archive retention tags, excluding non-editable tags such as Default.",
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validate.InvokeValidator("ibm_logs_log_data_retention_tags", "tags"),
				},
				MinItems: 3,
				MaxItems: 3,
			},
		},
	}
}

func ResourceIbmLogsLogDataRetentionTagsValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "tags",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[a-zA-Z0-9_-]+$`,
			MinValueLength:             1,
			MaxValueLength:             256,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_logs_log_data_retention_tags", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmLogsLogDataRetentionTagsCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsClient, err := meta.(conns.ClientSession).LogsV0()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_logs_log_data_retention_tags", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	region := getLogsInstanceRegion(logsClient, d)
	instanceId := d.Get("instance_id").(string)
	logsClient, err = getClientWithLogsInstanceEndpoint(logsClient, meta, instanceId, region, getLogsInstanceEndpointType(logsClient, d))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_logs_log_data_retention_tags", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateLogDataRetentionTagsOptions := &logsv0.UpdateLogDataRetentionTagsOptions{}

	tags := []string{}
	for _, tagsItem := range d.Get("tags").([]interface{}) {
		tags = append(tags, tagsItem.(string))
	}
	updateLogDataRetentionTagsOptions.SetTags(tags)

	response, err := logsClient.UpdateLogDataRetentionTagsWithContext(context, updateLogDataRetentionTagsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateLogDataRetentionTagsWithContext failed: %s", err.Error()), "ibm_logs_log_data_retention_tags", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	_ = response

	d.SetId(fmt.Sprintf("%s/%s", region, instanceId))

	return resourceIbmLogsLogDataRetentionTagsRead(context, d, meta)
}

func resourceIbmLogsLogDataRetentionTagsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsClient, err := meta.(conns.ClientSession).LogsV0()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_logs_log_data_retention_tags", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	logsClient, region, instanceId, _, err := updateClientURLWithInstanceEndpoint(d.Id(), meta, logsClient, d)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_logs_log_data_retention_tags", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getLogDataRetentionTagsOptions := &logsv0.GetLogDataRetentionTagsOptions{}

	logDataRetentionTags, response, err := logsClient.GetLogDataRetentionTagsWithContext(context, getLogDataRetentionTagsOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetLogDataRetentionTagsWithContext failed: %s", err.Error()), "ibm_logs_log_data_retention_tags", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("instance_id", instanceId); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting instance_id: %s", err), "ibm_logs_log_data_retention_tags", "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("region", region); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting region: %s", err), "ibm_logs_log_data_retention_tags", "read")
		return tfErr.GetDiag()
	}
	if !core.IsNil(logDataRetentionTags.Tags) {
		if err = d.Set("tags", logDataRetentionTags.Tags); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting tags: %s", err), "ibm_logs_log_data_retention_tags", "read")
			return tfErr.GetDiag()
		}
	}

	return nil
}

func resourceIbmLogsLogDataRetentionTagsUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsClient, err := meta.(conns.ClientSession).LogsV0()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_logs_log_data_retention_tags", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	logsClient, _, _, _, err = updateClientURLWithInstanceEndpoint(d.Id(), meta, logsClient, d)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_logs_log_data_retention_tags", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateLogDataRetentionTagsOptions := &logsv0.UpdateLogDataRetentionTagsOptions{}

	hasChange := false

	if d.HasChange("tags") {
		tags := []string{}
		for _, tagsItem := range d.Get("tags").([]interface{}) {
			tags = append(tags, tagsItem.(string))
		}
		updateLogDataRetentionTagsOptions.SetTags(tags)
		hasChange = true
	}

	if hasChange {
		response, err := logsClient.UpdateLogDataRetentionTagsWithContext(context, updateLogDataRetentionTagsOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateLogDataRetentionTagsWithContext failed: %s", err.Error()), "ibm_logs_log_data_retention_tags", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		_ = response
	}

	return resourceIbmLogsLogDataRetentionTagsRead(context, d, meta)
}

func resourceIbmLogsLogDataRetentionTagsDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource cannot be deleted - it can only be deactivated by removing the archive bucket
	// We'll just remove it from state
	d.SetId("")
	return nil
}
