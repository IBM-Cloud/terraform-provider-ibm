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

func DataSourceIbmPdrValidateWorkspace() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmPdrValidateWorkspaceRead,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "instance id of instance to provision.",
			},
			"workspace_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "standBy workspaceID value.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "crn value.",
			},
			"location_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "schematic_workspace_id value.",
			},
			"if_none_match": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ETag for conditional requests (optional).",
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIbmPdrValidateWorkspaceRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_validate_workspace", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	validatePowerVsWorkspaceOptions := &drautomationservicev1.ValidatePowerVsWorkspaceOptions{}

	validatePowerVsWorkspaceOptions.SetInstanceID(d.Get("instance_id").(string))
	validatePowerVsWorkspaceOptions.SetWorkspaceID(d.Get("workspace_id").(string))
	validatePowerVsWorkspaceOptions.SetCrn(d.Get("crn").(string))
	validatePowerVsWorkspaceOptions.SetLocationURL(d.Get("location_url").(string))
	if _, ok := d.GetOk("if_none_match"); ok {
		validatePowerVsWorkspaceOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	validateWorkspaceResponse, _, err := drAutomationServiceClient.ValidatePowerVsWorkspaceWithContext(context, validatePowerVsWorkspaceOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ValidatePowerVsWorkspaceWithContext failed: %s", err.Error()), "(Data) ibm_pdr_validate_workspace", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmPdrValidateWorkspaceID(d))

	if !core.IsNil(validateWorkspaceResponse.Description) {
		if err = d.Set("description", validateWorkspaceResponse.Description); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting description: %s", err), "(Data) ibm_pdr_validate_workspace", "read", "set-description").GetDiag()
		}
	}

	if !core.IsNil(validateWorkspaceResponse.Status) {
		if err = d.Set("status", validateWorkspaceResponse.Status); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_pdr_validate_workspace", "read", "set-status").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmPdrValidateWorkspaceID returns a reasonable ID for the list.
func dataSourceIbmPdrValidateWorkspaceID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
