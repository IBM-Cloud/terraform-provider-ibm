// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
 */

package drautomationservice

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/IBM/dra-go-sdk/drautomationservicev1"
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
			"managed_vm_list": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vm_id":           {Type: schema.TypeString, Computed: true},
						"core":            {Type: schema.TypeString, Computed: true},
						"dr_average_time": {Type: schema.TypeString, Computed: true},
						"dr_region":       {Type: schema.TypeString, Computed: true},
						"memory":          {Type: schema.TypeString, Computed: true},
						"region":          {Type: schema.TypeString, Computed: true},
						"vm_name":         {Type: schema.TypeString, Computed: true},
						"workgroup_name":  {Type: schema.TypeString, Computed: true},
						"workspace_name":  {Type: schema.TypeString, Computed: true},
					},
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

	// convertedMap := make(map[string]interface{}, len(managedVMMapResponse.ManagedVMList))
	list := make([]map[string]interface{}, 0, len(managedVMMapResponse.ManagedVMList))

	for vmID, vmDetails := range managedVMMapResponse.ManagedVMList {

		obj := map[string]interface{}{
			"vm_id": vmID,
		}

		if vmDetails.Core != nil {
			obj["core"] = *vmDetails.Core
		}
		if vmDetails.DrAverageTime != nil {
			obj["dr_average_time"] = *vmDetails.DrAverageTime
		}
		if vmDetails.DrRegion != nil {
			obj["dr_region"] = *vmDetails.DrRegion
		}
		if vmDetails.Memory != nil {
			obj["memory"] = *vmDetails.Memory
		}
		if vmDetails.Region != nil {
			obj["region"] = *vmDetails.Region
		}
		if vmDetails.VMName != nil {
			obj["vm_name"] = *vmDetails.VMName
		}
		if vmDetails.WorkgroupName != nil {
			obj["workgroup_name"] = *vmDetails.WorkgroupName
		}
		if vmDetails.WorkspaceName != nil {
			obj["workspace_name"] = *vmDetails.WorkspaceName
		}

		list = append(list, obj)
	}

	if err = d.Set("managed_vm_list", list); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting managed_vm_list: %s", err), "(Data) ibm_pdr_get_managed_vm_list", "read", "set-managed_vm_list").GetDiag()
	}

	return nil
}

// dataSourceIBMPdrGetManagedVMListID returns a reasonable ID for the list.
func dataSourceIBMPdrGetManagedVMListID(d *schema.ResourceData) string {
	parts := strings.Split(d.Get("instance_id").(string), ":")
	if len(parts) > 7 {
		return parts[7]
	}
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
