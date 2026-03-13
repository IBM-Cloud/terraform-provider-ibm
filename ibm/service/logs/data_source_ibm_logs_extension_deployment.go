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
	"github.com/IBM/logs-go-sdk/logsv0"
)

func DataSourceIbmLogsExtensionDeployment() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmLogsExtensionDeploymentRead,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the IBM Cloud Logs instance.",
			},
			"region": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The region of the IBM Cloud Logs instance.",
			},
			"endpoint_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "public or private.",
			},
			"logs_extension_deployment_id": &schema.Schema{
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
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_logs_extension_deployment", "read")
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

	getExtensionDeploymentOptions.SetID(d.Get("logs_extension_deployment_id").(string))

	extensionDeployment, response, err := logsClient.GetExtensionDeploymentWithContext(context, getExtensionDeploymentOptions)
	if err != nil {
		log.Printf("[DEBUG] GetExtensionDeploymentWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetExtensionDeploymentWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s", *getExtensionDeploymentOptions.ID))

	if err = d.Set("version", extensionDeployment.Version); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting version: %s", err))
	}

	if err = d.Set("item_ids", extensionDeployment.ItemIds); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting item_ids: %s", err))
	}

	if extensionDeployment.Applications != nil {
		if err = d.Set("applications", extensionDeployment.Applications); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting applications: %s", err))
		}
	}

	if extensionDeployment.Subsystems != nil {
		if err = d.Set("subsystems", extensionDeployment.Subsystems); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting subsystems: %s", err))
		}
	}

	return nil
}
