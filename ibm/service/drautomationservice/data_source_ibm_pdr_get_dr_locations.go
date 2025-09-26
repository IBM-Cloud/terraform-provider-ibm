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

func DataSourceIbmPdrGetDrLocations() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmPdrGetDrLocationsRead,

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
			"dr_locations": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of disaster recovery locations available for the service.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique identifier of the DR location.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the DR location.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmPdrGetDrLocationsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_get_dr_locations", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getDrLocationOptions := &drautomationservicev1.GetDrLocationOptions{}

	getDrLocationOptions.SetInstanceID(d.Get("instance_id").(string))
	if _, ok := d.GetOk("accept_language"); ok {
		getDrLocationOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	if _, ok := d.GetOk("if_none_match"); ok {
		getDrLocationOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	getDrLocationsResponse, _, err := drAutomationServiceClient.GetDrLocationWithContext(context, getDrLocationOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDrLocationWithContext failed: %s", err.Error()), "(Data) ibm_pdr_get_dr_locations", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmPdrGetDrLocationsID(d))

	drLocations := []map[string]interface{}{}
	for _, drLocationsItem := range getDrLocationsResponse.DrLocations {
		drLocationsItemMap, err := DataSourceIbmPdrGetDrLocationsDrLocationToMap(&drLocationsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_get_dr_locations", "read", "dr_locations-to-map").GetDiag()
		}
		drLocations = append(drLocations, drLocationsItemMap)
	}
	if err = d.Set("dr_locations", drLocations); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting dr_locations: %s", err), "(Data) ibm_pdr_get_dr_locations", "read", "set-dr_locations").GetDiag()
	}

	return nil
}

// dataSourceIbmPdrGetDrLocationsID returns a reasonable ID for the list.
func dataSourceIbmPdrGetDrLocationsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmPdrGetDrLocationsDrLocationToMap(model *drautomationservicev1.DrLocation) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	return modelMap, nil
}
