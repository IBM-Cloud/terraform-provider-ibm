// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.104.0-b4a47c49-20250418-184351
 */

package mqcloud

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/mqcloud-go-sdk/mqcloudv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
)

func DataSourceIbmMqcloudQueueManagerOptions() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmMqcloudQueueManagerOptionsRead,

		Schema: map[string]*schema.Schema{
			"service_instance_guid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The GUID that uniquely identifies the MQ SaaS service instance.",
			},
			"locations": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of deployment locations.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"sizes": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of queue manager sizes.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"versions": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of queue manager versions.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"latest_version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The latest Queue manager version.",
			},
		},
	}
}

func dataSourceIbmMqcloudQueueManagerOptionsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mqcloudClient, err := meta.(conns.ClientSession).MqcloudV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_mqcloud_queue_manager_options", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getOptionsOptions := &mqcloudv1.GetOptionsOptions{}

	getOptionsOptions.SetServiceInstanceGuid(d.Get("service_instance_guid").(string))

	configurationOptions, _, err := mqcloudClient.GetOptionsWithContext(context, getOptionsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetOptionsWithContext failed: %s", err.Error()), "(Data) ibm_mqcloud_queue_manager_options", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmMqcloudQueueManagerOptionsID(d))

	if !core.IsNil(configurationOptions.Locations) {
		locations := []interface{}{}
		for _, locationsItem := range configurationOptions.Locations {
			locations = append(locations, locationsItem)
		}
		if err = d.Set("locations", locations); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting locations: %s", err), "(Data) ibm_mqcloud_queue_manager_options", "read", "set-locations").GetDiag()
		}
	}

	if !core.IsNil(configurationOptions.Sizes) {
		sizes := []interface{}{}
		for _, sizesItem := range configurationOptions.Sizes {
			sizes = append(sizes, sizesItem)
		}
		if err = d.Set("sizes", sizes); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting sizes: %s", err), "(Data) ibm_mqcloud_queue_manager_options", "read", "set-sizes").GetDiag()
		}
	}

	if !core.IsNil(configurationOptions.Versions) {
		versions := []interface{}{}
		for _, versionsItem := range configurationOptions.Versions {
			versions = append(versions, versionsItem)
		}
		if err = d.Set("versions", versions); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting versions: %s", err), "(Data) ibm_mqcloud_queue_manager_options", "read", "set-versions").GetDiag()
		}
	}

	if !core.IsNil(configurationOptions.LatestVersion) {
		if err = d.Set("latest_version", configurationOptions.LatestVersion); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting latest_version: %s", err), "(Data) ibm_mqcloud_queue_manager_options", "read", "set-latest_version").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmMqcloudQueueManagerOptionsID returns a reasonable ID for the list.
func dataSourceIbmMqcloudQueueManagerOptionsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
