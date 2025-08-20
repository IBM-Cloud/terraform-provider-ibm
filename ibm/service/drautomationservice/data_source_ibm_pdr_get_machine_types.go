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

func DataSourceIbmPdrGetMachineTypes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmPdrGetMachineTypesRead,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "instance id of instance to provision.",
			},
			"primary_workspace_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Primary Workspace Name.",
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
			"standby_workspace_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Standby Workspace Name.",
			},
			"workspaces": &schema.Schema{
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Map of workspace IDs to lists of machine types.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceIbmPdrGetMachineTypesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_get_machine_types", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	machinetypesDetailsOptions := &drautomationservicev1.MachinetypesDetailsOptions{}

	machinetypesDetailsOptions.SetInstanceID(d.Get("instance_id").(string))
	machinetypesDetailsOptions.SetPrimaryWorkspaceName(d.Get("primary_workspace_name").(string))
	if _, ok := d.GetOk("accept_language"); ok {
		machinetypesDetailsOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	if _, ok := d.GetOk("if_none_match"); ok {
		machinetypesDetailsOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}
	if _, ok := d.GetOk("standby_workspace_name"); ok {
		machinetypesDetailsOptions.SetStandbyWorkspaceName(d.Get("standby_workspace_name").(string))
	}

	machineTypesByWorkspace, _, err := drAutomationServiceClient.MachinetypesDetailsWithContext(context, machinetypesDetailsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("MachinetypesDetailsWithContext failed: %s", err.Error()), "(Data) ibm_pdr_get_machine_types", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmPdrGetMachineTypesID(d))

	if !core.IsNil(machineTypesByWorkspace.Workspaces) {
		if err = d.Set("workspaces", machineTypesByWorkspace.Workspaces); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting workspaces: %s", err), "(Data) ibm_pdr_get_machine_types", "read", "set-workspaces").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmPdrGetMachineTypesID returns a reasonable ID for the list.
func dataSourceIbmPdrGetMachineTypesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
