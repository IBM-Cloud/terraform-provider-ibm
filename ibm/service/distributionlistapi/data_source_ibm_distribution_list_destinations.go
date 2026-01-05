// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
*/

package distributionlistapi

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/platform-services-go-sdk/distributionlistapiv1"
)

func DataSourceIbmDistributionListDestinations() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmDistributionListDestinationsRead,

		Schema: map[string]*schema.Schema{
			"account_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The IBM Cloud account ID.",
			},
			"destinations": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Array of destination entries.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The GUID of the Event Notifications instance.",
						},
						"destination_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the destination.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmDistributionListDestinationsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	distributionListApiClient, err := meta.(conns.ClientSession).DistributionListApiV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_distribution_list_destinations", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listDistributionListDestinationsOptions := &distributionlistapiv1.ListDistributionListDestinationsOptions{}

	listDistributionListDestinationsOptions.SetAccountID(d.Get("account_id").(string))

	addDestinationResponseBodyCollection, _, err := distributionListApiClient.ListDistributionListDestinationsWithContext(context, listDistributionListDestinationsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListDistributionListDestinationsWithContext failed: %s", err.Error()), "(Data) ibm_distribution_list_destinations", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmDistributionListDestinationsID(d))

	destinations := []map[string]interface{}{}
	for _, destinationsItem := range addDestinationResponseBodyCollection.Destinations {
		destinationsItemMap, err := DataSourceIbmDistributionListDestinationsAddDestinationResponseBodyToMap(destinationsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_distribution_list_destinations", "read", "destinations-to-map").GetDiag()
		}
		destinations = append(destinations, destinationsItemMap)
	}
	if err = d.Set("destinations", destinations); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting destinations: %s", err), "(Data) ibm_distribution_list_destinations", "read", "set-destinations").GetDiag()
	}

	return nil
}

// dataSourceIbmDistributionListDestinationsID returns a reasonable ID for the list.
func dataSourceIbmDistributionListDestinationsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmDistributionListDestinationsAddDestinationResponseBodyToMap(model distributionlistapiv1.AddDestinationResponseBodyIntf) (map[string]interface{}, error) {
	if _, ok := model.(*distributionlistapiv1.AddDestinationResponseBodyEventNotificationDestination); ok {
		return DataSourceIbmDistributionListDestinationsAddDestinationResponseBodyEventNotificationDestinationToMap(model.(*distributionlistapiv1.AddDestinationResponseBodyEventNotificationDestination))
	} else if _, ok := model.(*distributionlistapiv1.AddDestinationResponseBody); ok {
		modelMap := make(map[string]interface{})
		model := model.(*distributionlistapiv1.AddDestinationResponseBody)
		if model.ID != nil {
			modelMap["id"] = model.ID.String()
		}
		if model.DestinationType != nil {
			modelMap["destination_type"] = *model.DestinationType
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized distributionlistapiv1.AddDestinationResponseBodyIntf subtype encountered")
	}
}

func DataSourceIbmDistributionListDestinationsAddDestinationResponseBodyEventNotificationDestinationToMap(model *distributionlistapiv1.AddDestinationResponseBodyEventNotificationDestination) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID.String()
	modelMap["destination_type"] = *model.DestinationType
	return modelMap, nil
}
