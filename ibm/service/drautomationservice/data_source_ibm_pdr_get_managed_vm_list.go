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
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of managed VMs associated with the service instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vm_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique identifier of the VM.",
						},
						"vm_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the VM.",
						},
					},
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

	getDrManagedVmOptions := &drautomationservicev1.GetDrManagedVmOptions{}

	getDrManagedVmOptions.SetInstanceID(d.Get("instance_id").(string))
	if _, ok := d.GetOk("accept_language"); ok {
		getDrManagedVmOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	if _, ok := d.GetOk("if_none_match"); ok {
		getDrManagedVmOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	managedVmListResponse, _, err := drAutomationServiceClient.GetDrManagedVmWithContext(context, getDrManagedVmOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDrManagedVmWithContext failed: %s", err.Error()), "(Data) ibm_pdr_get_managed_vm_list", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmPdrGetManagedVmListID(d))

	managedVms := []map[string]interface{}{}
	for _, managedVmsItem := range managedVmListResponse.ManagedVms {
		managedVmsItemMap, err := DataSourceIbmPdrGetManagedVmListManagedVmListToMap(&managedVmsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_get_managed_vm_list", "read", "managed_vms-to-map").GetDiag()
		}
		managedVms = append(managedVms, managedVmsItemMap)
	}
	if err = d.Set("managed_vms", managedVms); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting managed_vms: %s", err), "(Data) ibm_pdr_get_managed_vm_list", "read", "set-managed_vms").GetDiag()
	}

	return nil
}

// dataSourceIbmPdrGetManagedVmListID returns a reasonable ID for the list.
func dataSourceIbmPdrGetManagedVmListID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmPdrGetManagedVmListManagedVmListToMap(model *drautomationservicev1.ManagedVmList) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.VmID != nil {
		modelMap["vm_id"] = *model.VmID
	}
	if model.VmName != nil {
		modelMap["vm_name"] = *model.VmName
	}
	return modelMap, nil
}
