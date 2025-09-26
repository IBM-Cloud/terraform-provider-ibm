// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.105.0-3c13b041-20250605-193116
*/

package drautomationservice

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

func DataSourceIbmPdrValidateClustertype() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmPdrValidateClustertypeRead,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "instance id of instance to provision.",
			},
			"orchestrator_cluster_type": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "orchestrator cluster type value.",
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
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Human-readable message explaining the cluster type validation result.",
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of the cluster type validation (for example, valid, invalid, or error).",
			},
		},
	}
}

func dataSourceIbmPdrValidateClustertypeRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_validate_clustertype", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getValidateClusterTypeOptions := &drautomationservicev1.GetValidateClusterTypeOptions{}

	getValidateClusterTypeOptions.SetInstanceID(d.Get("instance_id").(string))
	getValidateClusterTypeOptions.SetOrchestratorClusterType(d.Get("orchestrator_cluster_type").(string))
	if _, ok := d.GetOk("accept_language"); ok {
		getValidateClusterTypeOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	if _, ok := d.GetOk("if_none_match"); ok {
		getValidateClusterTypeOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	validateClusterTypeResponse, _, err := drAutomationServiceClient.GetValidateClusterTypeWithContext(context, getValidateClusterTypeOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetValidateClusterTypeWithContext failed: %s", err.Error()), "(Data) ibm_pdr_validate_clustertype", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmPdrValidateClustertypeID(d))

	if !core.IsNil(validateClusterTypeResponse.Description) {
		if err = d.Set("description", validateClusterTypeResponse.Description); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting description: %s", err), "(Data) ibm_pdr_validate_clustertype", "read", "set-description").GetDiag()
		}
	}

	if !core.IsNil(validateClusterTypeResponse.Status) {
		if err = d.Set("status", validateClusterTypeResponse.Status); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_pdr_validate_clustertype", "read", "set-status").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmPdrValidateClustertypeID returns a reasonable ID for the list.
func dataSourceIbmPdrValidateClustertypeID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
