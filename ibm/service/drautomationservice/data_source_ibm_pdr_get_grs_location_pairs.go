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

func DataSourceIBMPdrGetGrsLocationPairs() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPdrGetGrsLocationPairsRead,

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
			"location_pairs": &schema.Schema{
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "A map of GRS location pairs where each key is a primary location and the value is its paired location.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceIBMPdrGetGrsLocationPairsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_get_grs_location_pairs", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getDrGrsLocationPairOptions := &drautomationservicev1.GetDrGrsLocationPairOptions{}

	getDrGrsLocationPairOptions.SetInstanceID(d.Get("instance_id").(string))
	if _, ok := d.GetOk("accept_language"); ok {
		getDrGrsLocationPairOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}

	getGrsLocationPairResponse, response, err := drAutomationServiceClient.GetDrGrsLocationPairWithContext(context, getDrGrsLocationPairOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("GetDrGrsLocationPairWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"GetDrGrsLocationPairWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "(Data) ibm_pdr_get_grs_location_pairs", "read")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMPdrGetGrsLocationPairsID(d))

	if !core.IsNil(getGrsLocationPairResponse.LocationPairs) {
		if err = d.Set("location_pairs", getGrsLocationPairResponse.LocationPairs); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting location_pairs: %s", err), "(Data) ibm_pdr_get_grs_location_pairs", "read", "set-location_pairs").GetDiag()
		}
	}

	return nil
}

// dataSourceIBMPdrGetGrsLocationPairsID returns a reasonable ID for the list.
func dataSourceIBMPdrGetGrsLocationPairsID(d *schema.ResourceData) string {
	parts := strings.Split(d.Get("instance_id").(string), ":")
	if len(parts) > 7 {
		return parts[7]
	}
	return d.Get("instance_id").(string)
}
