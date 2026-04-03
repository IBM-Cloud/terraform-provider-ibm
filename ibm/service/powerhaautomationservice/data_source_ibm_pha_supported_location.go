// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.113.1-d76630af-20260320-135953
*/

package powerhaautomationservice

import (
	"context"
	"fmt"
	"log"
	"strings"

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
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique identifier of the provisioned instance.",
			},
			"if_none_match": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ETag for conditional requests (optional).",
			},
			"locations": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Array of PowerVS locations where PowerHA service is supported.",
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

	getSupportedLocationOptions.SetPhaInstanceID(d.Get("instance_id").(string))
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
	}

	d.SetId(dataSourceIBMPhaGetSupportedLocationID(d))

	locations := []map[string]interface{}{}
	for _, locationsItem := range phaSupportedLocationsResponse.Locations {
		locationsItemMap, err := DataSourceIBMPhaGetSupportedLocationPhaLocationToMap(&locationsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pha_get_supported_location", "read", "locations-to-map").GetDiag()
		}
		locations = append(locations, locationsItemMap)
	}
	if err = d.Set("locations", locations); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting locations: %s", err), "(Data) ibm_pha_get_supported_location", "read", "set-locations").GetDiag()
	}

	return nil
}

// dataSourceIBMPhaGetSupportedLocationID returns a reasonable ID for the list.
func dataSourceIBMPhaGetSupportedLocationID(d *schema.ResourceData) string {
	parts := strings.Split(d.Get("instance_id").(string), ":")
	if len(parts) > 7 {
		return parts[7]
	}
	return d.Get("instance_id").(string)
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
