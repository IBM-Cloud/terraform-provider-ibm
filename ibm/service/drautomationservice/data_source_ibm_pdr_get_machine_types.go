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
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

func DataSourceIBMPdrGetMachineTypes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPdrGetMachineTypesRead,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "instance id of instance to provision.",
			},
			"primary_workspace_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The primary Power virtual server workspace name.",
			},
			"accept_language": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The language requested for the return document.",
			},
			"standby_workspace_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The standby Power virtual server workspace name.",
			},
			"workspaces": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of workspaces with their machine types.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"machine_types": {
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMPdrGetMachineTypesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_get_machine_types", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getMachineTypeOptions := &drautomationservicev1.GetMachineTypeOptions{}

	getMachineTypeOptions.SetInstanceID(d.Get("instance_id").(string))
	getMachineTypeOptions.SetPrimaryWorkspaceName(d.Get("primary_workspace_name").(string))
	if _, ok := d.GetOk("accept_language"); ok {
		getMachineTypeOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}

	if _, ok := d.GetOk("standby_workspace_name"); ok {
		getMachineTypeOptions.SetStandbyWorkspaceName(d.Get("standby_workspace_name").(string))
	}

	machineTypesByWorkspace, response, err := drAutomationServiceClient.GetMachineTypeWithContext(context, getMachineTypeOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("GetMachineTypeWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"GetMachineTypeWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "(Data) ibm_pdr_get_machine_types", "read")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMPdrGetMachineTypesID(d))

	if !core.IsNil(machineTypesByWorkspace.Workspaces) {
		var workspacesList []map[string]interface{}
		for name, types := range machineTypesByWorkspace.Workspaces {
			entry := map[string]interface{}{
				"name":          name,
				"machine_types": convertToInterfaceList(types),
			}
			workspacesList = append(workspacesList, entry)
		}

		if err := d.Set("workspaces", workspacesList); err != nil {
			return flex.DiscriminatedTerraformErrorf(
				err, fmt.Sprintf("Error setting workspaces: %s", err),
				"(Data) ibm_pdr_get_machine_types", "read", "set-workspaces",
			).GetDiag()
		}
	}

	return nil
}

// dataSourceIBMPdrGetMachineTypesID returns a reasonable ID for the list.
func dataSourceIBMPdrGetMachineTypesID(d *schema.ResourceData) string {
	parts := strings.Split(d.Get("instance_id").(string), ":")
	if len(parts) > 7 {
		return parts[7]
	}
	return d.Get("instance_id").(string)
}

func convertToInterfaceList(items []string) []interface{} {
	if items == nil {
		return nil
	}

	interfaceList := make([]interface{}, len(items))
	for i, v := range items {
		interfaceList[i] = v
	}
	return interfaceList
}
