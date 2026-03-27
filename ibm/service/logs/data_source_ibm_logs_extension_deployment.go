// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
)

func DataSourceIbmLogsExtensionDeployment() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmLogsExtensionDeploymentRead,

		Schema: map[string]*schema.Schema{
			"logs_extension_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier of the extension.",
			},
			"version": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The version of the Extension revision to deploy.",
			},
			"item_ids": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of Extension item IDs to deploy.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"applications": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Applications that the Extension is deployed for. When this is empty, it is applied to all applications.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"subsystems": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Subsystems that the Extension is deployed. When this is empty, it is applied to all subsystems.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceIbmLogsExtensionDeploymentRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsClient, err := meta.(conns.ClientSession).LogsV0()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_extension_deployment", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	region := getLogsInstanceRegion(logsClient, d)
	instanceId := d.Get("instance_id").(string)
	logsClient, err = getClientWithLogsInstanceEndpoint(logsClient, meta, instanceId, region, getLogsInstanceEndpointType(logsClient, d))
	if err != nil {
		return diag.FromErr(fmt.Errorf("Unable to get updated logs instance client"))
	}

	getExtensionDeploymentOptions := &logsv0.GetExtensionDeploymentOptions{}

	getExtensionDeploymentOptions.SetID(d.Get("logs_extension_id").(string))

	extensionDeployment, _, err := logsClient.GetExtensionDeploymentWithContext(context, getExtensionDeploymentOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetExtensionDeploymentWithContext failed: %s", err.Error()), "(Data) ibm_logs_extension_deployment", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*getExtensionDeploymentOptions.ID)

	if err = d.Set("version", extensionDeployment.Version); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting version: %s", err), "(Data) ibm_logs_extension_deployment", "read", "set-version").GetDiag()
	}

	itemIds := []interface{}{}
	for _, itemIdsItem := range extensionDeployment.ItemIds {
		itemIds = append(itemIds, itemIdsItem)
	}
	if err = d.Set("item_ids", itemIds); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting item_ids: %s", err), "(Data) ibm_logs_extension_deployment", "read", "set-item_ids").GetDiag()
	}

	if !core.IsNil(extensionDeployment.Applications) {
		applications := []interface{}{}
		for _, applicationsItem := range extensionDeployment.Applications {
			applications = append(applications, applicationsItem)
		}
		if err = d.Set("applications", applications); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting applications: %s", err), "(Data) ibm_logs_extension_deployment", "read", "set-applications").GetDiag()
		}
	}

	if !core.IsNil(extensionDeployment.Subsystems) {
		subsystems := []interface{}{}
		for _, subsystemsItem := range extensionDeployment.Subsystems {
			subsystems = append(subsystems, subsystemsItem)
		}
		if err = d.Set("subsystems", subsystems); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting subsystems: %s", err), "(Data) ibm_logs_extension_deployment", "read", "set-subsystems").GetDiag()
		}
	}

	return nil
}
