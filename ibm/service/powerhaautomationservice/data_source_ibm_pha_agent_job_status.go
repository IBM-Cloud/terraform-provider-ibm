// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.113.1-d76630af-20260320-135953
 */

package powerhaautomationservice

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/DRAutomation/dra-go-sdk/powerhaautomationservicev1"
)

func DataSourceIBMPhaAgentJobStatus() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPhaAgentJobStatusRead,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique identifier of the provisioned instance.",
			},
			"job_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique ID to track the pha agent file download.",
			},
			"accept_language": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The language requested for the return document.",
			},
			"if_none_match": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ETag for conditional requests (optional).",
			},
			"bytes_downloaded": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of bytes downloaded so far.",
			},
			"creation_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when the job was created.",
			},
			"file_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the file that has been downloaded.",
			},
			"last_updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of the last update for this status.",
			},
			"service_instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Identifier of the service instance associated with the deployment.",
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current status of the deployment (e.g., running, completed, failed).",
			},
			"total_bytes": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Total size in bytes of the file that has to be downloaded.",
			},
			"vm_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Identifier of the virtual machine involved in the deployment.",
			},
		},
	}
}

func dataSourceIBMPhaAgentJobStatusRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	powerhaAutomationServiceClient, err := meta.(conns.ClientSession).PowerhaAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pha_agent_job_status", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getPhaAgentFileDownloadJobStatusOptions := &powerhaautomationservicev1.GetPhaAgentFileDownloadJobStatusOptions{}

	getPhaAgentFileDownloadJobStatusOptions.SetPhaInstanceID(d.Get("instance_id").(string))
	getPhaAgentFileDownloadJobStatusOptions.SetPhaJobID(d.Get("job_id").(string))
	if _, ok := d.GetOk("accept_language"); ok {
		getPhaAgentFileDownloadJobStatusOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	if _, ok := d.GetOk("if_none_match"); ok {
		getPhaAgentFileDownloadJobStatusOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	phaAgentJobStatusResponse, response, err := powerhaAutomationServiceClient.GetPhaAgentFileDownloadJobStatusWithContext(context, getPhaAgentFileDownloadJobStatusOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("GetPhaAgentFileDownloadJobStatusWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"GetPhaAgentFileDownloadJobStatusWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "ibm_pha_agent_job_status", "create")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
	}
	// if err != nil {
	// 	tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetPhaAgentFileDownloadJobStatusWithContext failed: %s", err.Error()), "(Data) ibm_pha_agent_job_status", "read")
	// 	log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
	// 	return tfErr.GetDiag()
	// }

	d.SetId(dataSourceIBMPhaAgentJobStatusID(d))

	if !core.IsNil(phaAgentJobStatusResponse.BytesDownloaded) {
		if err = d.Set("bytes_downloaded", flex.IntValue(phaAgentJobStatusResponse.BytesDownloaded)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting bytes_downloaded: %s", err), "(Data) ibm_pha_agent_job_status", "read", "set-bytes_downloaded").GetDiag()
		}
	}

	if !core.IsNil(phaAgentJobStatusResponse.CreationAt) {
		if err = d.Set("creation_at", flex.DateTimeToString(phaAgentJobStatusResponse.CreationAt)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting creation_at: %s", err), "(Data) ibm_pha_agent_job_status", "read", "set-creation_at").GetDiag()
		}
	}

	if !core.IsNil(phaAgentJobStatusResponse.FileName) {
		if err = d.Set("file_name", phaAgentJobStatusResponse.FileName); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting file_name: %s", err), "(Data) ibm_pha_agent_job_status", "read", "set-file_name").GetDiag()
		}
	}

	if !core.IsNil(phaAgentJobStatusResponse.LastUpdatedAt) {
		if err = d.Set("last_updated_at", flex.DateTimeToString(phaAgentJobStatusResponse.LastUpdatedAt)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting last_updated_at: %s", err), "(Data) ibm_pha_agent_job_status", "read", "set-last_updated_at").GetDiag()
		}
	}

	if !core.IsNil(phaAgentJobStatusResponse.ServiceInstanceID) {
		if err = d.Set("service_instance_id", phaAgentJobStatusResponse.ServiceInstanceID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting service_instance_id: %s", err), "(Data) ibm_pha_agent_job_status", "read", "set-service_instance_id").GetDiag()
		}
	}

	if !core.IsNil(phaAgentJobStatusResponse.Status) {
		if err = d.Set("status", phaAgentJobStatusResponse.Status); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_pha_agent_job_status", "read", "set-status").GetDiag()
		}
	}

	if !core.IsNil(phaAgentJobStatusResponse.TotalBytes) {
		if err = d.Set("total_bytes", flex.IntValue(phaAgentJobStatusResponse.TotalBytes)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting total_bytes: %s", err), "(Data) ibm_pha_agent_job_status", "read", "set-total_bytes").GetDiag()
		}
	}

	if !core.IsNil(phaAgentJobStatusResponse.VMID) {
		if err = d.Set("vm_id", phaAgentJobStatusResponse.VMID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting vm_id: %s", err), "(Data) ibm_pha_agent_job_status", "read", "set-vm_id").GetDiag()
		}
	}

	return nil
}

// dataSourceIBMPhaAgentJobStatusID returns a reasonable ID for the list.
func dataSourceIBMPhaAgentJobStatusID(d *schema.ResourceData) string {
	parts := strings.Split(d.Get("instance_id").(string), ":")
	if len(parts) > 7 {
		return parts[7]
	}
	return d.Get("instance_id").(string)
}
