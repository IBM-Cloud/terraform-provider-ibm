// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
 */

package powerhaautomationservice

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.ibm.com/DRAutomation/dra-go-sdk/powerhaautomationservicev1"
)

func DataSourceIBMPhaGetSupportedLocation() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPhaGetSupportedLocationRead,

		Schema: map[string]*schema.Schema{
			"pha_instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "instance id of instance to provision.",
			},
			"if_none_match": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ETag for conditional requests (optional).",
			},
			"dr_locations": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Array of supported DR locations.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique identifier for the location.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Human-readable name of the location.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMPhaGetSupportedLocationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	powerhaAutomationServiceClient, err := meta.(conns.ClientSession).PowerhaAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pha_get_supported_location", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getSupportedLocationOptions := &powerhaautomationservicev1.GetSupportedLocationOptions{}

	getSupportedLocationOptions.SetPhaInstanceID(d.Get("pha_instance_id").(string))
	if _, ok := d.GetOk("if_none_match"); ok {
		getSupportedLocationOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	phaSupportedLocationsResponse, response, err := powerhaAutomationServiceClient.GetSupportedLocationWithContext(context, getSupportedLocationOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("GetSupportedLocationWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"GetSupportedLocationWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "ibm_pha_get_supported_location", "create")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()

		// tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSupportedLocationWithContext failed: %s", err.Error()), "(Data) ibm_pha_get_supported_location", "read")
		// log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		// return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMPhaGetSupportedLocationID(d))

	drLocations := []map[string]interface{}{}
	for _, drLocationsItem := range phaSupportedLocationsResponse.DrLocations {
		drLocationsItemMap, err := DataSourceIBMPhaGetSupportedLocationPhaLocationToMap(&drLocationsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pha_get_supported_location", "read", "dr_locations-to-map").GetDiag()
		}
		drLocations = append(drLocations, drLocationsItemMap)
	}
	if err = d.Set("dr_locations", drLocations); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting dr_locations: %s", err), "(Data) ibm_pha_get_supported_location", "read", "set-dr_locations").GetDiag()
	}

	return nil
}

// dataSourceIBMPhaGetSupportedLocationID returns a reasonable ID for the list.
func dataSourceIBMPhaGetSupportedLocationID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIBMPhaGetSupportedLocationPhaLocationToMap(model *powerhaautomationservicev1.PhaLocation) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	return modelMap, nil
}
