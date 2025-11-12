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

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

func DataSourceIBMPdrGetManagedVMList() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPdrGetManagedVMListRead,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "instance id of instance to provision.",
			},
			"accept_language": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The language requested for the return document.",
			},
			"managed_vms": &schema.Schema{
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "A map where the key is the VM ID and the value is the corresponding ManagedVmDetails object.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceIBMPdrGetManagedVMListRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_get_managed_vm_list", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getDrManagedVMOptions := &drautomationservicev1.GetDrManagedVMOptions{}

	getDrManagedVMOptions.SetInstanceID(d.Get("instance_id").(string))
	if _, ok := d.GetOk("accept_language"); ok {
		getDrManagedVMOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	if _, ok := d.GetOk("if_none_match"); ok {
		getDrManagedVMOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	managedVMMapResponse, response, err := drAutomationServiceClient.GetDrManagedVMWithContext(context, getDrManagedVMOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("GetDrManagedVMWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"GetDrManagedVMWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "(Data) ibm_pdr_get_managed_vm_list", "read")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMPdrGetManagedVMListID(d))

	if !core.IsNil(managedVMMapResponse.ManagedVms) {
		convertedMap := make(map[string]interface{}, len(managedVMMapResponse.ManagedVms))
		for k, v := range managedVMMapResponse.ManagedVms {
			convertedMap[k] = v
		}
		if err = d.Set("managed_vms", flex.Flatten(convertedMap)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting managed_vms: %s", err), "(Data) ibm_pdr_get_managed_vm_list", "read", "set-managed_vms").GetDiag()
		}
	}

	return nil
}

// dataSourceIBMPdrGetManagedVMListID returns a reasonable ID for the list.
func dataSourceIBMPdrGetManagedVMListID(d *schema.ResourceData) string {
	return d.Get("instance_id").(string)
}

func DataSourceIBMPdrGetManagedVMListManagedVMDetailsToMap(model *drautomationservicev1.ManagedVMDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Core != nil {
		modelMap["core"] = *model.Core
	}
	if model.DrAverageTime != nil {
		modelMap["dr_average_time"] = *model.DrAverageTime
	}
	if model.DrRegion != nil {
		modelMap["dr_region"] = *model.DrRegion
	}
	if model.Memory != nil {
		modelMap["memory"] = *model.Memory
	}
	if model.Region != nil {
		modelMap["region"] = *model.Region
	}
	if model.VMName != nil {
		modelMap["vm_name"] = *model.VMName
	}
	if model.WorkgroupName != nil {
		modelMap["workgroup_name"] = *model.WorkgroupName
	}
	if model.WorkspaceName != nil {
		modelMap["workspace_name"] = *model.WorkspaceName
	}
	return modelMap, nil
}
