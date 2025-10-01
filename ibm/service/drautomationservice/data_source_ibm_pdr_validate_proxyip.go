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

func DataSourceIbmPdrValidateProxyip() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmPdrValidateProxyipRead,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "instance id of instance to provision.",
			},
			"proxyip": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "proxyip value.",
			},
			"vpc_location": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "vpc location value.",
			},
			"vpc_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "vpc id value.",
			},
			"if_none_match": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ETag for conditional requests (optional).",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Human-readable message explaining the proxy IP validation result.",
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of the proxy IP validation (for example, valid, invalid, or error).",
			},
			"warning": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether the proxy IP is valid but has an advisory (e.g., not in reserved IPs).",
			},
		},
	}
}

func dataSourceIbmPdrValidateProxyipRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_validate_proxyip", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getValidateProxyipOptions := &drautomationservicev1.GetValidateProxyipOptions{}

	getValidateProxyipOptions.SetInstanceID(d.Get("instance_id").(string))
	getValidateProxyipOptions.SetProxyip(d.Get("proxyip").(string))
	getValidateProxyipOptions.SetVpcLocation(d.Get("vpc_location").(string))
	getValidateProxyipOptions.SetVpcID(d.Get("vpc_id").(string))
	if _, ok := d.GetOk("if_none_match"); ok {
		getValidateProxyipOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	validateProxyipResponse, _, err := drAutomationServiceClient.GetValidateProxyipWithContext(context, getValidateProxyipOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetValidateProxyipWithContext failed: %s", err.Error()), "(Data) ibm_pdr_validate_proxyip", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmPdrValidateProxyipID(d))

	if !core.IsNil(validateProxyipResponse.Description) {
		if err = d.Set("description", validateProxyipResponse.Description); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting description: %s", err), "(Data) ibm_pdr_validate_proxyip", "read", "set-description").GetDiag()
		}
	}

	if !core.IsNil(validateProxyipResponse.Status) {
		if err = d.Set("status", validateProxyipResponse.Status); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_pdr_validate_proxyip", "read", "set-status").GetDiag()
		}
	}

	if !core.IsNil(validateProxyipResponse.Warning) {
		if err = d.Set("warning", validateProxyipResponse.Warning); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting warning: %s", err), "(Data) ibm_pdr_validate_proxyip", "read", "set-warning").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmPdrValidateProxyipID returns a reasonable ID for the list.
func dataSourceIbmPdrValidateProxyipID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
