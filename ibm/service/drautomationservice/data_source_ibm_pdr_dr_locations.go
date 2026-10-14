// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.116.0-df613dbc-20260803-154903
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

func DataSourceIBMPdrDrLocations() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPdrDrLocationsRead,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Service Instance ID.",
			},
			"accept_language": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The language requested for the return document.",
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
							Description: "The name of the Power virtual server DR location .",
						},
						"power_edge_router": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether the region is power edge router enabled or not.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMPdrDrLocationsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_dr_locations", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getDrLocationOptions := &drautomationservicev1.GetDrLocationOptions{}

	getDrLocationOptions.SetInstanceID(d.Get("instance_id").(string))
	if _, ok := d.GetOk("accept_language"); ok {
		getDrLocationOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}

	getDrLocationsResponse, response, err := drAutomationServiceClient.GetDrLocationWithContext(context, getDrLocationOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("GetDrLocationsWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"GetDrLocation failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "(Data) ibm_pdr_dr_locations", "read")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMPdrDrLocationsID(d))

	drLocations := []map[string]interface{}{}
	for _, drLocationsItem := range getDrLocationsResponse.DrLocations {
		drLocationsItemMap, err := DataSourceIBMPdrDrLocationsDrLocationToMap(&drLocationsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_dr_locations", "read", "dr_locations-to-map").GetDiag()
		}
		drLocations = append(drLocations, drLocationsItemMap)
	}
	if err = d.Set("dr_locations", drLocations); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting dr_locations: %s", err), "(Data) ibm_pdr_dr_locations", "read", "set-dr_locations").GetDiag()
	}

	return nil
}

// dataSourceIBMPdrDrLocationsID returns a reasonable ID for the list.
func dataSourceIBMPdrDrLocationsID(d *schema.ResourceData) string {
	parts := strings.Split(d.Get("instance_id").(string), ":")
	if len(parts) > 7 {
		return parts[7]
	}
	return d.Get("instance_id").(string)
}

func DataSourceIBMPdrDrLocationsDrLocationToMap(model *drautomationservicev1.DrLocation) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.PowerEdgeRouter != nil {
		modelMap["power_edge_router"] = *model.PowerEdgeRouter
	}
	return modelMap, nil
}
