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

func DataSourceIbmPdrGetManagedVmList() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmPdrGetManagedVmListRead,

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
			"if_none_match": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ETag for conditional requests (optional).",
			},
			"managed_vms": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceIbmPdrGetManagedVmListRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_get_managed_vm_list", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getDrManagedVmOptions := &drautomationservicev1.GetDrManagedVMOptions{}

	getDrManagedVmOptions.SetInstanceID(d.Get("instance_id").(string))
	if _, ok := d.GetOk("accept_language"); ok {
		getDrManagedVmOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	if _, ok := d.GetOk("if_none_match"); ok {
		getDrManagedVmOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	managedVmMapResponse, _, err := drAutomationServiceClient.GetDrManagedVMWithContext(context, getDrManagedVmOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDrManagedVmWithContext failed: %s", err.Error()), "(Data) ibm_pdr_get_managed_vm_list", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmPdrGetManagedVmListID(d))

	if !core.IsNil(managedVmMapResponse.ManagedVms) {
		convertedMap := make(map[string]interface{}, len(managedVmMapResponse.ManagedVms))
		for k, v := range managedVmMapResponse.ManagedVms {
			convertedMap[k] = v
		}
		if err = d.Set("managed_vms", flex.Flatten(convertedMap)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting managed_vms: %s", err), "(Data) ibm_pdr_get_managed_vm_list", "read", "set-managed_vms").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmPdrGetManagedVmListID returns a reasonable ID for the list.
func dataSourceIbmPdrGetManagedVmListID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmPdrGetManagedVmListManagedVmDetailsToMap(model *drautomationservicev1.ManagedVMDetails) (map[string]interface{}, error) {
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
