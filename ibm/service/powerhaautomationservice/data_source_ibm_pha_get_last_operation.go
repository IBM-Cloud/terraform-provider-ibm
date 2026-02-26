// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
 */

package powerhaautomationservice

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.ibm.com/DRAutomation/dra-go-sdk/powerhaautomationservicev1"
)

func DataSourceIBMPhaGetLastOperation() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPhaGetLastOperationRead,

		Schema: map[string]*schema.Schema{
			"pha_instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "instance id of instance to provision.",
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
			"deployment_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the deployment associated with the service instance.",
			},
			"provision_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique identifier for the provisioning operation.",
			},
			"resource_group": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource Group.",
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current operational status of the service instance.",
			},
		},
	}
}

func dataSourceIBMPhaGetLastOperationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	powerhaAutomationServiceClient, err := meta.(conns.ClientSession).PowerhaAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pha_get_last_operation", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getLastOperationOptions := &powerhaautomationservicev1.GetLastOperationOptions{}

	getLastOperationOptions.SetPhaInstanceID(d.Get("pha_instance_id").(string))
	if _, ok := d.GetOk("accept_language"); ok {
		getLastOperationOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	if _, ok := d.GetOk("if_none_match"); ok {
		getLastOperationOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	serviceInstancePhaStatus, response, err := powerhaAutomationServiceClient.GetLastOperationWithContext(context, getLastOperationOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("GetLastOperationWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"GetLastOperationWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "ibm_pha_get_last_operation", "create")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()

		// tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetLastOperationWithContext failed: %s", err.Error()), "(Data) ibm_pha_get_last_operation", "read")
		// log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		// return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMPhaGetLastOperationID(d))

	if err = d.Set("deployment_name", serviceInstancePhaStatus.DeploymentName); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting deployment_name: %s", err), "(Data) ibm_pha_get_last_operation", "read", "set-deployment_name").GetDiag()
	}

	if err = d.Set("provision_id", serviceInstancePhaStatus.ProvisionID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting provision_id: %s", err), "(Data) ibm_pha_get_last_operation", "read", "set-provision_id").GetDiag()
	}

	if err = d.Set("resource_group", serviceInstancePhaStatus.ResourceGroup); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group: %s", err), "(Data) ibm_pha_get_last_operation", "read", "set-resource_group").GetDiag()
	}

	if err = d.Set("status", serviceInstancePhaStatus.Status); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_pha_get_last_operation", "read", "set-status").GetDiag()
	}

	return nil
}

// dataSourceIBMPhaGetLastOperationID returns a reasonable ID for the list.
func dataSourceIBMPhaGetLastOperationID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
